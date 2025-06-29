pub fn series(digits: &str, len: usize) -> Vec<String> {
    (len != 0).then(|| Vec::from_iter(digits.chars())
        .windows(len)
        .map(String::from_iter)
        .collect()
    ).unwrap_or_else(Vec::new)
}