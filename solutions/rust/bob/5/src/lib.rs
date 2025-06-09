pub fn reply(message: &str) -> &str {
    let trimmed = message.trim();

    if trimmed.is_empty() {
        return "Fine. Be that way!";
    }
    
    let is_question = trimmed.ends_with("?");

    let mut only_alphabets = trimmed.chars().filter(|c| c.is_alphabetic()).peekable();
    
    let is_yelling = only_alphabets.peek().is_some() && only_alphabets.all(char::is_uppercase);

    match (is_yelling, is_question) {
        (false, true) => "Sure.",
        (true, false) => "Whoa, chill out!",
        (true, true) => "Calm down, I know what I'm doing!",
        (false, false) => "Whatever.",
    }
}