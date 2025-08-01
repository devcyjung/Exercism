with mapping (color, code) as (
    values
        ('black', 0),
        ('brown', 1),
        ('red', 2),
        ('orange', 3),
        ('yellow', 4),
        ('green', 5),
        ('blue', 6),
        ('violet', 7),
        ('grey', 8),
        ('white', 9)
), numeric_value (id, value) as (
    select color_code.rowid, (d1.code * 10 + d2.code) * pow(10, d3.code)
    from color_code
    join mapping as d1
    on d1.color = color1
    join mapping as d2
    on d2.color = color2
    join mapping as d3
    on d3.color = color3
), units (unit_word, unit_size) as (
    values
        ('ohms', 1),
        ('kiloohms', 1000),
        ('megaohms', 1000000),
        ('gigaohms', 1000000000)
), representation (id, numeral, unit) as (
    select id, value / u1.unit_size, u1.unit_word
    from numeric_value
    join units u1
    on u1.unit_size = (
        select min(u2.unit_size)
        from units u2
        where u2.unit_size * 1000 > value
    )
)
update color_code
set result = replace(cast(representation.numeral as text), '.0', '') || ' ' || representation.unit
from representation
where color_code.rowid = representation.id;
