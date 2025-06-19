def find_anagrams(word, candidates):
    word_sorted = sorted(ch for ch in word.lower())
    return (
        candidate
        for candidate in candidates
        if len(candidate) == len(word)
        if candidate.lower() != word.lower()
        if sorted(ch for ch in candidate.lower()) == word_sorted
    )