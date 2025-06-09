pub fn encode(source: &str) -> String {
    source.chars()
        .fold(Vec::<(char, usize)>::new(), |mut acc, ch| {
            match acc.last_mut() {
                Some(last) if last.0 == ch => last.1 += 1,
                _ => acc.push((ch, 1)),
            };
            acc
        })
        .into_iter()
        .map(|(ch, count)| {
            if count == 1 {
                ch.to_string()
            } else {
                count.to_string() + &ch.to_string()
            }
        })
        .collect()
}

pub fn decode(source: &str) -> String {
    source.chars()
        .fold((String::new(), Vec::<(char, usize)>::new()), |mut acc, ch| {
            if ch.is_ascii_digit() {
                acc.0.push(ch);
            } else {
                let repetition: usize = if acc.0.is_empty() {
                    1
                } else {
                    acc.0.parse::<usize>().expect("digits are always decimal")
                };
                acc.0.clear();
                acc.1.push((ch, repetition));
            }
            acc
        })
        .1
        .into_iter()
        .map(|(ch, size)| {
            ch.to_string().repeat(size)
        })
        .collect()
}
