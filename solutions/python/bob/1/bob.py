"""bob"""
def response(hey_bob):
    trimmed = hey_bob.rstrip()
    yelling = hey_bob.isupper()
    silent = trimmed == ""
    asking = trimmed.endswith("?")

    if asking and yelling:
        return "Calm down, I know what I'm doing!"
    if asking:
        return "Sure."
    if yelling:
        return "Whoa, chill out!"
    if silent:
        return "Fine. Be that way!"
    return "Whatever."