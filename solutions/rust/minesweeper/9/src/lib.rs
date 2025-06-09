/// # Panics
/// When each rows of `minefield` don't have equal length.
/// When the result is in unreachable state. i.e. some grid has 9 or more neighboring mines
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

    let byte_minefield = minefield.iter().map(|row| row.as_bytes()).collect::<Vec<_>>();

    (0..row_size)
        .map(|i| {
            (0..col_size)
                .map(|j| {
                    if byte_minefield[i][j] == b'*' {
                        return '*';
                    }
                    let mine_cnt_result = (i.saturating_sub(1)..=i + 1)
                        .map(|x| {
                            (j.saturating_sub(1)..=j + 1)
                                .filter(|&y| {
                                    !(x == i && y == j)
                                    && x < row_size
                                    && y < col_size
                                    && byte_minefield[x][y] == b'*'
                                })
                                .fold(0, |acc, _| acc + 1)
                        })
                        .sum();
                    match mine_cnt_result {
                        0 => ' ',
                        byt @ 1..=8 => char::from_digit(byt, 10).unwrap(),
                        _ => unreachable!("a grid should only have 0 - 8 mines nearby")
                    }
                })
                .collect()
        })
        .collect()
}
