"""Functions for organizing and calculating student exam scores."""

import itertools as it


"""
Round all provided student scores.

:param scores: list - float or int of student exam scores.
:return: list - student scores *rounded* to nearest integer value.
"""

round_scores = lambda scores : \
    [round(score) for score in scores]


"""
Count the number of failing students out of the group provided.

:param scores: list - containing int student scores.
:return: int - count of student scores at or below 40.
"""

count_failed_students = lambda scores : \
    len([0 for score in scores if score <= 40])


"""
Determine how many of the provided student scores were 'the best' based on the provided threshold.

:param scores: list - of integer scores.
:param threshold: int - threshold to cross to be the "best" score.
:return: list - of integer scores that are at or above the "best" threshold.
"""

above_threshold = lambda scores, threshold : \
    [score for score in scores if score >= threshold]


"""
Create a list of grade thresholds based on the provided highest grade.

:param highest: int - value of highest exam score.
:param [lowest]: int - value of lowest exam score.
:return: list - of lower threshold scores for each D-A letter grade interval.
    For example, where the highest score is 100, and failing is <= 40,
    The result would be [41, 56, 71, 86]:

    41 <= "D" <= 55
    56 <= "C" <= 70
    71 <= "B" <= 85
    86 <= "A" <= 100
"""

letter_grades = lambda highest, lowest = 40 : \
    _get_intervals(lowest + 1, highest - lowest >> 2)

_get_intervals = lambda first, interval : \
    [first + interval * idx for idx in range(4)]


"""
Organize the student's rank, name, and grade information in descending order.

:param scores: list - of scores in descending order.
:param names: list - of string names by exam score in descending order.
:return: list - of strings in format ["<rank>. <student name>: <score>"].
"""

student_ranking = lambda scores, names : \
    [f'{rank}. {name}: {score}' for (rank, name, score) in zip(it.count(1), names, scores)]


"""
Create a list that contains the name and grade of the first student to make a perfect score on the exam.

:param students: list - of [<student name>, <score>] lists.
:return: list - first `[<student name>, 100]` or `[]` if no student score of 100 is found.
"""

perfect_score = lambda students : \
    next((student for student in students if len(student) == 2 if student[1] == 100), [])
