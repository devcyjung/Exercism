local Character = {}
local ability, roll_dice, modifier

function Character:new(name)
  self.name = name
  self.strength = ability(roll_dice())
  self.dexterity = ability(roll_dice())
  self.constitution = ability(roll_dice())
  self.intelligence = ability(roll_dice())
  self.wisdom = ability(roll_dice())
  self.charisma = ability(roll_dice())
  self.hitpoints = 10 + modifier(self.constitution)
  return self
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