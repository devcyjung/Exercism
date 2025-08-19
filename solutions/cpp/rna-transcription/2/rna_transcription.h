#pragma once

#include <string>
#include <string_view>

namespace rna_transcription {

char to_rna(char) noexcept;
std::string to_rna(std::string_view) noexcept;

}  // namespace rna_transcription
