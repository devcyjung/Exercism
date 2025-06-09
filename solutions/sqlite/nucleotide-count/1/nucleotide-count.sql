DROP TABLE IF EXISTS "nucleotide-count";
CREATE TABLE "nucleotide-count" (
    "strand" TEXT,
    "result" TEXT
    CHECK (
        NOT glob("*[^ACGT]*", strand)
    )
);

.mode csv
.import ./data.csv "nucleotide-count"

WITH RECURSIVE
    chars(pos, ch) AS (
        SELECT 1, substr(strand, 1, 1)
        UNION
        SELECT pos + 1, substr(strand, pos + 1, 1)
        FROM chars
        WHERE pos < length(strand)
    ),
    char_count(ch, cnt) AS (
        SELECT ch, COUNT(pos)
        FROM chars
        GROUP BY ch
    )
UPDATE "nucleotide-count"
SET result = (
    SELECT json_object(
        "A", coalesce((
            SELECT cnt FROM char_count
            WHERE ch = "A"
        ), 0),
        "C", coalesce((
            SELECT cnt FROM char_count
            WHERE ch = "C"
        ), 0),
        "G", coalesce((
            SELECT cnt FROM char_count
            WHERE ch = "G"
        ), 0),
        "T", coalesce((
            SELECT cnt FROM char_count
            WHERE ch = "T"
        ), 0)
    )
    FROM char_count
);