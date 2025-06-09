export class DnDCharacter {
  strength: number
  dexterity: number
  constitution: number
  intelligence: number
  wisdom: number
  charisma: number
  hitpoints: number
  
  constructor() {
    this.strength = DnDCharacter.generateAbilityScore()
    this.dexterity = DnDCharacter.generateAbilityScore()
    this.constitution = DnDCharacter.generateAbilityScore()
    this.intelligence = DnDCharacter.generateAbilityScore()
    this.wisdom = DnDCharacter.generateAbilityScore()
    this.charisma = DnDCharacter.generateAbilityScore()
    this.hitpoints = DnDCharacter.getModifierFor(this.constitution) + 10
  }
  
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
