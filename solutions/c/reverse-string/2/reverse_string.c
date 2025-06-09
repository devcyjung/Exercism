#include "reverse_string.h"

#include <stdlib.h>
#include <string.h>

char *reverse(const char *value) {
    size_t len = strlen(value);
    char *dest = malloc(sizeof(char[len + 1]));
    dest[len] = 0;
    for (size_t i = 0; i < len; ++i) {
        dest[len - 1 - i] = value[i];
    }
    return dest;
}