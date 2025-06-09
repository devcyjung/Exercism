UPDATE etl
SET result = (
    SELECT json_group_object(object_key, object_value)
    FROM (
        SELECT
            lower(sublevel.value) AS object_key,
            CAST(toplevel.key AS INTEGER) AS object_value
        FROM
            json_each(input) AS toplevel,
            json_each(toplevel.value) AS sublevel
        ORDER BY object_key
    )
);