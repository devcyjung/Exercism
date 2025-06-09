#include "two_fer.h"

#include <string.h>

void two_fer(char *const buffer, const char *const name) {
    static const char *const PREFIX = "One for ";
    static const char *const DEFAULT = "you";
    static const char *const SUFFIX = ", one for me.";
    strcpy(buffer, PREFIX);
    strcat(buffer, !name || strlen(name) == 0 ? DEFAULT : name);
    strcat(buffer, SUFFIX);
}