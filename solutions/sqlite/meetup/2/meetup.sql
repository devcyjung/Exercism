-- Schema: CREATE TABLE "meetup" ( "year" INTEGER, "month" INTEGER, "week" TEXT, "dayofweek" TEXT, "result" TEXT);
WITH
    nth_offset(nth, offset_days) AS (
        VALUES
            ("second", 7),
            ("third", 14),
            ("fourth", 21),
            ("teenth", 12)
    ),
    day_index(day_string, day_index) AS (
        VALUES
            ("Sunday", 0),
            ("Monday", 1),
            ("Tuesday", 2),
            ("Wednesday", 3),
            ("Thursday", 4),
            ("Friday", 5),
            ("Saturday", 6)
    ),
    first_date_of_month AS(
        SELECT date(printf('%04d-%02d-01', year, month)) AS first_date_of_month
    ),
    start_date AS (
        SELECT CASE
            WHEN week = "first" THEN
                first_date_of_month
            WHEN week = "last" THEN
                date(
                    first_date_of_month,
                    "+1 month",
                    "start of month",
                    "-7 days"
                )
            ELSE
                date(
                    first_date_of_month,
                    "+" || (
                        SELECT offset_days FROM nth_offset
                        WHERE nth = week
                    ) || " days"
                )
        END AS start_date
        FROM first_date_of_month, nth_offset
    )
UPDATE meetup
SET result = (
    SELECT date(
        start_date,
        "weekday " || day_index
    )
    FROM start_date, day_index
    WHERE day_string = dayofweek
);

SELECT * FROM meetup;