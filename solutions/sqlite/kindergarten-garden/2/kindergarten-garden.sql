WITH
    util AS (
        SELECT
            2 * (unicode(substr(student, 1, 1)) - unicode('A')) + 1 AS first_index,
            coalesce(
                trim(
                    substr(diagram, 1, instr(diagram, CHAR(10))),
                    char(9, 10, 13, 32)
                ),
                ""
            ) AS first_row,
            coalesce(
                trim(
                    substr(diagram, instr(diagram, CHAR(10))),
                    char(9, 10, 13, 32)
                ),
                ""
            ) AS second_row
    ),
    mapping(encoding, plant) AS (
        VALUES
            ("G", "grass"),
            ("C", "clover"),
            ("R", "radishes"),
            ("V", "violets")
    )
UPDATE "kindergarten-garden"
SET result = (
    SELECT (
        SELECT plant FROM mapping
        WHERE encoding = substr(first_row, first_index, 1)
    ) || "," || (
        SELECT plant FROM mapping
        WHERE encoding = substr(first_row, first_index + 1, 1)
    ) || "," || (
        SELECT plant FROM mapping
        WHERE encoding = substr(second_row, first_index, 1)
    ) || "," || (
        SELECT plant FROM mapping
        WHERE encoding = substr(second_row, first_index + 1, 1)
    )
    FROM util
);