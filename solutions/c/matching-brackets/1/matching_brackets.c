#include "matching_brackets.h"

#include <string.h>

static int translate(const char ch) {
    switch (ch) {
        case '(':
            return 1;
        case ')':
            return -1;
        case '{':
            return 2;
        case '}':
            return -2;
        case '[':
            return 3;
        case ']':
            return -3;
        default:
            return 0;    
    }
}

bool is_paired(const char *input) {
    int stack[strlen(input)];
    size_t stack_idx = 0;
    for (char ch = *input; ch; ch = *(++input)) {
        int i = translate(ch);
        if (i > 0) {
            stack[stack_idx++] = i;
        } else if (i < 0) {
            if (i + stack[--stack_idx] != 0) {
                return false;
            }
        }
    }
    return stack_idx == 0;
}