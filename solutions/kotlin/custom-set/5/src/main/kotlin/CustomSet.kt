typealias CustomSet = GenericSet<Int>

class GenericSet<T>(val values: MutableSet<T>) {
    constructor(vararg elements: T): this(mutableSetOf(*elements))

    fun isEmpty(): Boolean =
        values.isEmpty()

    fun isSubset(other: GenericSet<*>): Boolean =
        values.all{ elem ->
            other.contains(elem)
        }

    fun isDisjoint(other: GenericSet<*>): Boolean =
        !values.any{ elem ->
            other.contains(elem)
        }

    fun contains(value: @UnsafeVariance T): Boolean =
        values.contains(value)

    fun intersection(other: GenericSet<out T>): GenericSet<T> =
        GenericSet((values intersect other.values).toMutableSet())

    fun add(element: T) = run {
        values += element
    }

    override fun equals(other: Any?): Boolean =
        this === other || (
            other is GenericSet<*>
            && other.values.size == values.size
            && isSubset(other)
        )

    override fun hashCode(): Int = values.hashCode()

    operator fun plus(other: GenericSet<out T>): GenericSet<T> =
        GenericSet((values + other.values).toMutableSet())

    operator fun minus(other: GenericSet<out T>): GenericSet<T> =
        GenericSet((values - other.values).toMutableSet())
}