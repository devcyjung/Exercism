#include "resistor_color_trio.h"
#include <math.h>

resistor_value_t color_code(const resistor_band_t colors[static 3]) {
    unit_t unit = colors[2] / 3;
    uint16_t value = (10 * colors[0] + colors[1]) * pow(10, colors[2] % 3);
    while (value >= 1000) {
        value /= 1000;
        ++unit;
    }
    return (resistor_value_t) {
        .unit = unit,
        .value = value,
    };
}