#ifndef HIGH_SCORES_H
#define HIGH_SCORES_H

#include <stddef.h>
#include <stdint.h>

/// Return the latest score.
int32_t latest(const int32_t *const scores, const size_t scores_len);

/// Return the highest score.
int32_t personal_best(const int32_t *const scores, const size_t scores_len);

/// Write the highest scores to `output` (in non-ascending order).
/// Return the number of scores written.
size_t personal_top_three(const int32_t *const scores, const size_t scores_len, int32_t *const output);

#endif
