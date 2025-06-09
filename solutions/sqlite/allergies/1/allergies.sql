WITH shr_offset(food, offsets) AS (
    VALUES
        ("eggs", 0),
        ("peanuts", 1),
        ("shellfish", 2),
        ("strawberries", 3),
        ("tomatoes", 4),
        ("chocolate", 5),
        ("pollen", 6),
        ("cats", 7)
)
UPDATE allergies
SET result =
    CASE
        WHEN task = "allergicTo" THEN (
            SELECT
                CASE
                    WHEN (score >> offsets) & 1 = 1 THEN
                        "true"
                    ELSE
                        "false"
                END
            FROM shr_offset
            WHERE food = item
        )
        WHEN task = "list" THEN (
            SELECT ifnull(group_concat(food, ", "), "")
            FROM shr_offset
            WHERE (score >> offsets) & 1 = 1
            ORDER BY offsets
        )
    END;


select * from allergies;