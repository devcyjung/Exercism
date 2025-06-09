#include "high_scores.h"

int32_t latest(const int32_t *scores, size_t scores_len) {
    return scores[scores_len - 1];
}

int32_t personal_best(const int32_t *scores, size_t scores_len) {
    int32_t max_0 = 1 << 31;
    for (size_t i = 0; i < scores_len; ++i) {
        if (scores[i] > max_0) {
            max_0 = scores[i];
        }
    }
    return max_0;
}

/// Write the highest scores to `output` (in non-ascending order).
/// Return the number of scores written.
size_t personal_top_three(const int32_t *const scores, const size_t scores_len, int32_t *const output) {
    for (size_t j = 0; j < 3; ++j) {
        output[j] = 1 << 31;
    }
    int32_t temp;
    for (size_t i = 0; i < scores_len; ++i) {
        for (size_t j = 0; j < 3; ++j) {
            if (output[j] < scores[i]) {
                temp = output[j];
                output[j] = scores[i];
                for (size_t k = j + 1; k < 3; ++k) {
                    temp ^= output[k];
                    output[k] ^= temp;
                    temp ^= output[k];
                }
                break;
            }
        }
    }
    return (scores_len < 3) ? scores_len : 3;
}