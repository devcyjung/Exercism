#include "leap.h"

bool leap_year(int year) {
    return (!(year & 3) && year % 100) || !(year % 400);
}