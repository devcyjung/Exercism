#include <array>
#include <algorithm>
#include <iterator>
#include <string>
#include <vector>

std::vector<int> round_down_scores(
    std::vector<double> student_scores
) {
    std::vector<int> result(student_scores.size());
    std::transform(
        student_scores.begin(),
        student_scores.end(),
        result.begin(),
        [](double d){ return static_cast<int>(d); }
    );
    return result;
}

int count_failed_students(
    std::vector<int> student_scores
) {
    return std::count_if(
        student_scores.begin(),
        student_scores.end(),
        [](int i){ return i <= 40; }
    );
}

std::array<int, 4> letter_grades(int highest_score) {
    int interval = (highest_score - 40) / 4;
    return {41, 41 + interval, 41 + 2 * interval, 41 + 3 * interval};
}

std::vector<std::string> student_ranking(
    std::vector<int> student_scores,
    std::vector<std::string> student_names
) {
    std::vector<std::string> result(student_scores.size());
    auto begin = student_scores.begin();
    for (auto it = begin; it != student_scores.end(); ++it) {
        int rank = 1 + std::distance(begin, it);
        result[rank - 1] = std::to_string(rank) + ". " + student_names[rank - 1] + ": " + std::to_string(*it);
    }
    return result;
}

std::string perfect_score(std::vector<int> student_scores,
                          std::vector<std::string> student_names) {
    auto it = std::find(student_scores.begin(), student_scores.end(), 100);
    if (it != student_scores.end()) {
        return student_names[std::distance(student_scores.begin(), it)];
    }
    return "";
}
