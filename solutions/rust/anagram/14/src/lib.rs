use std::collections::HashSet;

pub fn anagrams_for<'a>(key: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let lower_key = key.to_lowercase();
    let sorted_key = to_sorted(&lower_key);
    possible_anagrams.iter()
        .filter_map(|&word| {
            (word.len() == lower_key.len())
                .then(|| word.to_lowercase())
                .filter(|lower_word| *lower_word != lower_key)
                .map(|lower_word| to_sorted(&lower_word))
                .filter(|sorted_word| *sorted_word == sorted_key)
                .map(|_| word)
        })
        .collect()
}

fn to_sorted<T>(input: T) -> String
where
    T: AsRef<str>
{
    let mut chars = input.as_ref().chars().collect::<Vec<_>>();
    chars.sort();
    chars.iter().collect()
}