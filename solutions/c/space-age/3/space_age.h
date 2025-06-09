#ifndef SPACE_AGE_H
#define SPACE_AGE_H

#include <stdint.h>

#define INVALID_INPUT -1

typedef enum planet {
   MERCURY,
   VENUS,
   EARTH,
   MARS,
   JUPITER,
   SATURN,
   URANUS,
   NEPTUNE,
} planet_t;

float age(const planet_t planet, const int64_t seconds);

#endif
