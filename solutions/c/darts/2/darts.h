#ifndef DARTS_H
#define DARTS_H

#include <stdint.h>

typedef struct {
    double x; double y;
} coordinate_t;

uint8_t score(const coordinate_t position);

#endif
