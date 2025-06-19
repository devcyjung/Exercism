def is_valid(isbn):
    input = [
        ord(ch) - ord('0') if ch.isdigit() else \
        10 if i == 9 and ch.upper() == 'X' else -1 \
        for i, ch in enumerate((
            ch for ch in isbn if ch != '-'
        ))
    ]
    return len(input) == 10 and all((n >= 0 for n in input)) and sum((
        x * y for x, y in zip(input, range(10, 0, -1))
    )) % 11 == 0