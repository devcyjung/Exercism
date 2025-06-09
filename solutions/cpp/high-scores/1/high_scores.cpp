#include "high_scores.h"

#include <algorithm>
#include <functional>
#include <stdexcept>

namespace arcade {

std::vector<int> HighScores::list_scores() {
    return scores;
}

int HighScores::latest_score() {
    if (scores.empty()) {
        throw std::runtime_error("empty scoreboard");
    }
    return *std::prev(scores.end(), 1);
}

int HighScores::personal_best() {
    if (scores.empty()) {
        throw std::runtime_error("empty scoreboard");
    }
    return *std::max_element(scores.begin(), scores.end());
}

std::vector<int> HighScores::top_three() {
    if (scores.empty()) {
        throw std::runtime_error("empty scoreboard");
    }
    auto temp = scores;
    std::sort(temp.begin(), temp.end(), std::greater<int>());
    return std::vector<int>(temp.begin(), std::min(temp.begin()+3, temp.end()));
}

}  // namespace arcade
