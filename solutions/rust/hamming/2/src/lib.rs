pub fn hamming_distance(s1: &str, s2: &str) -> Option<usize> {
    (s1.len() == s2.len()).then(|| s1.chars().zip(s2.chars()).filter(|(ch1, ch2)| ch1 != ch2).count())
}