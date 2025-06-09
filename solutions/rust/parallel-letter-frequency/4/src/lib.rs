use std::{
    collections::HashMap,
    thread,
    sync::mpsc,
    iter,
};

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let (tx, rx) = mpsc::channel();
    let worker_count = worker_count.max(1);
    let chunk_size = input.len().div_ceil(worker_count);
    thread::scope(|scope| {
        for (chunk, tx) in input.chunks(chunk_size.max(1)).zip(iter::repeat(tx)) {
            scope.spawn(move || {
                for text in chunk {
                    tx.send(frequency_map(text.to_lowercase())).expect("send failed"); 
                }
            });
        }
    });
    rx.iter().fold(Default::default(), |mut acc, mut res| {
        for (k, v) in res.drain() {
            *acc.entry(k).or_insert(0) += v;
        }
        acc
    })
}

fn frequency_map(text: String) -> HashMap<char, usize> {
    text.chars().fold(Default::default(), |mut acc, ch| {
        if ch.is_alphabetic() {
            *acc.entry(ch).or_insert(0) += 1;
        }
        acc
    })
}