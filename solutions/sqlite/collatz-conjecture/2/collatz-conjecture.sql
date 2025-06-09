WITH RECURSIVE rec(step, rem) AS (
    SELECT 0, CAST(number AS INTEGER)
    UNION
    SELECT step + 1, iif(rem & 1 = 0, rem >> 1, rem * 3 + 1)
    FROM rec
    WHERE rem > 1
)
UPDATE collatz
SET steps = (
    SELECT step
    FROM rec
    WHERE rem = 1
);