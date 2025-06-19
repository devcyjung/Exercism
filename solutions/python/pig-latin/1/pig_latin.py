from typing import Final



def translate(text: Final[str]) -> str:
    size: Final[int] = len(text)
    def _get_word_begin(cursor: int) -> int:
        return next((
            index for index in range(cursor + 1, size) if not text[index].isspace()
        ), size)
    def _get_word_end(cursor: int) -> int:
        return next((
            index for index in range(cursor + 1, size) if text[index].isspace()
        ), size)
    def _cons_vowel(substr: str) -> tuple[str, str]:
        sep: int = next((
            index for index in range(len(substr)) if substr[index].lower() in 'aeiou'
        ), len(substr))
        return substr[:sep], substr[sep:]
    cursor: int = 0
    next_cursor: int = 0
    result: list[str] = []
    while cursor < size:
        if text[cursor].isspace():
            next_cursor = _get_word_begin(cursor)
            result.append(text[cursor:next_cursor])
            cursor = next_cursor
            continue
        next_cursor = _get_word_end(cursor)
        substr: Final[str] = text[cursor:next_cursor]
        cons, vowel = _cons_vowel(substr)
        if len(cons) == 0 or cons.startswith('xr') or cons.startswith('yt'):
            result.append(substr)
            result.append('ay')
            cursor = next_cursor
            continue
        ypos: int = cons.find('y')
        if ypos > 0:
            result.append(substr[ypos:])
            result.append(substr[:ypos])
            result.append('ay')
            cursor = next_cursor
            continue
        if cons.endswith('q') and vowel.startswith('u'):
            result.append(vowel[1:])
            result.append(substr[:len(cons) + 1])
            result.append('ay')
            cursor = next_cursor
            continue
        result.append(vowel)
        result.append(cons)
        result.append('ay')
        cursor = next_cursor
    return ''.join(result)