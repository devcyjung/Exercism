pub fn reply(message: &str) -> &str {
    let trimmed = message.trim();

    if trimmed.is_empty() {
        return "Fine. Be that way!";
    }
    
    let is_question = trimmed.ends_with("?");
    let is_yelling = trimmed.chars().any(char::is_alphabetic) && !trimmed.chars().any(char::is_lowercase);

    
    if is_yelling && is_question {
        return "Calm down, I know what I'm doing!";
    }
    if is_yelling {
        return "Whoa, chill out!";
    }
    if is_question {
        return "Sure.";
    }
    "Whatever."
}