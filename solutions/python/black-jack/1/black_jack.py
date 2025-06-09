"""Functions to help play and score a game of blackjack.
"""

def value_of_card(card):
    match card:
        case 'J' | 'Q' | 'K':
            return 10
        case 'A':
            return 1
        case _:
            return int(card)

def higher_card(card_one, card_two):
    match (value_of_card(card_one), value_of_card(card_two)):
        case (x, y) if x > y:
            return card_one
        case (x, y) if x < y:
            return card_two
        case _:
            return (card_one, card_two)

def value_of_ace(card_one, card_two):
    if card_one == 'A' or card_two == 'A':
        return 1
    match (value_of_card(card_one), value_of_card(card_two)):
        case (x, y) if x + y + 11 <= 21:
            return 11
        case _:
            return 1

def is_blackjack(card_one, card_two):
    return (
        (value_of_card(card_one) == 10 and card_two == 'A')
        or (value_of_card(card_two) == 10 and card_one == 'A')
    )


def can_split_pairs(card_one, card_two):
    return value_of_card(card_one) == value_of_card(card_two)


def can_double_down(card_one, card_two):
    return value_of_card(card_one) + value_of_card(card_two) in [9, 10, 11]
