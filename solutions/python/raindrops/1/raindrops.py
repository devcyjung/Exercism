sound_map = {
    3: "Pling",
    5: "Plang",
    7: "Plong",
}

def convert(number):
    result = ""
    for divisor, sound in sound_map.items():
        if number % divisor == 0:
            result += sound
    if len(result) == 0:
        result += str(number)
    return result