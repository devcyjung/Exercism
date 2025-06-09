pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }
    if n == 1 {
        return Some(0);
    }
    if n % 2 == 0 {
        return Some(collatz(n / 2)? + 1);
    }
    Some(collatz(3 * n + 1)? + 1)
}
