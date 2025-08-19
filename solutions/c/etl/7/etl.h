#ifndef ETL_H
#define ETL_H

#include <stddef.h>

typedef struct
{
    int value;
    const char* keys;
} legacy_map;

typedef struct
{
    char key;
    int value;
} new_map;

ptrdiff_t convert(const legacy_map* input, size_t input_len, new_map** output);

#endif
