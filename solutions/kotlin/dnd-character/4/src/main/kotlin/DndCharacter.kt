import kotlin.random.nextInt
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
        fun ability(): Int =
            generateSequence{
                Random.nextInt(1..6)
            }.take(4).sortedDescending().take(3).sum()

        fun modifier(score: Int): Int =
            (score - 10).floorDiv(2)
    }
}