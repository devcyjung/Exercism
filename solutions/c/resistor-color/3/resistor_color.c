#include "resistor_color.h"

unsigned int color_code(resistor_band_t color) {
    return (unsigned int) color;
}

const resistor_band_t *colors() {
    static const resistor_band_t COLORS[10] = {
        0, 1, 2, 3, 4, 5, 6, 7, 8, 9
    };
    return COLORS;
}