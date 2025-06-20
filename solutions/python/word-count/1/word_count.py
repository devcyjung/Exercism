from collections import Counter
from string import ascii_letters, ascii_lowercase, whitespace, punctuation

_non_words = whitespace + punctuation.replace("'", "")
_trimmer = str.maketrans(
    ascii_letters + _non_words, ascii_lowercase * 2 + " " * len(_non_words),
)

def count_words(sentence):
    return Counter(
        word for word in (
            trimmed_except_apos.strip("'")
            for trimmed_except_apos in
            sentence.translate(_trimmer).split()
        ) if word
    )