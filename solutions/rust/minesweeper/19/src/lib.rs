pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let rows = minefield.len();
    let cols = minefield.first().map_or(0, |first| first.len());
    let bytes = minefield.iter().map(|row| row.as_bytes()).collect::<Vec<_>>();
    let count_mines = |i: usize, j: usize| {
        (i.saturating_sub(1)..=(i + 1).min(rows - 1))
            .flat_map(|x| (j.saturating_sub(1)..=(j + 1).min(cols - 1))
                .map(move |y| (x, y))
            )
            .filter_map(|(x, y)| (bytes[x][y] == b'*').then_some(1))
            .sum()
    };
    minefield.iter().enumerate().map(|(i, row)| {
        row.chars().enumerate().flat_map(|(j, ch)| {
            (ch == '*')
                .then_some(ch)
                .or_else(|| {
                    let mines = count_mines(i, j);
                    (mines == 0)
                        .then_some(' ')
                        .or_else(|| char::from_digit(mines, 10))
                })
        }).collect()
    }).collect()
}