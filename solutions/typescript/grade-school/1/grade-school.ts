declare global {
  interface ObjectConstructor {
    groupBy<T, K extends PropertyKey>(
      items: Iterable<T>,
      callback: (item: T) => K
    ): Record<K, T[]>
  }
}

export class GradeSchool {
  private school: Map<string, number> = new Map()
  
  roster(): Record<number, string[]> {
    return Object.fromEntries(
      Object.entries(
        Object
          .groupBy(this.school.entries(), ([_, grade]) => grade))
          .map(([grade, students]) => [grade, students.map(([name, _]) => name).sort()])
    )
  }

  add(student: string, grade: number): void {
    this.school.set(student, grade)
  }

  grade(level: number): string[] {
    return this.roster()[level] || []
  }
}