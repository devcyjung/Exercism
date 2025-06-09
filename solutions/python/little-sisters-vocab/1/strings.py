import re

def add_prefix_un(word):
    return "un" + word


def make_word_groups(vocab_words):
    acc = []
    prefix = ""
    for i, word in enumerate(vocab_words):
        if i == 0:
            prefix = word
            acc.append(prefix)
            continue
        acc.append(prefix + word)
    return " :: ".join(acc)

def remove_suffix_ness(word):
    if not word.endswith("ness"):
        return word
    s = word.removesuffix("ness")
    if s[-1] == "i" and s[-2].lower() not in "aeiou":
        return s[:-1] + "y"
    return s
    
def adjective_to_verb(sentence, index):
    return re.sub(r"[^A-Za-z]", "", sentence.split()[index]) + "en"