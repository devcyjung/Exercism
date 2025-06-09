typealias CustomSet = Set<Int>

class Set<T>(vararg values: T) {
    private val internalSet: HashSet<T> = HashSet<T>()

    init {
        internalSet.addAll(values)
    }

    fun isEmpty(): Boolean {
        return internalSet.isEmpty()
    }    

    fun isSubset(other: Set<T>): Boolean {
        return internalSet.all {
            other.internalSet.contains(it)
        }
    }

    fun isDisjoint(other: Set<T>): Boolean {
        return internalSet.all {
            !other.internalSet.contains(it)
        }
    }

    fun contains(other: T): Boolean {
        return internalSet.contains(other)
    }

    fun intersection(other: Set<T>): Set<T> {
        val intersection = Set<T>()
        internalSet.filterTo(intersection.internalSet) {
            other.internalSet.contains(it)
        }
        return intersection
    }

    fun add(other: T) {
        internalSet.add(other)
    }

    override fun equals(other: Any?): Boolean {
        return other is Set<*> && internalSet.equals(other.internalSet)
    }

    operator fun plus(other: Set<T>): Set<T> {
        val union = Set<T>()
        union.internalSet.addAll(internalSet)
        union.internalSet.addAll(other.internalSet)
        return union
    }

    operator fun minus(other: Set<T>): Set<T> {
        val difference = Set<T>()
        difference.internalSet.addAll(internalSet)
        difference.internalSet.removeAll(other.internalSet)
        return difference
    }
}