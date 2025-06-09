typealias CustomSet = Set<Int>

class Set<T>(vararg values: T) {
    private val internalSet: HashSet<T> = HashSet<T>()

    init {
        values.forEach {
            internalSet.add(it)
        }
    }

    fun isEmpty(): Boolean {
        return internalSet.isEmpty()
    }    

    fun isSubset(other: Set<T>): Boolean {
        var isSubset: Boolean = true
        run loop@{
            internalSet.forEach {
                if (!other.contains(it)) {
                    isSubset = false
                    return@loop
                }
            }
        }
        return isSubset
    }

    fun isDisjoint(other: Set<T>): Boolean {
        var isDisjoint: Boolean = true
        run loop@{
            internalSet.forEach {
                if (other.contains(it)) {
                    isDisjoint = false
                    return@loop
                }
            }
        }
        return isDisjoint
    }

    fun contains(other: T): Boolean {
        return internalSet.contains(other)
    }

    fun intersection(other: Set<T>): Set<T> {
        val intersection = Set<T>()
        internalSet.forEach {
            if (other.contains(it)) {
                intersection.add(it)
            }
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
        internalSet.forEach {
            union.add(it)
        }
        other.internalSet.forEach {
            union.add(it)
        }
        return union
    }

    operator fun minus(other: Set<T>): Set<T> {
        val difference = Set<T>()
        internalSet.forEach {
            difference.internalSet.add(it)
        }
        other.internalSet.forEach {
            difference.internalSet.remove(it)
        }
        return difference
    }

}