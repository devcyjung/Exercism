#include "darts.h"

uint8_t score(const coordinate_t position) {
    double radius_square = position.x * position.x + position.y * position.y;
    if (radius_square > 10 * 10) {
        return 0;
    }
    if (radius_square > 5 * 5) {
        return 1;
    }
    if (radius_square > 1 * 1) {
        return 5;
    }
    return 10;
}