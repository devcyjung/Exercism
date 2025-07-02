pub fn is_armstrong_number(num: u32) -> bool {
    std::iter::successors(
        Some((
            num.checked_ilog10().unwrap_or(0u32).saturating_add(1u32), num, 0u32
        )), |&(len, rem, acc)| (rem.ne(&0)).then(|| (
            len, rem.div_euclid(10), acc.saturating_add(rem.rem_euclid(10).pow(len))
        ))
    ).last().unwrap_or((0,0,0)).2.eq(&num)
}