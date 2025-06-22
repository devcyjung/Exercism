pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let rows = minefield.len();
    let cols = minefield.first().map_or(0, |first| first.len());
    let bytes: Vec<_> = minefield.iter().map(|row| row.as_bytes()).collect();
    let count_mines = |i: usize, j: usize| -> u8 {
        (i.saturating_sub(1)..=(i + 1).min(rows - 1))
            .flat_map(|x| (j.saturating_sub(1)..=(j + 1).min(cols - 1)).map(move |y| (x, y)))
            .filter(|&(x, y)| bytes[x][y] == b'*')
            .map(|_| 1)
            .sum()
    };
    bytes
        .iter()
        .enumerate()
        .flat_map(|(i, row)| {
            let u8s = row.iter().enumerate().map(|(j, &ch)| {
                if ch == b'*' {
                    ch
                } else {
                    let mines = count_mines(i, j);
                    if mines == 0 {
                        b' '
                    } else {
                        mines + b'0'
                    }
                }
            });
            String::from_utf8(u8s.collect())
        })
        .collect()
}