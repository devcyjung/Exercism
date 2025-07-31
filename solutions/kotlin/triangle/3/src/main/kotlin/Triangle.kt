class Triangle<out T : Number>(val a: T, val b: T, val c: T) {
    val sides by lazy {
        val inputs = listOf(a, b, c)
        Array(3) { inputs[it].toDouble() }.sorted()
    }
    
    init {
        require(sides[0] + sides[1] > sides[2]) { "Sides violate triangle inequality" }
    }

    val isEquilateral: Boolean = sides[0] == sides[2]
    val isIsosceles: Boolean = sides[0] == sides[1] || sides[1] == sides[2]
    val isScalene: Boolean = !isEquilateral && !isIsosceles
}