WITH
    codes(color, code) AS (
        VALUES
            ("black", 0),
            ("brown", 1),
            ("red", 2),
            ("orange", 3),
            ("yellow", 4),
            ("green", 5),
            ("blue", 6),
            ("violet", 7),
            ("grey", 8),
            ("white", 9)
    ),
    code1 AS (
        SELECT code AS code1
        FROM codes
        WHERE color = color1
    ),
    code2 AS (
        SELECT code AS code2
        FROM codes
        WHERE color = color2
    )
UPDATE color_code
SET result = (
    SELECT code1 * 10 + code2
    FROM code1, code2
);