use once_cell::sync::Lazy;
use std::sync::Mutex;

static PRIMES: Lazy<Mutex<Vec<u32>>> = Lazy::new(|| Mutex::new(vec![2, 3]));

pub fn nth(n: u32) -> u32 {
    let mut primes = PRIMES.lock().expect("lock error");
    let n: usize = n.try_into().expect("u32 -> usize error");
    if n < primes.len() {
        return primes[n];
    }
    for i in (primes.last().expect("PRIMES is empty").saturating_add(1)..) {
        let is_prime = primes.iter().all(|p| i % p != 0);
        if is_prime {
            primes.push(i);
        }
        if primes.len() == n + 1 {
            break;
        }
    }
    primes.last().copied().expect("PRIMES is empty")
}