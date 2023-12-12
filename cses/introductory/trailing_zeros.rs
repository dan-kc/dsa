use std::io;
fn main() {
    let mut n = String::new();
    io::stdin().read_line(&mut n).expect("could not read line");

    let mut n: usize = n.trim().parse().expect("could not parse");
    let mut ans = 0;

    while n > 0 {
        n /= 5;
        ans += n;
    }
    println!("{}", ans)
}
