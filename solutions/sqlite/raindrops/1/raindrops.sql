WITH
    pling AS (
        SELECT CASE
            WHEN number % 3 = 0 THEN
                "Pling"
            ELSE
                ""
        END AS pling
    ),
    plang AS (
        SELECT CASE
            WHEN number % 5 = 0 THEN
                "Plang"
            ELSE
                ""
        END AS plang
    ),
    plong AS (
        SELECT CASE
            WHEN number % 7 = 0 THEN
                "Plong"
            ELSE    
                ""
        END AS plong
    ),
    concat AS (
        SELECT
            pling || plang || plong AS concat
        FROM pling, plang, plong
    )
UPDATE raindrops
SET sound = (
    SELECT CASE
        WHEN concat = "" THEN
            CAST(number AS TEXT)
        ELSE
            concat
    END
    FROM concat
);