#![feature(portable_simd)]

use std::{collections::HashMap, simd::prelude::*, thread};

/// Checks whether the entire input contains only ASCII characters.
fn is_ascii_only(input: &[&str]) -> bool {
    input.iter().all(|s| s.is_ascii())
}

/// Fast SIMD implementation for ASCII-only input
fn ascii_frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let joined: Vec<u8> = input.join("").into_bytes();
    let chunk_size = joined.len()
        .div_ceil(worker_count.max(1))
        .max(1);

    let mut threads = Vec::new();

    for chunk in joined.chunks(chunk_size) {
        let chunk = chunk.to_vec();
        
        threads.push(thread::spawn(move || {
            let mut freq = [0usize; 256];
            let chunks = chunk.chunks_exact(16);
            let remainder = chunks.remainder();
            const UPPER_TO_LOWER: u8 = (b'a').abs_diff(b'A');

            for c in chunks {
                let vec = u8x16::from_slice(c);

                let is_upper = vec.simd_ge(u8x16::splat(b'A')) & vec.simd_le(u8x16::splat(b'Z'));
                let is_lower = vec.simd_ge(u8x16::splat(b'a')) & vec.simd_le(u8x16::splat(b'z'));
                let is_alpha = is_upper | is_lower;

                let mut ascii = vec.to_array();
                let mask = is_upper.to_array();
                let alpha = is_alpha.to_array();

                for i in 0..16 {
                    if mask[i] {
                        ascii[i] += UPPER_TO_LOWER;
                    }
                }

                for i in 0..16 {
                    if alpha[i] {
                        freq[ascii[i] as usize] += 1;
                    }
                }
            }

            for &b in remainder {
                if b.is_ascii_alphabetic() {
                    freq[b.to_ascii_lowercase() as usize] += 1;
                }
            }

            freq
        }));
    }

    let mut combined = [0usize; 256];
    for t in threads {
        let local = t.join().unwrap();
        for (i, &v) in local.iter().enumerate() {
            combined[i] += v;
        }
    }

    combined
        .iter()
        .enumerate()
        .filter_map(|(i, &v)| {
            if v > 0 {
                Some((i as u8 as char, v))
            } else {
                None
            }
        })
        .collect()
}

/// Unicode-aware fallback using `chars()` and `.to_lowercase()`
fn unicode_frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let chunks = input.chunks(
        input.len()
            .div_ceil(worker_count.max(1))
            .max(1)
    );

    let threads: Vec<_> = chunks
        .map(|chunk| {
            let strings: Vec<String> = chunk.iter().map(|s| s.to_string()).collect();
            thread::spawn(move || {
                let mut local_freq = HashMap::new();
                for s in &strings {
                    for ch in s.chars().flat_map(|c| c.to_lowercase()) {
                        if ch.is_alphabetic() {
                            *local_freq.entry(ch).or_insert(0) += 1;
                        }
                    }
                }
                local_freq
            })
        })
        .collect();

    let mut final_freq = HashMap::new();
    for thread in threads {
        for (ch, count) in thread.join().unwrap() {
            *final_freq.entry(ch).or_insert(0) += count;
        }
    }

    final_freq
}

/// Public API: Selects fastest method based on input
pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    if is_ascii_only(input) {
        ascii_frequency(input, worker_count)
    } else {
        unicode_frequency(input, worker_count)
    }
}