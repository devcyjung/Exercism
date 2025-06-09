#include "protein_translation.h"

#include <stdexcept>
#include <unordered_map>
namespace protein_translation {

    const std::unordered_map<std::string, std::string> codon_map = {
        {"AUG", "Methionine"},
        {"UUU", "Phenylalanine"}, {"UUC", "Phenylalanine"},
        {"UUA", "Leucine"},       {"UUG", "Leucine"},
        {"UCU", "Serine"},        {"UCC", "Serine"},
        {"UCA", "Serine"},        {"UCG", "Serine"},
        {"UAU", "Tyrosine"},      {"UAC", "Tyrosine"},
        {"UGU", "Cysteine"},      {"UGC", "Cysteine"},
        {"UGG", "Tryptophan"},
        {"UAA", "STOP"}, {"UAG", "STOP"}, {"UGA", "STOP"}
    };

    std::vector<std::string> proteins(const std::string& rna) {
        std::vector<std::string> result(0);
        if (rna.size() % 3 != 0) {
            throw std::invalid_argument("rna length must be multiple of 3");
        }
        for (size_t i = 0; i < rna.size(); i += 3) {
            auto amino = codon_map.find(rna.substr(i, 3)) ;
            if (amino == codon_map.end()) {
                throw std::invalid_argument("rna strand contains invalid codon");
            }
            if (amino->second == "STOP") {
                break;
            }
            result.push_back(amino->second);
        }
        return result;
    }

}  // namespace protein_translation
