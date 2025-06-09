pub fn encrypt(input: &str) -> String {
    let sanitized: Vec<_> = input.chars()
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
    (0..c).flat_map(|ci| {
        let sanitized = &sanitized;
        (0..r).map(move |ri| {
            sanitized.get(ri*c+ci).copied().unwrap_or(' ')
        }).chain(if ci != c - 1 { Some(' ') } else { None })
    }).collect()
}