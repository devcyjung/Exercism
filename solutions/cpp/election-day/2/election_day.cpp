#include <algorithm>
#include <string>
#include <vector>

namespace election {

struct ElectionResult {
    std::string name{};
    int votes{};
};

int vote_count(const ElectionResult& result) {
    return result.votes;
}

void increment_vote_count(ElectionResult& result, int votes) {
    result.votes += votes;
}

ElectionResult& determine_result(std::vector<ElectionResult>& final_count) {
    auto winner = std::max_element(final_count.begin(), final_count.end(), [](ElectionResult a, ElectionResult b) {
        return a.votes < b.votes;
    });
    winner->name = "President " + winner->name;
    return *winner;
}

}  // namespace election