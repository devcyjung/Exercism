use std::collections::{HashMap,HashSet};

pub fn anagrams_for<'a>(word: &'a str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_lowercase();
    let word_map = get_wc_map(&word);
    possible_anagrams
        .iter()
        .filter(|&&elem| {
            let elem = elem.to_lowercase();
            elem != word && get_wc_map(&elem) == word_map
        })
        .copied()
        .collect()
}

fn get_wc_map(word: &str) -> HashMap<char, u64> {
    let mut wc = HashMap::new();
    for ch in word.chars() {
        *wc.entry(ch).or_insert(0u64) += 1;
    }
    wc
}