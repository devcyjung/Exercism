#include "matching_brackets.h"
#include <string.h>

static signed char translate(const char ch)
{
    switch (ch)
    {
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

bool is_paired(const char *input)
{
    if (!input)
        return false;
    signed char translations[strlen(input)];
    size_t cursor = 0;
    signed char translation;
    for (char ch = *input; ch; ch = *(++input))
    {
        translation = translate(ch);
        if (translation == 0)
            continue;
        if (translation > 0)
        {
            translations[cursor++] = translation;
            continue;
        }
        if (cursor == 0 || translation + translations[--cursor] != 0)
            return false;
    }
    return cursor == 0;
}