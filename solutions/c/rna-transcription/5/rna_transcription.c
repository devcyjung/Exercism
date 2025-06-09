#include "rna_transcription.h"

#include <stdlib.h>
#include <string.h>

static signed char translate(const char dna) {
    switch (dna) {
        case 'G': return 'C';
        case 'C': return 'G';
        case 'T': return 'A';
        case 'A': return 'U';
    }
    return ERROR_INVALID_CHAR;
}

char *to_rna(const char *const dna) {
    size_t len = strlen(dna);
    char *rna = malloc(sizeof(char) * (len + 1));
    if (!rna) {
        return NULL;
    }
    rna[len] = '\0';
    for (size_t i = 0; i < len; ++i) {
        rna[i] = translate(dna[i]);
    }
    return rna;
}