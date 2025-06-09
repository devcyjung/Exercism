class EmptyBufferException: IllegalStateException()
class BufferFullException: IllegalStateException()

class CircularBuffer<T>(val capacity: Int) {
    init {
        require(capacity >= 0) {
            "Capacity must be non-negative"
        }
    }

    @Suppress("UNCHECKED_CAST")
    private val buffer = arrayOfNulls<Any?>(capacity) as Array<T?>
    private var readCursor: Int = 0
    private var writeCursor: Int = 0

    fun read() : T {
        val value = buffer.get(readCursor)
        if (value == null) {
            throw EmptyBufferException()
        }
        buffer.set(readCursor, null)
        readCursor = (readCursor + 1) % buffer.size
        return value
    }

    fun write(value: T) {
        if (buffer.get(writeCursor) != null) {
            throw BufferFullException()
        }
        buffer.set(writeCursor, value)
        writeCursor = (writeCursor + 1) % buffer.size
    }

    fun overwrite(value: T) {
        if (buffer.get(writeCursor) != null) {
            readCursor = (readCursor + 1) % buffer.size
        }
        buffer.set(writeCursor, value)
        writeCursor = (writeCursor + 1) % buffer.size
    }

    fun clear() {
        buffer.fill(null)
        readCursor = 0
        writeCursor = 0
    }
}