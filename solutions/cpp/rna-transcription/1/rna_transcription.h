#pragma once

#include <string>
#include <string_view>

namespace rna_transcription {

char to_rna(char);
std::string to_rna(std::string_view);

}  // namespace rna_transcription
