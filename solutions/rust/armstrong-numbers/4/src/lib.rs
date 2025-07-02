pub fn is_armstrong_number(num: u32) -> bool {
    std::iter::successors(
        Some((
            num.checked_ilog10().unwrap_or(0) + 1, num, 0
        )), |&(len, rem, acc)| (rem != 0).then(|| (
            len, rem / 10, acc + (rem % 10).pow(len)
        ))
    ).last().unwrap_or((0,0,0)).2 == num
}