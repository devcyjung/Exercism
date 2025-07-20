#include "all_your_base.h"

size_t rebase(int8_t *const digits, const int16_t input_base,
              const int16_t output_base, const size_t input_length) {
    if (!digits || !input_length) return 0;
    if (input_base <= 1 || output_base <= 1) {
        digits[0] = 0;
        return 0;
    }
    unsigned long long number = 0;
    for (size_t i = 0; i < input_length; ++i) {
        if (digits[i] < 0 || digits[i] >= input_base) {
            digits[0] = 0;
            return 0;
        }
        number *= input_base;
        number += digits[i];
    }
    size_t index = 0;
    do {
        digits[index++] = number % output_base;
        number /= output_base;
    } while (number > 0);
    int8_t temp;
    for (size_t left = 0, right = index - 1; left < right; ++left, --right) {
        temp = digits[left];
        digits[left] = digits[right];
        digits[right] = temp;
    }
    return index;
}