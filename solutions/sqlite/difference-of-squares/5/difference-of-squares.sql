WITH RECURSIVE numbers(n, sum, square_sum) AS (
    SELECT 0, 0, 0
    UNION
    SELECT n + 1, sum + n + 1, square_sum + power(n + 1, 2)
    FROM numbers
    WHERE n <= diff.number
)
UPDATE "difference-of-squares" AS diff
SET result = 
    CASE
        WHEN property = "squareOfSum" THEN
            (SELECT power(sum, 2) FROM numbers WHERE numbers.n = diff.number)
        WHEN property = "sumOfSquares" THEN
            (SELECT square_sum FROM numbers WHERE numbers.n = diff.number)
        WHEN property = "differenceOfSquares" THEN
            (SELECT power(sum, 2) - square_sum FROM numbers WHERE numbers.n = diff.number)
    END;