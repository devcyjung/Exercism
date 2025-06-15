#include "gigasecond.h"

void gigasecond(const time_t input, char *const output, const size_t size) {
    time_t gigaseconds_later = input + 1e9;
    strftime(output, size, "%F %T", gmtime(&gigaseconds_later));
}