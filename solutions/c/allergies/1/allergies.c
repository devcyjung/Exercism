#include "allergies.h"

bool is_allergic_to(allergen_t allergen, unsigned int score) {
    if (allergen >= ALLERGEN_COUNT)
        return false;
    return ((score >> allergen) & 1) == 1;
}

allergen_list_t get_allergens(unsigned int score) {
    allergen_list_t result = (allergen_list_t){.count = 0};
    for (allergen_t allergen = 0; allergen < ALLERGEN_COUNT; ++allergen)
        if ((result.allergens[allergen] = is_allergic_to(allergen, score)))
            ++result.count;
    return result;
}