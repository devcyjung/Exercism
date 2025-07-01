#include "luhn.h"
#include <string.h>

bool luhn(const char *num)
{
    if (!num)
        return false;
    size_t i = strlen(num);
    if (i == 0)
        return false;
    char ch;
    bool double_flag = false;
    size_t sum = 0;
    size_t cur;
    size_t digits = 0;
    do
    {
        ch = num[--i];
        if (ch == ' ')
        {
            continue;
        } else if (ch >= '0' && ch <= '9')
        {
            ++digits;
            cur = ch - '0';
            sum += !double_flag ? cur
                : cur > 4 ? (cur << 1) - 9
                    : (cur << 1);
            double_flag = !double_flag;
        } else
        {
            return false;
        }
    } while (i != 0);
    return digits > 1 && sum % 10 == 0;
}