package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const url3 = "https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/mrybtb.jsp"
const url4 = "https://sso.ecust.edu.cn/authserver/login?service=https%3A%2F%2Fworkflow.ecust.edu.cn%2Fdefault%2Fwork%2Fuust%2Fzxxsmryb%2Fmrybtb.jsp"
const url5 = "https://sso.ecust.edu.cn/authserver/login;"
const url6 = "?service=https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/mrybtb.jsp"
const contentType = "application/x-www-form-urlencoded"
const agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:100.0) Gecko/20100101 Firefox/100.0"
const cookie = "route=ff2bdac9793518841c6ac6099a5af441; org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=zh_CN; fid=99352; 99352userinfo=6cb2606b4382afd179d923599539fcf9c49d67c0c30ca5043e655d84b33e56a223ed28be35d4b8a06016e50eee50bfdac98bf515e9147871e19df3f81a99d14a; 99352UID=196243147; _uid=196243147; 99352enc=F8B7471E696BC21AD1E79E256F1BF088; UID=196243147; vc=F8B7471E696BC21AD1E79E256F1BF088; vc2=E3A597DA92A2C43B3F5C1B2B646050BA; uf=14b6b1b3f40d8f908c0e8a3c35437931ac47db3197db753f4b008a02658943f2d95d3e405e0cc04d01786dfe1cb15b83428b5a98cdb27be888b83130e7eb4704e0cc755ad25f904e0246270a56330e6c57113d8a7ab747e4760b8a420cebd0c4e9022a21abd4cdc8; _d=1650501135775; vc3_mirror=FLmWJQ8sTPIK%2B1uICLxI2gkJ2f5QkBbAgttb1RK05NMpzBlWHQr7W27vyzwvFfQvbPVMiM6MB76NqHmZP2sN8eoaojpX72WQFPdXBVOOLtex6ZfPcnhprmkrRieeVokVKZh7Lt9tczFK9gApQIPcpqbyKO35414grkySmYAsOFs%3D14d0440231bdf5210e2a5bd2b2732c70; DSSTASH_LOG=C_38-UN_781-US_196243147-T_1650501135777; JSESSIONID_AUTH=upacL4gnyr2EWcEAy1FkJLEDgNxSZREJCI6pmO1-zwV_O1tn8BYi!879369270"
const part1 = "username=Y80210177&password=S08087878&lt="
const part2 = "&dllt=userNamePasswordLogin&execution="
const part3 = "&_eventId=submit&rmShown=1"

func main() {
	//login()
	// 每日6点30分执行 打印“早上好” 与 “起床学习”的两个任务
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	dateTime := time.Date(now.Year(), now.Month(), now.Day(), 8, rand.Intn(59), 0, 0, now.Location())
	DailyCron(dateTime, login)

}

var login = func() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	log.Println("--------------------------------")
	request1, err := http.NewRequest("GET", url3, nil)
	if err != nil {
		log.Println("get failed, err:", err)
		log.Fatal(err)
	}

	request1.Header.Add("User-Agent", agent)
	response1, _ := client.Do(request1)
	loc, err := response1.Location()

	request2, err := http.NewRequest("GET", loc.String(), nil)
	if err != nil {
		log.Println("get failed, err:", err)
		log.Fatal(err)
	}
	request2.Header.Add("User-Agent", agent)
	response2, _ := client.Do(request2)
	doc, err := goquery.NewDocumentFromReader(response2.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response2.Body)
	var lt, execution string
	doc.Find("input[name='lt']").Each(func(i int, s *goquery.Selection) {
		lt, _ = s.Attr("value")
	})
	doc.Find("input[name='execution']").Each(func(i int, s *goquery.Selection) {
		execution, _ = s.Attr("value")
	})
	coo2 := fmt.Sprintf("%s=%s", response2.Cookies()[1].Name, response2.Cookies()[1].Value)

	LoginUrl := fmt.Sprintf("%s%s%s", url5, coo2, url6)
	buf := fmt.Sprintf("%s%s%s%s%s", part1, lt, part2, execution, part3)
	request3, err := http.NewRequest("POST", LoginUrl, strings.NewReader(buf))
	if err != nil {
		log.Println("get failed, err:", err)
		log.Fatal(err)
	}
	request3.Header.Add("User-Agent", agent)
	request3.Header.Add("Origin", "https://sso.ecust.edu.cn")
	request3.Header.Add("Content-Type", contentType)
	request3.Header.Add("Referer", "https://sso.ecust.edu.cn/authserver/login?service=https%3A%2F%2Fworkflow.ecust.edu.cn%2Fdefault%2Fwork%2Fuust%2Fzxxsmryb%2Fmrybcn.jsp")
	request3.Header.Add("Sec-Fetch-Site", "same-origin")
	cookie := http.Cookie{Name: response2.Cookies()[0].Name, Value: response2.Cookies()[0].Value, Expires: time.Now().Add(111 * time.Second)}
	request3.AddCookie(&cookie)
	cookie = http.Cookie{Name: response2.Cookies()[1].Name, Value: response2.Cookies()[1].Value, Expires: time.Now().Add(111 * time.Second)}
	request3.AddCookie(&cookie)
	cookie = http.Cookie{Name: "org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE", Value: "en", Expires: time.Now().Add(111 * time.Second)}
	request3.AddCookie(&cookie)

	response3, _ := client.Do(request3)
	loc, err = response3.Location()
	client2 := &http.Client{}
	//request4, err := http.NewRequest("GET", loc.String(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//request4.Header.Add("User-Agent", agent)
	//
	//cookie = http.Cookie{Name: response3.Cookies()[2].Name, Value: response3.Cookies()[2].Value, Expires: time.Now().Add(111 * time.Second)}
	//request4.AddCookie(&cookie)
	//cookie = http.Cookie{Name: response1.Cookies()[0].Name, Value: response1.Cookies()[0].Value, Expires: time.Now().Add(111 * time.Second)}
	//request4.AddCookie(&cookie)
	//cookie = http.Cookie{Name: response1.Cookies()[1].Name, Value: response1.Cookies()[1].Value, Expires: time.Now().Add(111 * time.Second)}
	//request4.AddCookie(&cookie)
	//response4, _ := client2.Do(request4)
	//log.Println("response 4 status :  ", response4.Status)

	request5, err := http.NewRequest("POST", "https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/com.sudytech.work.uust.zxxsmryb.xxsmryb.queryDrsj.biz.ext", nil)
	if err != nil {
		log.Fatal(err)
	}
	request5.Header.Add("User-Agent", agent)
	request5.Header.Add("Content-Type", contentType)
	request5.Header.Add("Origin", "https://workflow.ecust.edu.cn")
	request5.Header.Add("Referer", "https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/mrybtb.jsp")
	request5.Header.Add("Sec-Fetch-Site", "same-origin")
	request5.Header.Add("Sec-Fetch-Dest", "empty")
	request5.Header.Add("Sec-Fetch-Mode", "cors")
	cookie = http.Cookie{Name: response3.Cookies()[2].Name, Value: response3.Cookies()[2].Value, Expires: time.Now().Add(111 * time.Second)}
	request5.AddCookie(&cookie)
	cookie = http.Cookie{Name: response1.Cookies()[0].Name, Value: response1.Cookies()[0].Value, Expires: time.Now().Add(111 * time.Second)}
	request5.AddCookie(&cookie)
	cookie = http.Cookie{Name: response1.Cookies()[1].Name, Value: response1.Cookies()[1].Value, Expires: time.Now().Add(111 * time.Second)}
	request5.AddCookie(&cookie)
	response5, _ := client2.Do(request5)
	body, err := ioutil.ReadAll(response5.Body)
	if err != nil {
		log.Fatal(err)
	}
	if string(body) == "{\"result\":{}}" {
		log.Println("未签到")
		info1 := []byte(`{"entity":{"ryid":130266,"xm":"范玉奇","xh":"Y80210177","xydm":"06000000","xy":"信息科学与工程学院","bj":"21计算机专硕班","zy":"电子信息","xb":"男","lxdh":"18255877595","rysf":"研究生","xq":"奉贤校区","ss":"奉贤校区-学生公寓2号楼-707","mph":"707","sfzh":"341221199808087878","jtzz":"安徽省临泉县老集镇老集村","jjlxr":"范中超","jjlxrdh":"18756836135","fdygh":"09061","fdy":"吕奕","swjkzk":"健康","xcm":"是","jkmtp":"","xcmtp":"","hsjcbgtp":"","twsfzc":"是","swdqtw":"","swbz":"","jkmsflm":"是","sfycxxwc":"否","tUustMrybhdgjs":"[]","_ext":"{\"jkmtp\":{\"type\":\"fileUpload\",\"value\":[],\"nameStyle\":\"\"},\"xcmtp\":{\"type\":\"fileUpload\",\"value\":[],\"nameStyle\":\"\"},\"hsjcbgtp\":{\"type\":\"fileUpload\",\"value\":[],\"nameStyle\":\"\"}}","tjsj":"`)
		info2 := []byte(time.Now().Format("2006-01-02 15:04"))
		info3 := []byte(`","tjrq":"`)
		info4 := []byte(time.Now().Format("2006-01-02"))
		info5 := []byte(`","zb":"[]","__type":"sdo:com.sudytech.work.uust.zxxsmryb.zxxsmryb.TUustZxxsmryb"}}`)
		var buffer bytes.Buffer
		buffer.Write(info1)
		buffer.Write(info2)
		buffer.Write(info3)
		buffer.Write(info4)
		buffer.Write(info5)
		request6, err := http.NewRequest("POST", "https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/com.sudytech.work.uust.zxxsmryb.xxsmryb.saveOrupte.biz.ext", bytes.NewBuffer(buffer.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
		request6.Header.Add("User-Agent", agent)
		request6.Header.Add("Content-Type", "text/json")
		request6.Header.Add("Origin", "https://workflow.ecust.edu.cn")
		request6.Header.Add("Referer", "https://workflow.ecust.edu.cn/default/work/uust/zxxsmryb/mrybtb.jsp")
		request6.Header.Add("Sec-Fetch-Site", "same-origin")
		request6.Header.Add("Sec-Fetch-Dest", "empty")
		request6.Header.Add("Sec-Fetch-Mode", "cors")
		request6.Header.Add("X-Requested-With", "XMLHttpRequest")
		cookie = http.Cookie{Name: response3.Cookies()[2].Name, Value: response3.Cookies()[2].Value, Expires: time.Now().Add(111 * time.Second)}
		request6.AddCookie(&cookie)
		cookie = http.Cookie{Name: response1.Cookies()[0].Name, Value: response1.Cookies()[0].Value, Expires: time.Now().Add(111 * time.Second)}
		request6.AddCookie(&cookie)
		cookie = http.Cookie{Name: response1.Cookies()[1].Name, Value: response1.Cookies()[1].Value, Expires: time.Now().Add(111 * time.Second)}
		request6.AddCookie(&cookie)
		response6, _ := client2.Do(request6)
		body, err = ioutil.ReadAll(response6.Body)
		if err != nil {
			log.Fatal(err)
		}
		if string(body) == "{\"result\":{\"result\":\"1\"}}" {
			log.Println("今日签到成功！")
		} else {
			log.Println("签到失败！")
		}
	} else {
		log.Println("今日已签到，无需签到")
	}

}

func init() {
	// 获取日志文件句柄
	// 以 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`./qiandao.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}

// DailyCron 每日指定时间 执行任意个无参任务
// 若今天已经超过了执行时间则等到第二天的指定时间再执行任务
func DailyCron(dateTime time.Time, task func()) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		// 检查是否超过当日的时间
		if next.Sub(now) < 0 {
			next = now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		task()

	}
}
