pub fn abbreviate(phrase: &str) -> String {
    let mut peek_iter = phrase.chars().peekable();
    let mut take_next = true;
    let mut result = String::with_capacity(phrase.len());
    let mut prev_char: Option<char> = None;
    while let Some(ch) = peek_iter.next() {
        if !ch.is_alphanumeric() {
            take_next = true;
            if let (Some(&next_ch), Some(prev_ch)) = (peek_iter.peek(), prev_char) {
                if ch == '\'' && next_ch.is_alphanumeric() && prev_ch.is_alphanumeric() {
                    take_next = false;
                }
            }
        } else if !take_next {
            if let Some(&next_ch) = peek_iter.peek() {
                if ch.is_lowercase() && next_ch.is_uppercase() {
                    take_next = true;
                }
            }
        } else {
            result.push(ch);
            take_next = false;
        }
        prev_char = Some(ch);
    }
    result.to_uppercase()
}