"""Functions for organizing and calculating student exam scores."""

def round_scores(scores): return [round(score) for score in scores]

def count_failed_students(scores): return len([0 for score in scores if score <= 40])

def above_threshold(scores, threshold): return [score for score in scores if score >= threshold]

def letter_grades(highest, lowest = 40): return _get_intervals(lowest + 1, highest - lowest >> 2)

def _get_intervals(first, interval): return [first + interval * idx for idx in range(4)]

def student_ranking(scores, names): return [f'{idx + 1}. {name}: {score}' for (idx, (name, score)) in enumerate(zip(names, scores))]

def perfect_score(students): return next((student for student in students if len(student) == 2 if student[1] == 100), [])