import string

def is_pangram(sentence):
    charset = set(sentence.lower())
    return len(charset) >= 26 and all((ch in charset for ch in string.ascii_lowercase))