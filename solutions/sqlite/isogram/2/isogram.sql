WITH ws_hyphen(ws_hyphen) AS (
    VALUES
        (CHAR(9, 10, 11, 12, 13, 32, 45))
), chars(ch) AS (
    SELECT substr(lower(phrase), value, 1)
    FROM generate_series(1, length(phrase))
), char_count(cnt) AS (
    SELECT count(ch)
    FROM chars
    LEFT JOIN ws_hyphen
    WHERE instr(ws_hyphen, ch) = 0
    GROUP BY ch
)
UPDATE isogram
SET is_isogram = (
    SELECT coalesce(max(cnt), 0)
    FROM char_count
) < 2;

SELECT * from isogram