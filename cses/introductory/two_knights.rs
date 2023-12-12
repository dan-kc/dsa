use std::io;

fn main() {
    let mut n = String::new();
    io::stdin().read_line(&mut n).expect("Failed to read input");
    let n: usize = n.trim().parse().expect("Invalid input");

    println!("0");

    for k in 2..=n {
        // We divide both by 2 due to double counting
        let total = k * k * ((k * k) - 1) / 2;
        let attacks = 4 * 2 * (k - 1) * (k - 2) / 2;
        println!("{}", total - attacks);
    }
}
