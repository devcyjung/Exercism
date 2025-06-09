WITH RECURSIVE chars(ch, pos) AS (
    SELECT substr(lower, 1, 1), 1
    FROM lower
    UNION
    SELECT substr(lower, pos + 1, 1), pos + 1
    FROM lower, chars
    WHERE pos < length(lower)
), lower(lower) AS (
    SELECT lower(phrase)
)
UPDATE isogram
SET is_isogram = (
    SELECT NOT EXISTS (
        SELECT ch
        FROM (
            SELECT ch
            FROM (
                SELECT ch, count(ch) AS cnt
                FROM chars
                GROUP BY ch
            )
            WHERE cnt > 1
        )
        WHERE ch NOT IN (CHAR(9), CHAR(10), CHAR(13), CHAR(32), CHAR(45))
    )
);