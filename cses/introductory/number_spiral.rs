use std::io;

fn main() {
    let mut n = String::new();
    io::stdin().read_line(&mut n).expect("couldn't read line");
    let n: usize = n.trim().parse().expect("couldn't parse");
    for _ in 0..n {
        let mut x_y = String::new();
        io::stdin().read_line(&mut x_y).expect("Couldn't read line");
        let mut iter = x_y.split_whitespace();
        let y: usize = iter.next().unwrap().parse().expect("couldn't parse");
        let x: usize = iter.next().unwrap().parse().expect("couldn't parse");
        let val: usize;
        let max = y.max(x);
        if max % 2 == 0 {
            if y < x {
                val = (max - 1) * (max - 1) + y;
            } else {
                val = (max - 1) * (max - 1) + y + (max - x);
            }
        } else {
            if y < x {
                val = (max - 1) * (max - 1) + x + (max - y);
            } else {
                val = (max - 1) * (max - 1) + x;
            }
        }
        println!("{}", val)
    }
}
