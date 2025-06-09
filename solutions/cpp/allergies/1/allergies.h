#if !defined(ALLERGIES_H)
#define ALLERGIES_H

#include <array>
#include <string>
#include <unordered_set>

namespace allergies {

const std::array<std::string, 8> allergens = {"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"};
    
class Allergy {
    public:
        Allergy(int n);
        const std::unordered_set<std::string>& get_allergies();
        bool is_allergic_to(const std::string&);
    private:
        int n{};
        std::unordered_set<std::string> allergy_set{};
};

Allergy allergy_test(int n);

}  // namespace allergies

#endif  // ALLERGIES_H