use std::cmp;

#[derive(Debug)]
pub struct Item {
    pub weight: usize,
    pub value: usize,
}

pub fn maximum_value(max_weight: usize, items: &[Item]) -> usize {
    let size= items.len();
    let mut memo = vec![vec![0; max_weight + 1]; size + 1];
    for i in 0..size {
        let item_weight = items[i].weight;
        let item_value = items[i].value;
        for j in 1..=max_weight {
            memo[i + 1][j] = if item_weight > j {
                memo[i][j]
            } else {
                cmp::max(memo[i][j], memo[i][j - item_weight] + item_value)
            };
        }
    }
    memo[size][max_weight]
}