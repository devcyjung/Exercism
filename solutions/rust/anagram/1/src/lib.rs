use std::collections::{HashMap,HashSet};

pub fn anagrams_for<'a>(word: &'a str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut anagrams = HashSet::new();
    let target_lower = word.to_lowercase();
    let target = target_lower.as_str();
    let target_map = get_anagram_map(target);
    for &cand in possible_anagrams {
        let source_lower = cand.to_lowercase();
        let source = source_lower.as_str();
        if target == source {
            continue
        }
        if target_map != get_anagram_map(source) {
            continue
        }
        anagrams.insert(cand);
    }
    anagrams
}

fn get_anagram_map<'a>(word: &'a str) -> HashMap<char, u64> {
    let mut m = HashMap::new();
    for ch in word.chars() {
        *m.entry(ch).or_insert(0u64) += 1;
    }
    m
}