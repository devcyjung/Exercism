from random import randrange

class Character:
    def __init__(self):
        self.strength = Character.ability()
        self.dexterity = Character.ability()
        self.constitution = Character.ability()
        self.intelligence = Character.ability()
        self.wisdom = Character.ability()
        self.charisma = Character.ability()
        self.hitpoints = modifier(self.constitution) + 10

    @staticmethod
    def ability():
        return sum(sorted(randrange(1, 7) for _ in range(4))[1:])

def modifier(value):
    return (value - 10) >> 1