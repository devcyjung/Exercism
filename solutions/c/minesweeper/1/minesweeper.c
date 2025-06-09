#include "minesweeper.h"
#include <stdlib.h>
#include <string.h>

static char mines(
    const size_t i, const size_t j, const size_t rows, const size_t cols,
    const char **minefield
) {
    if (minefield[i][j] == '*') {
        return '*';
    }
    char count = '0';
    for (int ni = i - 1; ni <= (int) i + 1; ++ni) {
        for (int nj = j - 1; nj <= (int) j + 1; ++nj) {
            if (ni < 0 || ni >= (int) rows || nj < 0 || nj >= (int) cols || (ni == (int) i && nj == (int) j)) {
                continue;
            }
            if ('*' == minefield[ni][nj]) {
                ++count;
            }
        }
    }
    if (count == '0') {
        return ' ';
    }
    return count;
}

char **annotate(const char **minefield, const size_t rows) {
    if (!minefield || !minefield[0]) {
        return NULL;
    }
    size_t rowlen = strlen(minefield[0]);
    char* buffer = malloc(sizeof(char) * (rowlen + 1) * rows);
    char** matrix = malloc(sizeof(char*) * rows);
    if (!buffer || !matrix) {
        free(buffer);
        free(matrix);
        return NULL;
    }
    for (size_t i = 0; i < rows; ++i) {
        matrix[i] = &buffer[i * (rowlen+1)];
        matrix[i][rowlen] = '\0';
        for (size_t j = 0; j < rowlen; ++j) {
            matrix[i][j] = mines(i, j, rows, rowlen, minefield);
        }
    }
    return matrix;
}

void free_annotation(char **annotation) {
    if (!annotation) {
        return;
    }
    free(annotation[0]);
    free(annotation);
}