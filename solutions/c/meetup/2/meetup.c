#include "meetup.h"

int meetup_day_of_month(
    unsigned int year, unsigned int month, const char *week, const char *day_of_week
) {
    int start_date;
    int is_last = 0;
    if (string_equal(week, "first")) {
        start_date = 1;
    } else if (string_equal(week, "second")) {
        start_date = 8;
    } else if (string_equal(week, "third")) {
        start_date = 15;
    } else if (string_equal(week, "fourth")) {
        start_date = 22;
    } else if (string_equal(week, "teenth")) {
        start_date = 13;
    } else if (string_equal(week, "last")) {
        start_date = last_date_in_month(year, month);
        is_last = 1;
    } else {
        return -1;
    }
    int start_weekday = weekday(year, month, start_date);
    int target_weekday;
    if (string_equal(day_of_week, "Sunday")) {
        target_weekday = 0;
    }
    else if (string_equal(day_of_week, "Monday")) {
        target_weekday = 1;
    } else if (string_equal(day_of_week, "Tuesday")) {
        target_weekday = 2;
    } else if (string_equal(day_of_week, "Wednesday")) {
        target_weekday = 3;
    } else if (string_equal(day_of_week, "Thursday")) {
        target_weekday = 4;
    } else if (string_equal(day_of_week, "Friday")) {
        target_weekday = 5;
    } else if (string_equal(day_of_week, "Saturday")) {
        target_weekday = 6;
    } else {
        return -1;
    }
    if (!is_last) {
        return start_date + (target_weekday - start_weekday + 7) % 7;
    } else {
        return start_date - (start_weekday - target_weekday + 7) % 7;
    }
}

int leap(unsigned int year) {
    return (year % 4 == 0 && year % 100 != 0) || year % 400 == 0;
}

int last_date_in_month(unsigned int year, unsigned int month) {
    switch (month) {
        case 1:
        case 3:
        case 5:
        case 7:
        case 8:
        case 10:
        case 12:
            return 31;
        case 4:
        case 6:
        case 9:
        case 11:
            return 30;
        default:
            if (leap(year)) {
                return 29;
            }
            return 28;
    }
}

int string_equal(const char *str1, const char *str2) {
    while (*str1 != 0 && *str2 != 0) {
        if (*str1 != *str2)
        {
            return 0;
        }
        ++str1;
        ++str2;
    }
    if (*str1 == *str2 && *str1 == 0) {
        return 1;
    }
    return 0;
}

int weekday(unsigned int year, unsigned int month, unsigned int date) {
    static int table[12] = {0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4};
    if (month < 3) {
        --year;
    }
    return (year + year / 4 - year / 100 + year / 400 + table[month - 1] + date) % 7;
}