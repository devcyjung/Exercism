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
    (0..c).flat_map(|ci| {
        let text = &text;
        (0..r).map(move |ri| {
            text.chars().nth(ri*c+ci).unwrap_or(' ')
        }).chain(
            if ci != c - 1 {
                " ".chars()
            } else {
                "".chars()
            }
        )
    }).collect()
}