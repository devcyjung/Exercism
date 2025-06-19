from string import ascii_letters, ascii_lowercase, digits

def encode(plain_text):
    encoded = plain_text.translate(_encoding)
    return ' '.join(encoded[i:i+5] for i in range(0, len(encoded), 5))

def decode(ciphered_text):
    return ciphered_text.translate(_decoding)

_encoding = str.maketrans(
    ascii_letters + digits,
    ascii_lowercase[::-1] * 2 + digits,
    ''.join(
        chr(codepoint) for codepoint in range(0x110000)
        if chr(codepoint) not in ascii_letters + digits
    )
)

_decoding = str.maketrans(
    ascii_lowercase + digits,
    ascii_lowercase[::-1] + digits,
    ''.join(
        chr(codepoint) for codepoint in range(0x110000)
        if chr(codepoint) not in ascii_lowercase + digits
    )
)