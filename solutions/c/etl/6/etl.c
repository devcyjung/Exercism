#include "etl.h"

#include <ctype.h>
#include <stdlib.h>
#include <string.h>

static int compare_new_map(const void* lhs, const void* rhs)
{
    return ((const new_map*) lhs)->key - ((const new_map*) rhs)->key;
}

ptrdiff_t convert(const legacy_map* input, const size_t input_len, new_map** output)
{
    if (!input || !output)
        return -1;
    size_t total_keys = 0;
    for (size_t i = 0; i < input_len; ++i)
        total_keys += strlen(input[i].keys);
    const size_t output_length = total_keys;
    *output = malloc(sizeof(new_map) * output_length);
    new_map* output_buffer = *output;
    if (!output_buffer)
        return -1;
    size_t output_index = 0;
    for (size_t i = 0; i < input_len; ++i)
    {
        const int value = input[i].value;
        const char* key_ptr = input[i].keys;
        while (*key_ptr)
        {
            output_buffer[output_index++] = (new_map)
            {
                .key = tolower(*key_ptr++),
                .value = value
            };
        }
    }
    qsort(output_buffer, output_length, sizeof(new_map), compare_new_map);
    return output_length;
}
