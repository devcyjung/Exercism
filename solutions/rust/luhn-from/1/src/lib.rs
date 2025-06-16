pub struct Luhn(String);

impl Luhn {
    pub fn is_valid(&self) -> bool {
        let code = &self.0;
        if code.chars().any(|ch| !ch.is_ascii_digit() && ch != ' ') {
            return false;
        }
        let trimmed = code.chars().filter(char::is_ascii_digit).collect::<Vec<_>>();
        if trimmed.len() <= 1 {
            return false;
        }
        trimmed.iter().filter_map(|ch| ch.to_digit(10))
            .rfold((false, 0), |(double, sum), num| (
                !double,
                if double {
                    let mut doubled = num * 2;
                    if doubled > 9 {
                        doubled -= 9;
                    }
                    sum + doubled
                } else { sum + num }
            )).1 % 10 == 0
    }
}

impl<T> From<T> for Luhn
where
    T: ToString
{
    fn from(input: T) -> Self {
        Self(input.to_string())
    }
}