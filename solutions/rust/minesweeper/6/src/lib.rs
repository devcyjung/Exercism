/// # Panics
/// When each rows of `minefield` don't have equal length, this function panics.
/// Or when parsing the annotation result (u8 slices) into utf-8 strings fail, it panics.
#[must_use]
pub fn annotate(minefield: &[&str]) -> Vec<String> {
    if minefield.is_empty() {
        return vec![];
    }

    if minefield[0].is_empty() {
        return vec![String::new()];
    }

    let row_size = minefield.len();
    let col_size = minefield[0].len();

    for s in minefield {
        if s.len() != col_size {
            unreachable!("input length is not equal")
        }
    }

    (0..row_size)
        .map(|i| {
            (0..col_size)
                .map(|j| {
                    if matches!(minefield[i].as_bytes()[j], b'*') {
                        return '*';
                    }
                    let mine_cnt_result = u32::try_from(
                        (i.saturating_sub(1)..=i + 1)
                            .map(|x| {
                                (j.saturating_sub(1)..=j + 1)
                                    .filter(|&y| {
                                        !(x == i && y == j)
                                            && x < row_size
                                            && y < col_size
                                            && matches!(minefield[x].as_bytes()[y], b'*')
                                    })
                                    .count()
                            })
                            .sum::<usize>(),
                    )
                    .map(|u| char::from_digit(u, 10));
                    match mine_cnt_result {
                        Err(e) => unreachable!("usize -> u32 conversion error, {}", e),
                        Ok(Some('0')) => ' ',
                        Ok(Some(c)) if ('1'..='8').contains(&c) => c,
                        Ok(Some(c)) => unreachable!("unexpected character {}", c),
                        _ => unreachable!("unexpected non digit count"),
                    }
                })
                .collect()
        })
        .collect()
}
