/// # Panics
/// When each rows of `minefield` don't have equal length, this function panics.
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
                    let mut mine_cnt = 0;
                    (if i == 0 { 0 } else { i - 1 }..=i + 1).for_each(|x| {
                        (if j == 0 { 0 } else { j - 1 }..=j + 1).for_each(|y| {
                            if !(x == i && y == j)
                                && x < row_size
                                && y < col_size
                                && matches!(minefield[x].as_bytes()[y], b'*')
                            {
                                mine_cnt += 1;
                            }
                        });
                    });
                    if mine_cnt == 0 {
                        ' '
                    } else {
                        char::from(mine_cnt + b'0')
                    }
                })
                .collect()
        })
        .collect()
}
