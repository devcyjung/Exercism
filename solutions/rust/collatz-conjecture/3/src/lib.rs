pub fn collatz(n: u64) -> Option<u64> {
    (n > 0).then(|| {
        std::iter::successors(
            Some((n, 0)),
            |&(num, acc)| (num & 1 ==0)
                .then(|| u64::from(num.trailing_zeros()))
                .map_or_else(
                    || Some((num + 1 + (num << 1), acc + 1)),
                    |tz| Some((num >> tz, acc + tz))
                )
        )
        .find(|&(num, _acc)| num == 1)
        .map_or_else(|| 0, |(_num, acc)| acc)
    })
}