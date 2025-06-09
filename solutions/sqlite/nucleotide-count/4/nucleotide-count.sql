DROP TABLE IF EXISTS "nucleotide-count";
CREATE TABLE "nucleotide-count" (
    "strand" TEXT,
    "result" TEXT
    CHECK(NOT glob("*[^ACGT]*", strand))
);

.mode csv
.import ./data.csv "nucleotide-count"

WITH chars(ch) AS (
    SELECT substr(strand, value, 1)
    FROM generate_series(1, length(strand))
), chars_enum(ch) AS (
    VALUES
        ('A'), ('C'), ('G'), ('T')
), char_count(ch, cnt) AS (
    SELECT e.ch, count(c.ch)
    FROM chars_enum AS e
    LEFT JOIN chars AS c
    USING(ch)
    GROUP BY e.ch
)
UPDATE "nucleotide-count"
SET result = (
    SELECT
        json_group_object(ch, cnt)
    FROM char_count
);