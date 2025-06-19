from typing import Final

def _space_word(substr: str) -> tuple[str, str]:
    sep: Final[int] = next((
        index for index in range(len(substr)) if not substr[index].isspace()
    ), len(substr))
    return substr[:sep], substr[sep:]

def _word_space(substr: str) -> tuple[str, str]:
    sep: Final[int] = next((
        index for index in range(len(substr)) if substr[index].isspace()
    ), len(substr))
    return substr[:sep], substr[sep:]
    
def _cons_vowel(substr: str) -> tuple[str, str]:
    sep: Final[int] = next((
        index for index in range(len(substr)) if substr[index].lower() in 'aeiou'
    ), len(substr))
    return substr[:sep], substr[sep:]

def translate(text: Final[str]) -> str:
    subtext: str = text[:]
    result: list[str] = []
    while len(subtext) > 0:
        if subtext[0].isspace():
            spaces, subtext = _space_word(subtext)
            result.append(spaces)
            continue
        substr, subtext = _word_space(subtext)
        cons, vowel = _cons_vowel(substr)
        if len(cons) == 0 or cons.startswith('xr') or cons.startswith('yt'):
            result.extend((substr, 'ay'))
            continue
        ypos: Final[int] = cons.find('y')
        if ypos > 0:
            result.extend((substr[ypos:], substr[:ypos], 'ay'))
            continue
        if cons.endswith('q') and vowel.startswith('u'):
            result.extend((vowel[1:], substr[:len(cons) + 1], 'ay'))
            continue
        result.extend((vowel, cons, 'ay'))
    return ''.join(result)