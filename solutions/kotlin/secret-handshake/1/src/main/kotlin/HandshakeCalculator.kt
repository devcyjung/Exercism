object HandshakeCalculator {
    private val SIGNALS: Array<Signal> = enumValues<Signal>()
    
    fun calculateHandshake(number: Int): List<Signal> = SIGNALS
        .filterIndexed { index, _ -> number and (1 shl index) > 0 }
        .let { if (number and (1 shl SIGNALS.size) > 0) it.reversed() else it }
}
