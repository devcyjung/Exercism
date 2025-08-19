#include "etl.h"

#include <ctype.h>
#include <stdlib.h>
#include <string.h>

static int compare_new_map(const void* lhs, const void* rhs)
{
    const char lhs_key = ((const new_map*) lhs)->key;
    const char rhs_key = ((const new_map*) rhs)->key;
    return (lhs_key > rhs_key) - (lhs_key < rhs_key);
}

ptrdiff_t convert(const legacy_map* input, const size_t input_len, new_map** output)
{
    if (!input || !output)
        return -1;
    size_t total_keys = 0;
    for (size_t i = 0; i < input_len; ++i)
        total_keys += input[i].keys ? strlen(input[i].keys) : 0;
    const size_t output_length = total_keys;
    new_map* output_buffer = malloc(sizeof(new_map) * output_length);
    if (!output_buffer)
        return -1;
    *output = output_buffer;
    size_t output_index = 0;
    for (size_t i = 0; i < input_len; ++i)
    {
        const int value = input[i].value;
        const char* key_ptr = input[i].keys;
        while (*key_ptr)
            output_buffer[output_index++] = (new_map)
            {
                .key = tolower((unsigned char) *key_ptr++),
                .value = value
            };
    }
    qsort(output_buffer, output_length, sizeof(new_map), compare_new_map);
    return output_length;
}
