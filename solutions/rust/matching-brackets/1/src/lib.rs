pub fn brackets_are_balanced(string: &str) -> bool {
    string
        .chars()
        .filter(|c| matches!(c, '(' | '[' | '{' | ')' | ']' | '}'))
        .map(|c| match c {
            '(' => 1, '[' => 2, '{' => 3, ')' => -1, ']' => -2, '}' => -3, _ => unreachable!(), 
        })
        .try_fold(vec![], |mut acc, n| {
            if (0..).contains(&n) {
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
        .map_or(false, |vec| vec.is_empty())
}