import re

def add_prefix_un(word):
    return "un" + word


def make_word_groups(vocab_words):
    acc = []
    prefix = ""
    for idx, word in enumerate(vocab_words):
        if idx == 0:
            prefix = word
            acc.append(prefix)
            continue
        acc.append(prefix + word)
    return " :: ".join(acc)

def remove_suffix_ness(word):
    if not word.endswith("ness"):
        return word
    trimmed = word.removesuffix("ness")
    if trimmed[-1] == "i" and trimmed[-2].lower() not in "aeiou":
        return trimmed[:-1] + "y"
    return trimmed
    
def adjective_to_verb(sentence, index):
    return re.sub(r"[^A-Za-z]", "", sentence.split()[index]) + "en"