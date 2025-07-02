local ability, roll_dice, modifier
local Character = { ability = ability, roll_dice = roll_dice, modifier = modifier }

function Character:new(name)
  local new = {}
  setmetatable(new, self)
  new.__index = self
  new.name = name
  new.strength = ability(roll_dice())
  new.dexterity = ability(roll_dice())
  new.constitution = ability(roll_dice())
  new.intelligence = ability(roll_dice())
  new.wisdom = ability(roll_dice())
  new.charisma = ability(roll_dice())
  new.hitpoints = 10 + modifier(new.constitution)
  return new
end

function ability(scores)
  table.sort(scores,
    function (a, b)
      return a > b
    end
  )
  local sum = 0
  for i = 1, 3 do
    sum = sum + scores[i]
  end
  return sum
end

function roll_dice()
  local dice = {}
  for i = 1, 4 do
    table.insert(dice, math.random(1, 6))
  end
  return dice
end

function modifier(input)
  return math.floor((input - 10) / 2)
end

return { Character = Character, ability = ability, roll_dice = roll_dice, modifier = modifier }