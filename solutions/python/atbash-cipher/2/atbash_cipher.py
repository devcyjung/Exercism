from string import ascii_letters, ascii_lowercase, digits, punctuation, whitespace

def encode(plain_text):
    encoded = plain_text.translate(_encoding)
    return ' '.join(encoded[i:i+5] for i in range(0, len(encoded), 5))

def decode(ciphered_text):
    return ciphered_text.translate(_decoding)

_encoding = str.maketrans(
    ascii_letters + digits,
    ascii_lowercase[::-1] * 2 + digits,
    punctuation + whitespace
)

_decoding = str.maketrans(
    ascii_lowercase + digits,
    ascii_lowercase[::-1] + digits,
    punctuation + whitespace
)