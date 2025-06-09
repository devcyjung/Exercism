pub fn reply(message: &str) -> &str {
    let trimmed = message.trim();
    let is_silent = trimmed.is_empty();
    let is_question = trimmed.ends_with("?");
    let is_yelling = trimmed.chars().any(char::is_alphabetic)
        && trimmed.chars().all(|c| !c.is_alphabetic() || c.is_uppercase());

    if is_silent {
        return "Fine. Be that way!";
    }
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