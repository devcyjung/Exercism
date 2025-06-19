pub fn collatz(n: u64) -> Option<u64> {
    (n > 0).then(|| {
        std::iter::successors(
            Some((n, 0u64)),
            |&(num, acc)| (num & 1 ==0)
                .then(|| u64::from(num.trailing_zeros()))
                .map_or_else(
                    || Some((num + 1 + (num << 1), acc.saturating_add(1))),
                    |tz| Some((num >> tz, acc.saturating_add(tz)))
                )
        )
        .find(|&(num, acc)| num == 1 || acc == u64::MAX)
        .map_or_else(|| u64::MAX, |(_num, acc)| acc)
    })
}