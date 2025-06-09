WITH day_index(day_string, day_index) AS (
        VALUES
            ("Sunday", 0),
            ("Monday", 1),
            ("Tuesday", 2),
            ("Wednesday", 3),
            ("Thursday", 4),
            ("Friday", 5),
            ("Saturday", 6)
    )
UPDATE meetup
SET result = (
    SELECT
        CASE week
            WHEN "teenth"
                THEN date(
                    printf("%04d-%02d-13", year, month),
                    "weekday " || day_index
                )
            WHEN "first"
                THEN date(
                    printf("%04d-%02d-01", year, month),
                    "weekday " || day_index
                )
            WHEN "second"
                THEN date(
                    printf("%04d-%02d-08", year, month),
                    "weekday " || day_index
                )
            WHEN "third"
                THEN date(
                    printf("%04d-%02d-15", year, month),
                    "weekday " || day_index
                )
            WHEN "fourth"
                THEN date(
                    printf("%04d-%02d-22", year, month),
                    "weekday " || day_index
                )
            WHEN "last"
                THEN date(
                    printf("%04d-%02d-01", year, month),
                    "+1 month", "start of month", "-7 days",
                    "weekday " || day_index
                )
        END
    FROM day_index
    WHERE day_string = dayofweek
);

SELECT * from meetup