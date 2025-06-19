from collections.abc import Callable
from typing import Final

def translate(text: Final[str]) -> str:
    subtext: str = text[:]
    result: list[str] = []
    while subtext:
        if subtext[0].isspace():
            spaces, subtext = _space_word(subtext)
            result.append(spaces)
            continue
        substr, subtext = _word_space(subtext)
        cons, vowel = _cons_vowel(substr)
        if not cons or cons.startswith('xr') or cons.startswith('yt'):
            result.extend((substr, 'ay'))
            continue
        ypos: Final[int] = cons.find('y')
        if ypos > 0:
            result.extend((substr[ypos:], substr[:ypos], 'ay'))
            continue
        if cons.endswith('q') and vowel.startswith('u'):
            result.extend((vowel[1:], cons, vowel[:1], 'ay'))
            continue
        else:
            result.extend((vowel, cons, 'ay'))
    return ''.join(result)

def _sep_util(substr: str, sep_at: Callable[[int], bool]) -> tuple[str, str]:
    last: int = len(substr)
    sep: int = next((pos for pos in range(last) if sep_at(pos)), last)
    return substr[:sep], substr[sep:]

def _space_word(substr: str) -> tuple[str, str]:
    return _sep_util(substr, lambda pos: not substr[pos].isspace())

def _word_space(substr: str) -> tuple[str, str]:
    return _sep_util(substr, lambda pos: substr[pos].isspace())
    
def _cons_vowel(substr: str) -> tuple[str, str]:
    return _sep_util(substr, lambda pos: substr[pos].lower() in 'aeiou')