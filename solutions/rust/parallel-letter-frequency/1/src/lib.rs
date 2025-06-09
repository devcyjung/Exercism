use std::{
    collections::HashMap,
    thread,
    sync::mpsc::{
        channel,
        Sender,
    },
};

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let (tx, rx) = channel();
    let worker_count = worker_count.max(1);
    let chunk_size = input.len().div_ceil(worker_count);
    input.chunks(chunk_size.max(1)).for_each(|chunk| {
        let tx = Sender::clone(&tx);
        let chunk : Vec<_> = chunk.iter().copied().map(str::to_lowercase).collect();
        thread::spawn(move || {
            chunk.iter().for_each(|text| {
                tx.send(frequency_map(text)).expect("send failed"); 
            });
        });
    });
    drop(tx);
    rx.iter().fold(HashMap::new(), |mut acc, res| {
        res.iter().for_each(|(&k, &v)| {
            acc.entry(k).and_modify(|e| { *e += v }).or_insert(v);
        });
        acc
    })
}

fn frequency_map(text: &str) -> HashMap<char, usize> {
    text.chars().fold(HashMap::new(), |mut acc, ch| {
        if ch.is_alphabetic() {
            acc.entry(ch).and_modify(|e| { *e += 1 }).or_insert(1);
        }
        acc
    })
}