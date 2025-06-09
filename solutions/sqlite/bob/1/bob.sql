WITH
    trimmed AS (
        SELECT trim(input, char(9, 10, 13, 32)) AS trimmed
    ),
    is_question AS (
        SELECT substr(trimmed, -1) = "?" AS is_question
        FROM trimmed
    ),
    is_yelling AS (
        SELECT
            glob("*[a-zA-Z]*", trimmed)
            AND upper(trimmed) = trimmed AS is_yelling
        FROM trimmed
    )
UPDATE bob
SET reply = (
    SELECT CASE
        WHEN trimmed = "" THEN
            "Fine. Be that way!"
        WHEN is_question AND is_yelling THEN
            "Calm down, I know what I'm doing!"
        WHEN is_yelling THEN
            "Whoa, chill out!"
        WHEN is_question THEN
            "Sure."
        ELSE
            "Whatever."
    END
    FROM trimmed, is_question, is_yelling
);