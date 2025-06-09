WITH grains(grid, grain) AS (
    SELECT
        value, pow(2, (value - 1))
    FROM generate_series(1, 64)
)
UPDATE grains
SET result = 
    CASE
        WHEN task = "single-square" THEN (
            SELECT grain
            FROM grains
            WHERE grid = square
        )
        WHEN task = "total" THEN (
            SELECT sum(grain)
            FROM grains
        )
    END;