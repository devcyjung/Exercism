class ListNode<TElement> {
  public value: TElement
  public next: ListNode<TElement> | undefined
  public prev: ListNode<TElement> | undefined
  constructor(element: TElement) {
    this.value = element
    this.next = undefined
    this.prev = undefined
  }  
}

export class LinkedList<TElement> {
  protected length: number = 0
  protected head: ListNode<TElement> | undefined = undefined
  protected tail: ListNode<TElement> | undefined = undefined
  
  public push(element: TElement) {
    ++this.length
    const newNode = new ListNode<TElement>(element)
    if (this.tail === undefined) {
      this.tail = newNode
      this.head = this.tail
      return
    }
    this.tail.next = newNode
    newNode.prev = this.tail
    this.tail = this.tail.next
  }

  public pop(): TElement {
    if (this.tail === undefined) {
      throw new Error("pop called on empty list")
    }
    --this.length
    const returnNode = this.tail
    this.tail = this.tail.prev
    this.tail && (this.tail.next = undefined)
    if (this.tail === undefined) {
      this.head = undefined
    }
    return returnNode.value
  }

  public shift(): TElement {
    if (this.head === undefined) {
      throw new Error("shift called on empty list")
    }
    --this.length
    const returnNode = this.head
    this.head = this.head.next
    this.head && (this.head.prev = undefined)
    if (this.head === undefined) {
      this.tail = undefined
    }
    return returnNode.value
  }

  public unshift(element: TElement) {
    ++this.length
    const newNode = new ListNode<TElement>(element)
    if (this.head === undefined) {
      this.head = newNode
      this.tail = this.head
      return
    }
    this.head.prev = newNode
    newNode.next = this.head
    this.head = this.head.prev
  }

  public delete(element: TElement) {
    let node = this.head
    let index = 0
    while (node) {
      if (element === node.value) {
        node.prev && (node.prev.next = node.next)
        node.next && (node.next.prev = node.prev)
        if (index === 0) {
          this.head = undefined
        }
        if (index === this.length - 1) {
          this.tail = undefined
        }
        --this.length
        return
      }
      node = node.next
      ++index
    }
  }

  public count(): number {
    return this.length
  }
}
