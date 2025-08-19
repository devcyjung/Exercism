#ifndef ETL_H
#define ETL_H

#if __has_include("sys/types.h")
    #include <sys/types.h>
    typedef ssize_t etl_ssize_t;
#else   // __has_include("sys/types.h")
    #include <stddef.h>
    typedef ptrdiff_t etl_ssize_t;
#endif  // __has_include("sys/types.h")

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

etl_ssize_t convert(const legacy_map* input, size_t input_len, new_map** output);

#endif
