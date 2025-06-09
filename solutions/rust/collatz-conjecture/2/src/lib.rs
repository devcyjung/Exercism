pub fn collatz(n: u64) -> Option<u64> {
    match n {
        0 => None,
        1 => Some(0),
        _ => {
            if n % 2 == 0 {
                return Some(collatz(n / 2)? + 1);
            }
            Some(collatz(n * 3 + 1)? + 1)
        }
    }
}
