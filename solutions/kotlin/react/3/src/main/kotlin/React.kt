import java.util.IdentityHashMap

class Reactor<T>() {
    private val dag = mutableListOf<Cell<T>>()

    private fun invokeChain(id: Int) {
        dag.asSequence().drop(id + 1).forEach{ cell ->
            cell.value
        }
    }
    
    fun InputCell(initial: T): InputCell<T> {
        val id = dag.size
        val input = InputCell(id = id, initial = initial) {
            invokeChain(id)
        }
        dag.add(input)
        return input
    }

    fun ComputeCell(vararg cells: Cell<T>, compute: (List<T>) -> T): ComputeCell<T> {
        val id = dag.size
        val com = ComputeCell(id = id) {
            compute(cells.map(Cell<T>::value))
        }
        dag.add(com)
        return com
    }
}

interface Subscription {
    fun cancel()
}

interface Cell<T> {
    val id: Int
    val value: T
}

class InputCell<T>(
    override val id: Int, initial: T, val runner: () -> Unit
): Cell<T> {
    override var value: T = initial
        set(newValue) {
            if (field == newValue) {
                return
            }
            field = newValue
            runner()
        }
}

class ComputeCell<T>(override val id: Int, val compute: () -> T): Cell<T> {
    private val registry = IdentityHashMap<Canceler<T>, (T) -> Unit>()
    override var value: T = compute()
        get() {
            val newValue = compute()
            if (field == newValue) {
                return field
            }
            field = newValue
            registry.forEach{ (_, callback) ->
                callback(field)
            }
            return field
        }

    fun addCallback(callback: (T) -> Unit): Canceler<T> {
        val newCanceler = Canceler(callback, registry)
        registry.put(newCanceler, callback)
        return newCanceler
    }
}

data class Canceler<T>(
    val callback: (T) -> Unit, val registry: MutableMap<Canceler<T>, (T) -> Unit>
): Subscription {
    override fun cancel() {
        registry.remove(this)
    }
}