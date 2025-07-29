fun <T> List<T>.customAppend(list: List<T>): List<T> = this + list

fun List<*>.customConcat(): List<*> = this.flatMap {
    when(it) {
        is List<*> -> if (it.isEmpty()) { it } else { it.customConcat() }
        else -> listOf(it)
    }
}

fun <T> List<T>.customFilter(predicate: (T) -> Boolean): List<T> = this.filter(predicate)

val List<*>.customSize: Int get() = this.size

fun <T, U> List<T>.customMap(transform: (T) -> U): List<U> = this.map(transform)

fun <T, U> List<T>.customFoldLeft(initial: U, f: (U, T) -> U): U = this.fold(initial, f)

fun <T, U> List<T>.customFoldRight(initial: U, f: (T, U) -> U): U = this.foldRight(initial, f)

fun <T> List<T>.customReverse(): List<T> = this.reversed()