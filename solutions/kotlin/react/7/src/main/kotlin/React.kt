import kotlin.properties.Delegates

class Reactor<T>() {
    fun interface Subscription {
        fun cancel()
    }

    open inner class InputCell(initialValue: T) {
        private val onChangeCallbacks: MutableSet<(T) -> Any?> = mutableSetOf()

        var value: T by Delegates.observable(initialValue) { _, old, new ->
            if (old != new) {
                onChangeCallbacks.forEach { callback -> callback(new) }
            }
        }

        open fun compute(): T = value

        fun addCallback(callback: (T) -> Unit): Subscription = Subscription {
            onChangeCallbacks.remove(callback)
        }.also { onChangeCallbacks.add(callback) }
    }

    inner class ComputeCell(vararg val cells: InputCell, val computeFn: (values: List<T>) -> T) :
        InputCell(computeFn(cells.map { it.compute() })) {
        init {
            cells.forEach { it.addCallback { value = compute() } }
        }

        override fun compute(): T = computeFn(cells.map { it.compute() })
    }
}
