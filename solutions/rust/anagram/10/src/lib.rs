use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let lower_word = word.to_lowercase();
    let word_sorted = to_sorted(lower_word.as_str());
    possible_anagrams.iter()
        .filter(|s| {
            if s.len() != word.len() {
                return false;
            }
            let lower_s = s.to_lowercase();
            if lower_s == lower_word {
                return false;
            }
            to_sorted(lower_s.as_str()) == word_sorted
        })
        .cloned()
        .collect()
}

fn to_sorted(input: &str) -> Vec<char> {
    let mut chars = input.chars().collect::<Vec<_>>();
    chars.sort();
    chars
}