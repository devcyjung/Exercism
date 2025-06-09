use std::collections::{HashMap,HashSet};

pub fn anagrams_for<'a>(word: &'a str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_uppercase();
    let word_len = word.len();
    let word_map = get_wc_map(&word);
    possible_anagrams
        .iter()
        .filter(|&elem| {
            let elem = elem.to_uppercase();
            word_len == elem.len()
                && word != elem
                && get_wc_map(&elem) == word_map
        })
        .copied()
        .collect()
}

fn get_wc_map(word: &str) -> HashMap<char, usize> {
    word.chars().fold(HashMap::new(), |mut acc, c| {
        acc.entry(c).and_modify(|e| { *e += 1 }).or_insert(1);
        acc
    })
}