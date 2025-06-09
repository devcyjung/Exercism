#include "clock.h"
#include <stdio.h>
#include <string.h>

clock_t clock_create(int hour, int minute) {
    clock_t c = { "00:00" };
    return clock_add(c, hour * 60 + minute);
}

clock_t clock_add(clock_t clock, int minute_add) {
    minute_add += ((clock.text[0] - '0') * 10 + (clock.text[1] - '0')) * 60
        + (clock.text[3] - '0') * 10 + clock.text[4] - '0';
    minute_add %= 24 * 60;
    minute_add += 24 * 60;
    minute_add %= 24 * 60;
    sprintf(clock.text, "%02d:%02d", minute_add / 60, minute_add % 60);
    return clock;
}

clock_t clock_subtract(clock_t clock, int minute_subtract) {
    return clock_add(clock, -minute_subtract);
}

bool clock_is_equal(clock_t a, clock_t b) {
    return strcmp(a.text, b.text) == 0;
}