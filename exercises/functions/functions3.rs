// functions3.rs
// Make me compile! Execute `rustlings hint functions3` for hints :)


fn main() {
    let x = 3;
    call_me(x);
}

fn call_me(num: u32) {
    for i in 0..num {
        println!("Ring! Call number {}", i + 1);
    }
}
