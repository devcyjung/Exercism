pub fn encrypt(input: &str) -> String {
    let text: String = input.chars()
        .filter(|c| c.is_alphanumeric())
        .flat_map(char::to_lowercase).collect();
    let size = text.len();
    let mut r = size.isqrt();
    let mut c = r;
    let mut change_r = false;
    while r * c < size {
        if change_r {
            r += 1
        } else {
            c += 1
        }
        change_r = !change_r
    }
    let bytes = text.as_bytes();
    (0..c).flat_map(|ci| {
        let bytes = &bytes;
        (0..r).map(move |ri| {
            bytes.get(ri*c+ci).map_or(' ', |b| *b as char)
        }).chain(
            if ci != c - 1 {
                " ".chars()
            } else {
                "".chars()
            }
        )
    }).collect()
}