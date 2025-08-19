#include "rna_transcription.h"

#include <algorithm>

namespace rna_transcription {

char to_rna(const char c)
{
    switch (c)
    {
        case 'G': return 'C';
        case 'C': return 'G';
        case 'T': return 'A';           
        case 'A': return 'U';
        default: return c;
    }
}

std::string to_rna(const std::string_view strand)
{
    std::string output {};
    output.reserve(strand.size());
    std::transform(std::begin(strand), std::end(strand), std::back_inserter(output), [](char c) { return to_rna(c); });
    return output;
}

}  // namespace rna_transcription
