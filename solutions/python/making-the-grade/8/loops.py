
# pylint: disable=C3001

round_scores = lambda scores : \
    [round(score) for score in scores]

count_failed_students = lambda scores : \
    len([0 for score in scores if score <= 40])

above_threshold = lambda scores, threshold : \
    [score for score in scores if score >= threshold]

letter_grades = lambda highest, lowest = 40 : \
    _get_intervals(lowest + 1, highest - lowest >> 2)

_get_intervals = lambda first, interval : \
    [first + interval * idx for idx in range(4)]

student_ranking = lambda scores, names : \
    [f'{idx + 1}. {name}: {score}' for (idx, (name, score)) in enumerate(zip(names, scores))]

perfect_score = lambda students : \
    next((student for student in students if len(student) == 2 if student[1] == 100), [])