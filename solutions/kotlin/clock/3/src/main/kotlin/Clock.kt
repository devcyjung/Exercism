import java.util.Objects

typealias Hour = Int
typealias Minute = Int

class Clock(var hours: Int, var minutes: Int) {
    val normalized: Pair<Hour, Minute>
        get() = (hours * 60 + minutes).mod(24 * 60).let { (it / 60) to (it % 60) }
    
    fun subtract(minutes: Int) = run {
        this.minutes -= minutes
    }

    fun add(minutes: Int) = run {
        this.minutes += minutes
    }

    override fun toString(): String = normalized.let { (h, m) ->
        "%02d:%02d".format(h, m)
    }

    override fun equals(other: Any?): Boolean = other is Clock && other.normalized == normalized

    override fun hashCode(): Int = normalized.hashCode()
}