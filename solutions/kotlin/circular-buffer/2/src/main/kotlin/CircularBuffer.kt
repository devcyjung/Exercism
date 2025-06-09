import kotlin.collections.ArrayDeque

class EmptyBufferException: IllegalStateException("Empty Buffer")

class BufferFullException: IllegalStateException("Full Buffer")

class CircularBuffer<T>(val size: Int) {

    private val queue = ArrayDeque<T>(size)
    private var len = 0
    private var readFrom = 0
    private var writeTo = 0

    fun read(): T {
        if (isEmpty) {
            throw EmptyBufferException()
        }
        return queue
            .get(readFrom)
            .also {
                --len
                ++readFrom
                readFrom %= size
            }
    }
    
    fun write(value: T) {
        if (isFull) {
            throw BufferFullException()
        }
        if (queue.size <= writeTo) {
            queue.add(writeTo, value)
        } else {
            queue.set(writeTo, value)
        }
        ++len
        ++writeTo
        writeTo %= size
    }

    fun overwrite(value: T) {
        if (isFull) {
            read()
        }
        write(value)
    }
    
    fun clear() {
        len = 0
        readFrom = 0
        writeTo = 0
    }

    private val isEmpty
        get() = len == 0

    private val isFull
        get() = len == size

}