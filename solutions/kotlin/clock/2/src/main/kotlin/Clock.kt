import java.util.Objects

class Clock(var hours: Int, var minutes: Int) {
    val realTimeInMinute: Int
        get() = (hours * 60 + minutes).mod(24 * 60)

    val realHours: Int
        get() = realTimeInMinute / 60

    val realMinutes: Int
        get() = realTimeInMinute % 60
    
    fun subtract(minutes: Int) {
        this.minutes -= minutes
    }

    fun add(minutes: Int) {
        this.minutes += minutes
    }

    override fun toString(): String = "${realHours.toString().padStart(2, '0')}:${realMinutes.toString().padStart(2, '0')}"

    override fun equals(other: Any?): Boolean = other is Clock && other.realHours == realHours && other.realMinutes == realMinutes

    override fun hashCode(): Int = Objects.hash(realHours, realMinutes)
}