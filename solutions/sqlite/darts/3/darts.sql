WITH radius AS (
    SELECT sqrt(pow(x, 2) + pow(y, 2)) AS radius
)
UPDATE darts
SET score = (
    SELECT
        CASE
            WHEN radius <= 1 THEN
                10
            WHEN radius <= 5 THEN
                5
            WHEN radius <= 10 THEN
                1
            ELSE
                0
        END
    FROM radius
);

SELECT * from darts;