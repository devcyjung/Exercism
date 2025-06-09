pub fn encrypt(input: &str) -> String {
    let sanitized: String = input.chars()
        .filter(|c| c.is_alphanumeric())
        .flat_map(char::to_lowercase).collect();
    let size = sanitized.len();
    let root = size.isqrt();
    let (r, c) = if root * root == size {
        (root, root)
    } else if root * (root + 1) >= size {
        (root, root + 1)
    } else {
        (root + 1, root + 1)
    };
    let sanitized = sanitized.as_bytes();
    (0..c).flat_map(|ci| {
        let sanitized = &sanitized;
        (0..r).map(move |ri| {
            sanitized.get(ri*c+ci).map_or(' ', |&b| b as char)
        }).chain(if ci != c - 1 { " ".chars() } else { "".chars() })
    }).collect()
}