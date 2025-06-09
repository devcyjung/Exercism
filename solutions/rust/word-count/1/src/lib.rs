use std::collections::HashMap;
use regex::Regex;

/// Count occurrences of words.
pub fn word_count(words: &str) -> HashMap<String, u32> {
    Regex::new(r"[^A-Za-z0-9']")
        .unwrap()
        .split(words)
        .map(|w| {
            w.trim_start_matches(|c: char| !c.is_alphanumeric())
                .trim_end_matches(|c: char| !c.is_alphanumeric())
                .to_lowercase()
        })
        .filter(|w| !w.is_empty())
        .fold(HashMap::new(), |mut acc, w| {
            acc.entry(w).and_modify(|e| *e += 1).or_insert(1);
            acc
        })
}
