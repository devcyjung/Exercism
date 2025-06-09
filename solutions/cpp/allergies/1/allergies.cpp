#include "allergies.h"

#include <iterator>

namespace allergies {

    Allergy::Allergy(int n): n{n} {
        for (auto it = allergens.begin(); it != allergens.end(); ++it) {
            if (((n >> std::distance(allergens.begin(), it)) & 1) == 1) {
                allergy_set.emplace(*it);
            }
        }
    }
    
    const std::unordered_set<std::string>& Allergy::get_allergies() {
        return allergy_set; 
    }

    bool Allergy::is_allergic_to(const std::string& thing) {
        return allergy_set.find(thing) != allergy_set.end();
    }

    Allergy allergy_test(int n) {
        return Allergy(n);
    }

}  // namespace allergies
