WITH RECURSIVE grain(n, grain) AS (
    SELECT 1, 1
    UNION
    SELECT n + 1, grain * 2
    FROM grain
    LIMIT 64
)
UPDATE grains

SET result = 
    CASE
        WHEN task = "single-square" THEN (
            SELECT grain
            FROM grain
            WHERE n = square
        )
        WHEN task = "total" THEN (
            SELECT SUM(grain)
            FROM grain
        )
    END;