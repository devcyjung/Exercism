WITH ws_removed(trimmed) AS (
    VALUES
        (replace(value, " ", ""))
), digits(back_ind, digit) AS (
    SELECT
        length(trimmed) - series.value,
        CAST(substr(trimmed, series.value, 1) AS INTEGER)
    FROM ws_removed
    LEFT JOIN generate_series(1, length(trimmed)) AS series
), luhn_number(luhn_num) AS (
    SELECT sum(
        iif(
            back_ind & 1 = 0,
            digit,
            iif(
                digit * 2 > 9,
                digit * 2 - 9,
                digit * 2
            )
        )
    )
    FROM digits
)
UPDATE luhn
SET result = (
    SELECT iif(
        NOT glob("*[^0-9]*", trimmed) AND length(trimmed) > 1,
        mod(luhn_num, 10) = 0,
        false
    )
    FROM ws_removed, luhn_number
);