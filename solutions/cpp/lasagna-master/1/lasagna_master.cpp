#include "lasagna_master.h"

#include <algorithm>

namespace lasagna_master {

int preparationTime(const std::vector<std::string>& layers, int n) {
    return n * layers.size();
}

amount quantities(const std::vector<std::string>& layers) {
    amount a{};
    a.noodles = std::count(layers.begin(), layers.end(), "noodles") * 50;
    a.sauce = std::count(layers.begin(), layers.end(), "sauce") * 0.2;
    return a;
}
    
void addSecretIngredient(std::vector<std::string>& my_recipe, const std::vector<std::string>& friends_recipe) {
    my_recipe.back() = friends_recipe.back();
}
    
std::vector<double> scaleRecipe(const std::vector<double>& two_portions, int n) {
    std::vector<double> result(two_portions.size());
    std::transform(two_portions.begin(), two_portions.end(), result.begin(), [&n](const double& for_two){ return for_two / 2 * n; });
    return result;
}
    
void addSecretIngredient(std::vector<std::string>& my_recipe, const std::string secret) {
    my_recipe.back() = secret;    
}

}  // namespace lasagna_master
