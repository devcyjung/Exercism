#include "meetup.h"

#include <string.h>
#include <time.h>

int meetup_day_of_month(
    unsigned int year, unsigned int month, const char *week, const char *day_of_week
) {
    int start_date;
    int is_last = 0;
    if (!strcmp(week, "first")) {
        start_date = 1;
    } else if (!strcmp(week, "second")) {
        start_date = 8;
    } else if (!strcmp(week, "third")) {
        start_date = 15;
    } else if (!strcmp(week, "fourth")) {
        start_date = 22;
    } else if (!strcmp(week, "teenth")) {
        start_date = 13;
    } else if (!strcmp(week, "last")) {
        ++month;
        start_date = 0;
        is_last = 1;
    } else {
        return -1;
    }
    struct tm *start_tm = &((struct tm){
        .tm_year = year - 1900,
        .tm_mon = month - 1,
        .tm_mday = start_date,
    });
    if (mktime(start_tm) == -1) {
        return -1;
    }
    int target_weekday;
    if (!strcmp(day_of_week, "Sunday")) {
        target_weekday = 0;
    }
    else if (!strcmp(day_of_week, "Monday")) {
        target_weekday = 1;
    } else if (!strcmp(day_of_week, "Tuesday")) {
        target_weekday = 2;
    } else if (!strcmp(day_of_week, "Wednesday")) {
        target_weekday = 3;
    } else if (!strcmp(day_of_week, "Thursday")) {
        target_weekday = 4;
    } else if (!strcmp(day_of_week, "Friday")) {
        target_weekday = 5;
    } else if (!strcmp(day_of_week, "Saturday")) {
        target_weekday = 6;
    } else {
        return -1;
    }
    if (!is_last) {
        return start_tm->tm_mday + (target_weekday - start_tm->tm_wday + 7) % 7;
    } else {
        return start_tm->tm_mday - (start_tm->tm_wday - target_weekday + 7) % 7;
    }
}