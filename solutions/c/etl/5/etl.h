#ifndef ETL_H
#define ETL_H

#include <unistd.h>

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

ssize_t convert(const legacy_map* input, size_t input_len, new_map** output);

#endif
