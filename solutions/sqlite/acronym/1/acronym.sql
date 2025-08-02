with recursive trim(id, acc, remaining) as (
    select rowid, '', ' ' || upper(replace(phrase, '-', ' '))
    from acronym
    union all
    select
            id,
            case
                when length(ltrim(remaining)) != length(remaining) then
                    acc || ' '
                when substr(remaining, 1, 1) between 'A' and 'Z' then
                    acc || substr(remaining, 1, 1)
                else
                    acc
            end,
            substr(remaining, 2)
    from trim
    where length(remaining) > 0
), abbr(id, acc, previous, remaining) as (
    select id, '', '', acc
    from trim
    where length(remaining) == 0
    union all
    select
        id,
        iif(
                substr(remaining, 1, 1) between 'A' and 'Z'
                and substr(previous, -1, 1) == ' ',
                acc || substr(remaining, 1, 1),
                acc
            ),
        previous || substr(remaining, 1, 1),
        substr(remaining, 2)
    from abbr
    where length(remaining) > 0
)
update acronym
set result = acc
from abbr
where id == rowid and length(remaining) == 0;
