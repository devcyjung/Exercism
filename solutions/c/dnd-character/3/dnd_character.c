#include "dnd_character.h"

#include <stdlib.h>
#include <time.h>
#include <math.h>

static int dice_roll() {
    srand((unsigned int)time(NULL));
    return rand() % 6 + 1;
}

int ability() {
    int mini = dice_roll();
    int roll = 0;
    int sum = 0;
    for (unsigned int i = 0; i < 3; ++i) {
        roll = dice_roll();
        if (roll < mini) {
            sum += mini;
            mini = roll;
        } else {
            sum += roll;
        }
    }
    return sum;
}

int modifier(const int score) {
    return floor((score - 10) / 2.0);
}

dnd_character_t make_dnd_character() {
    dnd_character_t chr = {
        .strength = ability(),
        .dexterity = ability(),
        .constitution = ability(),
        .intelligence = ability(),
        .wisdom = ability(),
        .charisma = ability(),
    };
    chr.hitpoints = modifier(chr.constitution) + 10;
    return chr;
}