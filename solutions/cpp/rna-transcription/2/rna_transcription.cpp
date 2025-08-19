#include "rna_transcription.h"

#include <algorithm>

namespace rna_transcription {

char to_rna(const char dna) noexcept
{
    switch (dna)
    {
        case 'G': return 'C';
        case 'C': return 'G';
        case 'T': return 'A';           
        case 'A': return 'U';
        default: return dna;
    }
}

std::string to_rna(const std::string_view dna_strand) noexcept
{
    std::string rna_strand {};
    rna_strand.reserve(dna_strand.size());
    std::transform(std::begin(dna_strand), std::end(dna_strand), std::back_inserter(rna_strand), [](char c) { return to_rna(c); });
    return rna_strand;
}

}  // namespace rna_transcription
