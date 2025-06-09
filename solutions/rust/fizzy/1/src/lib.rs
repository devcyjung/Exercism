use std::ops::Rem;

pub struct Matcher<T: ToString + Clone> {
    predicate: Box<dyn Fn(T) -> bool>,
    subs: Box<dyn ToString>,
}

impl<T: ToString + Clone> Matcher<T> {
    pub fn new<F, S>(matcher: F, subs: S) -> Self
    where
        F: Fn(T) -> bool + 'static,
        S: ToString + 'static,
    {
        Self {
            predicate: Box::new(matcher),
            subs: Box::new(subs),
        }
    }
}

pub struct Fizzy<T: ToString + Clone> {
    matchers: Vec<Matcher<T>>,
}

impl<T: ToString + Clone> Fizzy<T> {
    pub fn new() -> Self {
        Self {
            matchers: vec![],
        }
    }

    #[must_use]
    pub fn add_matcher(mut self, matcher: Matcher<T>) -> Self {
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

/// convenience function: return a Fizzy which applies the standard fizz-buzz rules
pub fn fizz_buzz<T>() -> Fizzy<T>
where
    T: ToString + Clone + Rem<T, Output = T> + PartialEq<T> + From<u8>,
    {
    let fizzy = Fizzy::<T>::new();
    let fizzy = fizzy.add_matcher(Matcher::new(|t: T| t % T::from(3) == T::from(0), "fizz"));
    let fizzy = fizzy.add_matcher(Matcher::new(|t: T| t % T::from(5) == T::from(0), "buzz"));
    fizzy
}
