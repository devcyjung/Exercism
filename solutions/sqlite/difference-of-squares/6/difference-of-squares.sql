WITH numbers(sum, square_sum) AS (
    SELECT sum(value), sum(value * value)
    FROM generate_series(0, number)
)
UPDATE "difference-of-squares"
SET result = (
    SELECT
        CASE
            WHEN property = "squareOfSum" THEN
                sum * sum
            WHEN property = "sumOfSquares" THEN
                square_sum
            WHEN property = "differenceOfSquares" THEN
                sum * sum - square_sum
        END
    FROM numbers
);