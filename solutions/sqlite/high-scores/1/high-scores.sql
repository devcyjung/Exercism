UPDATE results
SET result = (
    CASE property
        WHEN "scores" THEN (
            SELECT group_concat(score, ",")
            FROM scores
            WHERE scores.game_id = results.game_id
        )
        WHEN "latest" THEN (
            SELECT score
            FROM scores
            WHERE scores.game_id = results.game_id
            ORDER BY rowid DESC
            LIMIT 1
        )
        WHEN "personalBest" THEN (
            SELECT score
            FROM scores
            WHERE scores.game_id = results.game_id
            ORDER BY score DESC
            LIMIT 1
        )
        WHEN "personalTopThree" THEN (
            SELECT group_concat(score, ",")
            FROM (
                SELECT score
                FROM scores
                WHERE scores.game_id = results.game_id
                ORDER BY score DESC
                LIMIT 3
            )
        )
    END
);

select * from scores;
select * from results;