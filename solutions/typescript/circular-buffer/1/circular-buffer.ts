class EmptyMarker {}

export default class CircularBuffer<T> {
  private buffer: Array<T | EmptyMarker>
  private readCursor: number = 0
  private writeCursor: number = 0
  private static EMPTY = new EmptyMarker()
  
  constructor(initialCapacity: number) {
    this.buffer = Array.from(
      { length: Math.max(1, Math.floor(initialCapacity)) },
      () => CircularBuffer.EMPTY
    )
  }

  write(value: T): void | never {
    const oldValue = this.buffer[this.writeCursor]
    if (oldValue instanceof EmptyMarker) {
      this.buffer[this.writeCursor] = value
      this.writeCursor = (this.writeCursor + 1) % this.buffer.length
    } else {
      throw new BufferFullError()
    }
  }

  read(): T | never {
    const value = this.buffer[this.readCursor]
    if (value instanceof EmptyMarker) {
      throw new BufferEmptyError()
    } else {
      this.buffer[this.readCursor] = CircularBuffer.EMPTY
      this.readCursor = (this.readCursor + 1) % this.buffer.length
      return value
    }
  }

  forceWrite(value: T): void {
    const oldValue = this.buffer[this.writeCursor]
    if (oldValue instanceof EmptyMarker) {
      this.buffer[this.writeCursor] = value
      this.writeCursor = (this.writeCursor + 1) % this.buffer.length
    } else {
      this.read()
      this.write(value)
    }
  }

  clear(): void {
    this.buffer.fill(CircularBuffer.EMPTY)
    this.readCursor = 0
    this.writeCursor = 0
  }
}

export class BufferFullError extends Error {
  constructor(message?: string) {
    super(message || 'Buffer is full')
  }
}

export class BufferEmptyError extends Error {
  constructor(message?: string) {
    super(message || 'Buffer is empty')
  }
}