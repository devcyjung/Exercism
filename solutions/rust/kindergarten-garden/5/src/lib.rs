pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    STUDENTS.iter().position(|s| *s == student).map_or_else(
        Vec::new,
        |student_index| diagram
            .split('\n')
            .flat_map(
                |row| row
                      .chars()
                      .skip(student_index * 2)
                      .take(2)
                      .flat_map(
                          |ch| MAPPINGS
                                .iter()
                                .find_map(|(c, s)| (*c == ch).then_some(s))
                      )
            )
            .cloned()
            .collect()
    )
}

const MAPPINGS: [(char, &'_ str); 4] = [
    ('G', "grass"),
    ('R', "radishes"),
    ('C', "clover"),
    ('V', "violets"),
];

const STUDENTS: [&'_ str; 12] = [
    "Alice", "Bob", "Charlie", "David", "Eve", "Fred",
    "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry",
];