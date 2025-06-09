#include "meetup.h"

#include <string.h>
#include <time.h>

static const char *WEEKS[6] = {"first", "second", "third", "fourth", "teenth", "last"};
static const int START_DATES[6] = {1, 8, 15, 22, 13, 0};
static const int MONTH_DELTAS[6] = {-1, -1, -1, -1, -1, 0};
static const char *WEEKDAYS[7] = {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"};

int meetup_day_of_month(
    unsigned int year, unsigned int month, const char *week, const char *day_of_week
) {
    int date_idx = -1;
    for (size_t i = 0; i < 6; ++i) {
        if (strcmp(week, WEEKS[i]) == 0) {
            date_idx = i;
            break;
        }
    }
    int target_wday = -1;
    for (size_t i = 0; i < 7; ++i) {
        if (strcmp(day_of_week, WEEKDAYS[i]) == 0) {
            target_wday = i;
            break;
        }
    }
    if (date_idx == -1 || target_wday == -1) {
        return -1;
    }
    int is_last = date_idx == 5;
    struct tm *start_tm = &((struct tm){
        .tm_year = year - 1900,
        .tm_mon = month + MONTH_DELTAS[date_idx],
        .tm_mday = START_DATES[date_idx],
    });
    if (mktime(start_tm) == -1) {
        return -1;
    }
    if (!is_last) {
        return start_tm->tm_mday + (target_wday - start_tm->tm_wday + 7) % 7;
    } else {
        return start_tm->tm_mday - (start_tm->tm_wday - target_wday + 7) % 7;
    }
}