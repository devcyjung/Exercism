//! Kindergarten Garden::Version 2::Goodbye vectors, hello iterators.

/// Determine the plants belonging to `student` given a `diagram`.
pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    diagram
        .lines()
        .flat_map(student_plants_getter(student))
        .collect()
}

/// List of students, in alphabetic order.
pub const STUDENTS: [&str; 12] = [
    "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph",
    "Kincaid", "Larry",
];

fn get_plant_name(repr: char) -> &'static str {
    match repr {
        'C' => "clover",
        'G' => "grass",
        'R' => "radishes",
        'V' => "violets",
        _ => panic!("Uh oh, unknown plant, could be poisonous!"),
    }
}

fn student_plants_getter<'a>(
    student: &str,
) -> impl FnMut(&'a str) -> Box<(dyn Iterator<Item=&'static str> + 'a)> {
    let pos = STUDENTS
        .binary_search(&student)
        .unwrap_or_else(|_| panic!("Student {} not found", student));
    move |row| Box::new(row[pos * 2..(pos + 1) * 2].chars().map(get_plant_name))
}