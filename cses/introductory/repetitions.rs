use std::collections::HashMap;
use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read input");

    let mut rep_map = HashMap::new();
    for c in input.chars() {
        *rep_map.entry(c).or_insert(0) += 1;
    }

    let mut largest = 0;
    for &v in rep_map.values() {
        if v > largest {
            largest = v;
        }
    }

    println!("{}", largest);
}
