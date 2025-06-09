use std::cmp;

pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let row_len = input.len();
    let col_len = input[0].len();
    let mut row_maxs: Vec<u64> = vec![0; row_len];
    let mut col_mins: Vec<u64> = vec![u64::MAX; col_len];
    (0..row_len).for_each(|row| {
        (0..col_len).for_each(|col| {
            row_maxs[row] = cmp::max(row_maxs[row], input[row][col]);
            col_mins[col] = cmp::min(col_mins[col], input[row][col]);
        })
    });
    (0..row_len).flat_map(|row| {
        let input = &input;
        let row_maxs = &row_maxs;
        let col_mins = &col_mins;
        (0..col_len)
            .filter(move |&col| {
                row_maxs[row] == input[row][col] && col_mins[col] == input[row][col]
            })
            .map(move |col| (row, col))
    }).collect()
}
