#ifndef MINESWEEPER_H
#define MINESWEEPER_H
#include <stddef.h>

char **annotate(const char *const *const minefield, const size_t rows);
void free_annotation(char **annotation);

#endif
