export class DnDCharacter {
  public readonly strength = DnDCharacter.generateAbilityScore()
  public readonly dexterity = DnDCharacter.generateAbilityScore()
  public readonly constitution = DnDCharacter.generateAbilityScore()
  public readonly intelligence = DnDCharacter.generateAbilityScore()
  public readonly wisdom = DnDCharacter.generateAbilityScore()
  public readonly charisma = DnDCharacter.generateAbilityScore()
  public readonly hitpoints = DnDCharacter.getModifierFor(this.constitution) + 10

  public static generateAbilityScore(): number {
    let mini = DnDCharacter.diceRoll()
    let dice: number
    let total = 0
    for (let i = 0; i < 3; ++i) {
      dice = DnDCharacter.diceRoll()
      if (dice < mini) {
        total += mini
        mini = dice
      } else {
        total += dice
      }
    }
    return total
  }

  private static diceRoll(): number {
    return Math.random() * 6 + 1
  }

  public static getModifierFor(abilityValue: number): number {
    return Math.floor((abilityValue - 10) / 2)
  }
}
