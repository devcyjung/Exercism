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
    )
UPDATE color_code
SET result = (
    SELECT c1.code * 10 + c2.code
    FROM codes AS c1
    JOIN codes AS c2
    WHERE
        c1.color = color1
        AND c2.color = color2
);