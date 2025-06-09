#include "two_fer.h"

#include <stdio.h>
#include <string.h>

void two_fer(char *const buffer, const char *const name) {
    sprintf(buffer, "One for %s, one for me.", !name || strlen(name) == 0 ? "you" : name);
}