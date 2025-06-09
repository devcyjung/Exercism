pub fn brackets_are_balanced(string: &str) -> bool {
    string
        .chars()
        .filter_map(|c| match c {
            '(' => Some(1), '[' => Some(2), '{' => Some(3),
            ')' => Some(-1), ']' => Some(-2), '}' => Some(-3),
            _ => None, 
        })
        .try_fold(vec![], |mut acc, n| {
            if n > 0 {
                acc.push(n);
                Some(acc)
            } else {
                let last = acc.last()?;
                if last + n == 0 {
                    acc.pop();
                    Some(acc)
                } else {
                    None
                }
            }
        })
        .is_some_and(|vec| vec.is_empty())
}