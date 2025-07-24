local Anagram = {}

local function sort_string(str)
  local tbl = {}
  for i = 1, #str do
    table.insert(tbl, str:sub(i, i))
  end
  table.sort(tbl)
  return table.concat(tbl)
end

function Anagram:new(word)
  local lower_word = word:lower()
  return setmetatable({
      word = lower_word,
      sorted = sort_string(lower_word)
    }, {__index = self})
end

function Anagram:match(candidates)
  local anagrams = {}
  for _, candidate in ipairs(candidates) do
    local lower_candidate = candidate:lower()
    if self.word ~= lower_candidate and self.sorted == sort_string(lower_candidate) then
      table.insert(anagrams, candidate)
    end
  end
  return anagrams
end

return Anagram
