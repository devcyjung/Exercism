WITH RECURSIVE digits(digit, multiplier, rem) AS (
    SELECT
        cast_int % 10, 1, cast_int / 10
    FROM cast_int
    UNION
    SELECT
        rem % 10, multiplier * 10, rem / 10
    FROM digits
    WHERE rem > 0
), cast_int(cast_int) AS (
    VALUES
        (CAST(number AS INTEGER))
), literals(multiplier, literal) AS (
    VALUES
        (1, "I"),
        (5, "V"),
        (10, "X"),
        (50, "L"),
        (100, "C"),
        (500, "D"),
        (1000, "M")
), numerals(multiplier, literal) AS (
    SELECT
        d.multiplier,
        CASE
            WHEN digit = 0 THEN
                ""
            WHEN digit < 4 THEN
                printf("%.*c", digit, single.literal)
            WHEN digit = 4 THEN
                single.literal || five.literal
            WHEN digit = 5 THEN
                five.literal
            WHEN digit < 9 THEN
                five.literal || printf("%.*c", digit - 5, single.literal)      
            WHEN digit = 9 THEN
                single.literal || ten.literal
        END
    FROM digits AS d
    LEFT JOIN literals AS single
    USING(multiplier)
    LEFT JOIN literals AS five
    ON d.multiplier * 5 = five.multiplier
    LEFT JOIN literals AS ten
    ON d.multiplier * 10 = ten.multiplier
    ORDER BY d.multiplier DESC
)
UPDATE "roman-numerals"
SET result = (
    SELECT group_concat(literal, "")
    FROM numerals
);