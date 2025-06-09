#![feature(inherent_str_constructors)]
/// # Panics
/// When each rows of `minefield` don't have equal length, this function panics.
#[must_use] pub fn annotate(minefield: &[&str]) -> Vec<String> {
    if minefield.is_empty() {
        return vec![];
    }
    
    if minefield[0].is_empty() {
        return vec![String::from("")];
    }

    let row_size = minefield.len();
    let col_size = minefield[0].len();

    for s in minefield {
        if s.len() != col_size {
            unreachable!("input length is not equal")
        }
    }

    let mut result = vec![b'0'; row_size * col_size];

    (0..row_size).for_each(|i| {
        (0..col_size).for_each(|j| {
            if matches!(minefield[i].as_bytes()[j], b'*') {
                result[i * col_size + j] = b'*';
            } else {
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
                result[i * col_size + j] = if mine_cnt == 0 { b' ' } else { mine_cnt + b'0' };
            }
        });
    });

    (0..row_size)
        .map(|i| {
            String::from(
                str::from_utf8(&result[i * col_size..(i + 1) * col_size])
                    .expect("Invalid UTF-8"),
            )
        })
        .collect::<Vec<String>>()
}
