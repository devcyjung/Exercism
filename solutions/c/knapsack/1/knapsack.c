#include "knapsack.h"

unsigned int maximum_value(
    unsigned int maximum_weight, item_t *items, size_t item_count
) {
    if (item_count == 0) {
        return 0;
    }
    item_t last_item = items[item_count - 1];
    if (last_item.weight > maximum_weight) {
        return maximum_value(maximum_weight, items, item_count - 1);
    }
    int without = maximum_value(maximum_weight, items, item_count - 1);
    int with = last_item.value + maximum_value(
        maximum_weight - last_item.weight, items, item_count - 1
    );
    if (without > with) {
        return without;
    }
    return with;
}