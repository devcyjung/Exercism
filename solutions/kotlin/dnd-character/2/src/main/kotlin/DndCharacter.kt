import kotlin.random.Random

class DndCharacter {

    val strength: Int = ability()
    val dexterity: Int = ability()
    val constitution: Int = ability()
    val intelligence: Int = ability()
    val wisdom: Int = ability()
    val charisma: Int = ability()
    val hitpoints: Int = modifier(constitution) + 10

    companion object {

        fun ability(): Int {
            var mini = Random.nextInt(1, 7)
            var total = 0
            var dice: Int
            for (i in 1..3) {
                dice = Random.nextInt(1, 7)
                if (dice < mini) {
                    total += mini
                    mini = dice
                } else {
                    total += dice
                }
            }
            return total
        }

        fun modifier(score: Int): Int {
            return Math.floorDiv(score - 10, 2) 
        }
    }

}
