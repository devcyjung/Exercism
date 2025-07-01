#include "nucleotide_count.h"
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct dna_count {
    size_t A, C, G, T;
} dna_count_t;

char *count(const char *dna_strand)
{
    char* result = calloc(100, sizeof(char));
    assert(result);
    if (!dna_strand)
        return result;
    dna_count_t dna_count = (dna_count_t){.A = 0, .C = 0, .G = 0, .T = 0};
    for (char ch = *dna_strand; ch; ch = *(++dna_strand))
    {
        switch (ch)
        {
            case 'A':
                ++dna_count.A;
                break;
            case 'C':
                ++dna_count.C;
                break;
            case 'G':
                ++dna_count.G;
                break;
            case 'T':
                ++dna_count.T;
                break;
            default:
                return result;
        }
    }
    sprintf(result, "A:%lu C:%lu G:%lu T:%lu",
            dna_count.A, dna_count.C, dna_count.G, dna_count.T);
    return result;
}