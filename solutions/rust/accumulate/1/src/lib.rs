pub fn map<I, F, R>(input: I, function: F) -> Vec<R> 
where
    I: IntoIterator,
    F: FnMut(<I as IntoIterator>::Item) -> R,
{
    input.into_iter().map(function).collect()
}