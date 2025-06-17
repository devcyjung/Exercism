use std::collections::HashSet;

pub fn anagrams_for<'a>(key: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let lower_key = key.to_lowercase();
    let sorted_key = to_sorted(lower_key.as_str());
    possible_anagrams.iter()
        .filter_map(|&word| {
            let lower_word = word.to_lowercase();
            (lower_word.len() == lower_key.len() && lower_word != lower_key)
                .then(|| to_sorted(lower_word.as_str()))
                .filter(|sorted_word| *sorted_word == sorted_key)
                .map(|_| word)
        })
        .collect()
}

fn to_sorted(input: &str) -> String {
    let mut chars = input.chars().collect::<Vec<_>>();
    chars.sort();
    chars.iter().collect()
}