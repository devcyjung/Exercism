/// <reference path="./global.d.ts" />
// @ts-check

export function cookingStatus(timer) {
  switch (typeof timer) {
    case 'number':
      return (timer === 0) ? 'Lasagna is done.' : 'Not done, please wait.'
    default:
      return 'You forgot to set the timer.'
  }
}

export function preparationTime(layers, average = 2) {
  return layers.length * average
}

export function quantities(layers) {
  return {
    noodles: 50 * layers.filter(e => e === 'noodles').length,
    sauce: 0.2 * layers.filter(e => e === 'sauce').length,
  }
}

export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList.at(-1))
}

export function scaleRecipe(recipe, portions) {
  return Object.fromEntries(Object.entries(recipe).map(([k, v]) => [k, portions / 2 * v]))
}