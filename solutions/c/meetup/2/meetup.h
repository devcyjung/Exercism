#ifndef MEETUP_H
#define MEETUP_H

int meetup_day_of_month(
    unsigned int year, unsigned int month, const char *week, const char *day_of_week);
int weekday(unsigned int year, unsigned int month, unsigned int date);
int string_equal(const char *str1, const char *str2);
int last_date_in_month(unsigned int year, unsigned int month);
int leap(unsigned int year);

#endif    // MEETUP_H