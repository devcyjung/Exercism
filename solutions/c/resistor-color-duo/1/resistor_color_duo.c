#include "resistor_color_duo.h"

uint16_t color_code(const resistor_band_t codes[2]) {
    return codes[0] * 10 + codes[1];
}