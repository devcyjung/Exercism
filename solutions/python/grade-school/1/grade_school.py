class School:
    def __init__(self):
        self._name_grade = {}
        self._added = []

    def add_student(self, name, grade):
        added = name not in self._name_grade
        if added:
            self._name_grade[name] = grade
        self._added.append(added)

    def roster(self):
        return [
            name for name, _ in sorted(
                self._name_grade.items(),
                key = lambda item: (item[1], item[0])
            )
        ]

    def grade(self, grade_number):
        return [
            name for name, grade in sorted(
                self._name_grade.items(),
                key = lambda item: item[0]
            ) if grade == grade_number
        ]

    def added(self):
        return [added for added in self._added]
