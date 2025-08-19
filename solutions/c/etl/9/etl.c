#include "etl.h"

#include <ctype.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static int compare_new_map(const void* const lhs, const void* const rhs)
{
    const unsigned char lhs_key = (unsigned char) ((const new_map*) lhs)->key;
    const unsigned char rhs_key = (unsigned char) ((const new_map*) rhs)->key;
    return (lhs_key > rhs_key) - (lhs_key < rhs_key);
}

etl_ssize_t convert(const legacy_map* const input, const size_t input_len, new_map** const output)
{
    if (!input || !output)
        return -1;
    size_t total_keys = 0;
    for (size_t i = 0; i < input_len; ++i)
        total_keys += input[i].keys ? strlen(input[i].keys) : 0;
    const size_t output_length = total_keys;
    if (output_length > SIZE_MAX / sizeof(new_map))
        return -1;
    new_map* const output_buffer = malloc(sizeof(new_map) * output_length);
    if (!output_buffer)
        return -1;
    *output = output_buffer;
    size_t output_index = 0;
    for (size_t i = 0; i < input_len; ++i)
    {
        const int value = input[i].value;
        const char* key_ptr = input[i].keys;
        if (!key_ptr)
            continue;
        while (*key_ptr)
        {
            if (output_index >= output_length)
            {
                free(output_buffer);
                return -1;
            }
            output_buffer[output_index++] = (new_map)
            {
                .key = tolower((unsigned char) *key_ptr++),
                .value = value
            };
        }
    }
    qsort(output_buffer, output_length, sizeof(new_map), compare_new_map);
    return output_length;
}
