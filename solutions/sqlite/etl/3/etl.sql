WITH flattened(input, output) AS (
    SELECT
        input,
        json_group_object(object_key, object_value)
    FROM (
        SELECT
            input,
            lower(sublevel.value) AS object_key,
            CAST(toplevel.key AS INTEGER) AS object_value
        FROM
            etl,
            json_each(input) AS toplevel,
            json_each(toplevel.value) AS sublevel
        ORDER BY
            input, object_key
    )
    GROUP BY input
)
UPDATE etl
SET result = (
    SELECT output
    FROM flattened
    WHERE flattened.input = etl.input
);