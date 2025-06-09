WITH RECURSIVE rows(r_nth, line) AS (
    VALUES
        (2, json_array(1, 1))
    UNION
    SELECT
        r_nth + 1,
        json("[1," || (
            SELECT group_concat(
                CAST(line ->> (value - 1)  AS INTEGER)
                + CAST(line ->> value AS INTEGER),
                ","
            )
            FROM generate_series(1, json_array_length(line) - 1)
        ) || ",1]")
    FROM rows
    WHERE r_nth <= input
), pyramid_rows(pr_nth, pyramid_row) AS (
    VALUES
        (1, "1")
    UNION
    SELECT
        pr_nth + 1, (
            SELECT group_concat(value, " ")
            FROM json_each((
                SELECT line
                FROM rows
                WHERE pr_nth + 1 = r_nth 
            ))
        )
    FROM pyramid_rows
    WHERE pr_nth <= input
), pyramid(p_nth, pyramid) AS (
    SELECT i.value, (
        SELECT group_concat(pyramid_row, CHAR(10))
        FROM pyramid_rows
        WHERE pr_nth <= i.value
    )
    FROM generate_series(1, input) AS i
)
UPDATE "pascals-triangle"
SET result = (
    iif(input = 0, "", (
        SELECT pyramid
        FROM pyramid
        WHERE input = p_nth
    ))
);