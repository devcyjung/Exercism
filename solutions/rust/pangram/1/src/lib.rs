pub fn is_pangram(sentence: &str) -> bool {
    sentence.chars()
        .fold(vec![false; 26], |mut acc, ch| {
            if ch.is_ascii_alphabetic() {
                let idx: usize = (
                    ch
                        .to_digit(36)
                        .expect("ascii alphabets are always a digit") - 10
                )
                    .try_into()
                    .expect("u32 should fit into usize");
                acc[idx] = true;
            }
            acc
        })
        .iter()
        .all(|&b| b)
}
