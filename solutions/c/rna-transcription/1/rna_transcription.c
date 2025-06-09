#include "rna_transcription.h"

#include <stdlib.h>
#include <string.h>

static signed char translate(char dna) {
    switch (dna) {
        case 'G': return 'C';
        case 'C': return 'G';
        case 'T': return 'A';
        case 'A': return 'U';
    }
    return ERROR_INVALID_CHAR;
}

char *to_rna(const char *dna) {
    size_t len = strlen(dna);
    char *rna = malloc(sizeof(char) * len);
    if (!rna) {
        return NULL;
    }
    for (size_t i = 0; i < len; ++i) {
        rna[i] = translate(dna[i]);
    }
    return rna;
}