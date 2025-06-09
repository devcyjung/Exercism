typealias CustomSet = MySet<Int>

class MySet<T>(vararg values: T) {
    private val internalMySet: MutableSet<T> = mutableSetOf()

    init {
        internalMySet.addAll(values)
    }

    fun isEmpty(): Boolean {
        return internalMySet.isEmpty()
    }

    fun isSubset(other: MySet<T>): Boolean {
        return internalMySet.all {
            other.internalMySet.contains(it)
        }
    }

    fun isDisjoint(other: MySet<T>): Boolean {
        return internalMySet.all {
            !other.internalMySet.contains(it)
        }
    }

    fun contains(other: T): Boolean {
        return internalMySet.contains(other)
    }

    fun intersection(other: MySet<T>): MySet<T> {
        val intersection = MySet<T>()
        internalMySet.filterTo(intersection.internalMySet) {
            other.internalMySet.contains(it)
        }
        return intersection
    }

    fun add(other: T) {
        internalMySet.add(other)
    }

    override fun equals(other: Any?): Boolean {
        return other is MySet<*> && internalMySet.equals(other.internalMySet)
    }

    operator fun plus(other: MySet<T>): MySet<T> {
        val union = MySet<T>()
        union.internalMySet.addAll(internalMySet)
        union.internalMySet.addAll(other.internalMySet)
        return union
    }

    operator fun minus(other: MySet<T>): MySet<T> {
        val difference = MySet<T>()
        difference.internalMySet.addAll(internalMySet)
        difference.internalMySet.removeAll(other.internalMySet)
        return difference
    }
}