use std::ops::Rem;

pub struct Matcher<'a, T: ToString + Clone> {
    predicate: Box<dyn Fn(T) -> bool + 'a>,
    subs: Box<dyn ToString + 'a>,
}

impl<'a, T: ToString + Clone> Matcher<'a, T> {
    pub fn new<F, S>(matcher: F, subs: S) -> Self
    where
        F: Fn(T) -> bool + 'a,
        S: ToString + 'a,
    {
        Self {
            predicate: Box::new(matcher),
            subs: Box::new(subs),
        }
    }
}

pub struct Fizzy<'a, T: ToString + Clone> {
    matchers: Vec<Matcher<'a, T>>,
}

impl<'a, T: ToString + Clone> Fizzy<'a, T> {
    pub fn new() -> Self {
        Self {
            matchers: vec![],
        }
    }

    #[must_use]
    pub fn add_matcher(mut self, matcher: Matcher<'a, T>) -> Self {
        self.matchers.push(matcher);
        self
    }

    pub fn apply<I>(self, iter: I) -> impl Iterator<Item = String>
    where
        I: Iterator<Item = T>,
    {
        iter.map(move |t| {
            let sub_word = self.matchers
                .iter()
                .filter(|matcher| (matcher.predicate)(t.clone()))
                .map(|matcher| matcher.subs.to_string())
                .collect::<Vec<_>>()
                .join("");
            if sub_word.is_empty() {
                t.to_string()
            } else {
                sub_word
            }
        })
    }
}

pub fn fizz_buzz<'a, T>() -> Fizzy<'a, T>
where
    T: ToString + Clone + Rem<T, Output = T> + PartialEq<T> + From<u8>,
{
    Fizzy::<T>::new()
        .add_matcher(Matcher::new(|t: T| t % T::from(3) == T::from(0), "fizz"))
        .add_matcher(Matcher::new(|t: T| t % T::from(5) == T::from(0), "buzz"))
}
