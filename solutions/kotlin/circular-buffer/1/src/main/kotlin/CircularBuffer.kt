import kotlin.collections.ArrayDeque

class EmptyBufferException: IllegalStateException("Empty Buffer")

class BufferFullException: IllegalStateException("Full Buffer")

class CircularBuffer<T>(val size: Int) {

    val queue = ArrayDeque<T>(size)
    var len = 0
    var readFrom = 0
    var writeTo = 0

    fun read(): T = if (isEmpty())
        throw EmptyBufferException()
        else queue
            .get(readFrom)
            .also {
                --len
                ++readFrom
                readFrom %= size
            }

    fun write(value: T) = if (isFull())
        throw BufferFullException()
        else queue
            .let {
                if (it.size <= writeTo) {
                    it.add(writeTo, value)
                } else {
                    it.set(writeTo, value)
                }
            }
            .also {
                ++len
                ++writeTo
                writeTo %= size
            }

    fun overwrite(value: T) = this
        .also {
            if (isFull()) {
                read()
            }
            write(value)
        }

    fun clear() = this
        .also {
            len = 0
            readFrom = 0
            writeTo = 0
        }

    fun isEmpty(): Boolean = len == 0

    fun isFull(): Boolean = len == size

}