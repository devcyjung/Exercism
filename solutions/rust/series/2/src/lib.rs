pub fn series(digits: &str, len: usize) -> Vec<String> {
    digits
        .as_bytes()
        .windows(len)
        .flat_map(std::str::from_utf8)
        .map(String::from)
        .collect()
}