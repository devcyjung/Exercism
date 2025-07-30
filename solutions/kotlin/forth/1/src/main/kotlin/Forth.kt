class Forth {
    fun evaluate(vararg line: String): List<Int> = run {
        line
            .flatMap { it.split(' ') }
            .filter { it.isNotEmpty() }
            .map { it.uppercase() }
            .forEach { state = state.dispatch(it) }
        stack.toList()
    }

    private val stack = ArrayDeque<Int>()

    private var state: State = State.Interpret(this)

    private sealed class State(val outer: Forth) {
        open fun dispatch(token: String): State = when (token) {
            ":" -> NewWord(outer)
            else -> this
        }

        class NewWord(outer: Forth) : State(outer) {
            override fun dispatch(token: String): State = run {
                require(token.toIntOrNull() == null) { "illegal operation" }
                DefineWord(outer, token)
            }
        }

        class DefineWord(outer: Forth, val word: String) : State(outer) {
            val instructions = mutableListOf<String>()

            private fun isDefined(token: String): Boolean =
                token in outer.definitions || builtinFromString(token) != null || token.toIntOrNull() != null

            override fun dispatch(token: String): State = run {
                when (token) {
                    ";" -> {
                        outer.definitions[word] = instructions
                        Interpret(outer)
                    }

                    else -> {
                        require(isDefined(token)) { "undefined operation" }
                        val predefined = outer.definitions[token]
                        if (predefined != null) {
                            instructions.addAll(predefined)
                        } else {
                            instructions.add(token)
                        }
                        this
                    }
                }
            }
        }

        class Interpret(outer: Forth) : State(outer) {
            override fun dispatch(token: String): State = run {
                val defaultDispatch = super.dispatch(token)
                if (defaultDispatch != this) {
                    return@run defaultDispatch
                }
                outer.interpret(token)
                this
            }
        }
    }

    private val definitions = mutableMapOf<String, List<String>>()

    private fun interpret(token: String): Unit = token.toIntOrNull()?.let {
        stack.addLast(it)
    } ?: run {
        val upper = token.uppercase()
        val instructions = definitions[upper]
        if (instructions != null) {
            instructions.forEach { interpret(it) }
        } else {
            val builtin = builtinFromString(upper)
            require(builtin != null) { "undefined operation" }
            interpretBuiltin(builtin)
        }
    }

    enum class Builtin(val token: String) {
        PLUS("+"), MINUS("-"), MULTIPLY("*"), DIVIDE("/"),
        DUP("DUP"), OVER("OVER"), DROP("DROP"), SWAP("SWAP")
    }

    companion object {
        val BUILTINS = Builtin.values()
        private fun builtinFromString(token: String): Builtin? = BUILTINS.find { it.token == token }
    }

    private fun interpretBuiltin(builtin: Builtin): Unit = when (builtin) {
        Builtin.PLUS -> {
            binaryValidation()
            val last = stack.removeLast()
            stack[stack.lastIndex] += last
        }

        Builtin.MINUS -> {
            binaryValidation()
            val last = stack.removeLast()
            stack[stack.lastIndex] -= last
        }

        Builtin.MULTIPLY -> {
            binaryValidation()
            val last = stack.removeLast()
            stack[stack.lastIndex] *= last
        }

        Builtin.DIVIDE -> {
            binaryValidation()
            require(stack.last() != 0) { "divide by zero" }
            val last = stack.removeLast()
            stack[stack.lastIndex] /= last
        }

        Builtin.SWAP -> {
            binaryValidation()
            stack[stack.lastIndex - 1] = stack[stack.lastIndex]
                .also { stack[stack.lastIndex] = stack[stack.lastIndex - 1] }
        }

        Builtin.DROP -> {
            unaryValidation()
            stack.removeLast()
            Unit
        }

        Builtin.DUP -> {
            unaryValidation()
            stack.addLast(stack.last())
        }

        Builtin.OVER -> {
            binaryValidation()
            stack.addLast(stack[stack.lastIndex - 1])
        }
    }

    private fun unaryValidation() = run {
        require(stack.isNotEmpty()) { "empty stack" }
    }

    private fun binaryValidation() = run {
        unaryValidation()
        require(stack.size > 1) { "only one value on the stack" }
    }
}
