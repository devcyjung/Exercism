WITH dice(dice) AS (
    SELECT value
    FROM json_each(json("[" || dice_results || "]"))
), util AS (
    SELECT
        count(DISTINCT dice) AS uniq,
        min(dice) AS minimum,
        max(dice) AS maximum,
        sum(dice) AS total
    FROM dice
), dice_freq AS (
    SELECT
        dice AS dice_value,
        count(dice) AS freq,
        rank() OVER (
            ORDER BY count(dice) desc
        ) AS freq_rank
    FROM dice
    GROUP BY dice
), freq_util AS (
    SELECT (
        SELECT dice_value
        FROM dice_freq
        WHERE freq_rank = 1
    ) AS most_common_dice, (
        SELECT dice_value
        FROM dice_freq
        WHERE freq_rank = (
            SELECT max(freq_rank)
            FROM dice_freq
        )
    ) AS least_common_dice
), freq_count_util AS (
    SELECT (
        SELECT freq
        FROM dice_freq
        WHERE dice_value = most_common_dice
    ) AS most_common_dice_count, (
        SELECT freq
        FROM dice_freq
        WHERE dice_value = least_common_dice
    ) AS least_common_dice_count
    FROM freq_util
)
UPDATE yacht
SET result = (
    SELECT
        CASE category
            WHEN "ones" THEN
                1 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 1
                ), 0)
            WHEN "twos" THEN
                2 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 2
                ), 0)
            WHEN "threes" THEN
                3 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 3
                ), 0)
            WHEN "fours" THEN
                4 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 4
                ), 0)
            WHEN "fives" THEN
                5 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 5
                ), 0)
            WHEN "sixes" THEN
                6 * coalesce((
                    SELECT freq
                    FROM dice_freq
                    WHERE dice_value = 6
                ), 0)
            WHEN "full house" THEN
                iif(most_common_dice_count = 3, total, 0)
            WHEN "four of a kind" THEN
                iif(most_common_dice_count >= 4, 4 * most_common_dice, 0)
            WHEN "little straight" THEN
                iif(uniq = 5 AND minimum = 1 AND maximum = 5, 30, 0)
            WHEN "big straight" THEN
                iif(uniq = 5 AND minimum = 2 AND maximum = 6, 30, 0)
            WHEN "choice" THEN
                total
            WHEN "yacht" THEN
                iif(uniq = 1, 50, 0)
        END
    FROM dice, util, freq_util, freq_count_util
);