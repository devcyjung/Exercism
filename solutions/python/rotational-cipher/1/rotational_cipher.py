def rotate(text, key):
    return ''.join((
        chr((ord(ch) - ord('A') + key) % 26 + ord('A'))
        if ch.isupper() else
        chr((ord(ch) - ord('a') + key) % 26 + ord('a'))
        if ch.islower() else ch
        for ch in text
    ))