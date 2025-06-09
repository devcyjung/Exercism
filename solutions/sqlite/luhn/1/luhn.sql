WITH RECURSIVE digits(pos, rpos, digit) AS (
    SELECT 1, length(space_removed), CAST(substr(space_removed, 1, 1) AS INTEGER)
    UNION
    SELECT pos + 1, rpos - 1, CAST(substr(space_removed, pos + 1, 1) AS INTEGER)
    FROM digits
    WHERE rpos > 1
), space_removed(space_removed) AS (
    SELECT replace(value, " ", "")
), compute(compute) AS (
    SELECT sum(
        CASE
            WHEN rpos % 2 = 1 THEN digit
            WHEN 2 * digit > 9 THEN 2 * digit - 9
            ELSE 2 * digit
        END
    )
    FROM digits
)
UPDATE luhn
SET result = (
    SELECT
        iif(
            NOT glob("*[^0-9 ]*", value) AND length(space_removed) > 1,
            (
                SELECT compute % 10 = 0
                FROM compute
            ),
            1 != 1
        )
    FROM space_removed
);