WITH len(len) AS (
        SELECT length(cast(number as text))
    ),
    total(total) AS (
        SELECT sum(
            pow(
                mod(
                    floor(number / pow(10, value)),
                    10
                ),
                len
            )
        )
        FROM len
        JOIN generate_series(0, len - 1)
    )
UPDATE "armstrong-numbers"
SET result = (
    SELECT total = number
    FROM total
);