#pragma once

#include <string>
#include <vector>

namespace lasagna_master {

struct amount {
    int noodles;
    double sauce;
};

int preparationTime(const std::vector<std::string>& layers, int n = 2);
amount quantities(const std::vector<std::string>&);
void addSecretIngredient(std::vector<std::string>&, const std::vector<std::string>&);
std::vector<double> scaleRecipe(const std::vector<double>&, int);
void addSecretIngredient(std::vector<std::string>&, const std::string);
    
}  // namespace lasagna_master
