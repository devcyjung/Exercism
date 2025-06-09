enum class Classification {
    DEFICIENT, PERFECT, ABUNDANT
}

fun classify(naturalNumber: Int): Classification {
    require(naturalNumber > 0) {
        "Must provide a positive number"
    }
    return (1..<naturalNumber).filter{ divisor ->
        naturalNumber % divisor == 0
    }.sum().let{ aliquotSum ->
        when {
            aliquotSum == naturalNumber -> Classification.PERFECT
            aliquotSum < naturalNumber -> Classification.DEFICIENT
            else -> Classification.ABUNDANT
        }   
    }
}