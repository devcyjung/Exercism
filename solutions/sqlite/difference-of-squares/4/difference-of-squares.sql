UPDATE "difference-of-squares"
SET result = number * number * (number + 1) * (number + 1) / 4
WHERE property = "squareOfSum";

UPDATE "difference-of-squares"
SET result = number * (number + 1) * (2 * number + 1) / 6
WHERE property = "sumOfSquares";

WITH sums AS (
    SELECT
        number,
        MAX(CASE WHEN property = "squareOfSum" THEN result END) AS square_of_sum,
        MAX(CASE WHEN property = "sumOfSquares" THEN result END) AS sum_of_squares
    FROM "difference-of-squares"
    GROUP BY number
)
UPDATE "difference-of-squares" AS diff
SET result = (
    SELECT square_of_sum - sum_of_squares
    FROM sums
    WHERE sums.number = diff.number
)
WHERE property = "differenceOfSquares";