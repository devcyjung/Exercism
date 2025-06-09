use std::{
    collections::HashMap,
    iter,
    sync::mpsc::{
        self,
        SendError,
    },
    thread,
};

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let (tx, rx) = mpsc::channel();
    let chunk_size = input.len().div_ceil(worker_count.max(1)).max(1);
    thread::scope(|scope| {
        for (chunk, tx) in input.chunks(chunk_size).zip(iter::repeat(tx)) {
            scope.spawn(move || {
                for text in chunk {
                    tx.send(frequency_map(text)).unwrap_or_else(|SendError(failed)| {
                        eprintln!("Failed to send {failed:#?} because receiver is dropped.");
                    }); 
                }
            });
        }
        let mut acc = HashMap::default();
        for res in rx {
            for (k, v) in res {
                *acc.entry(k).or_default() += v;
            }
        }
        acc
    })
}

fn frequency_map(text: &str) -> HashMap<char, usize> {
    let mut acc = HashMap::default();
    for ch in text.to_lowercase().chars().filter(|ch| ch.is_alphabetic()) {
        *acc.entry(ch).or_default() += 1;
    }
    acc
}