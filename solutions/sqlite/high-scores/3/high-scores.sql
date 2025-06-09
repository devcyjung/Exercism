UPDATE results
SET result = (
    SELECT group_concat(score, ",")
    FROM scores
    WHERE scores.game_id = results.game_id
)
WHERE property = "scores";

UPDATE results
SET result = (
    SELECT score
    FROM scores
    WHERE scores.game_id = results.game_id
    ORDER BY rowid DESC
    LIMIT 1
)
WHERE property = "latest";

UPDATE results
SET result = (
    SELECT score
    FROM scores
    WHERE scores.game_id = results.game_id
    ORDER BY score DESC
    LIMIT 1
)
WHERE property = "personalBest";

UPDATE results
SET result = (
    SELECT group_concat(score, ",")
    FROM (
        SELECT score
        FROM scores
        WHERE scores.game_id = results.game_id
        ORDER BY score DESC
        LIMIT 3
    )
)
WHERE property = "personalTopThree";