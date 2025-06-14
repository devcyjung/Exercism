export class ValueError extends Error {
  constructor(message?: string) {
    super(message || 'Bank account error')
  }
}

export class BankAccount {
  private isOpen: boolean = false
  #balance: number = 0
  
  constructor() {}

  open(): void | never {
    if (this.isOpen) {
      throw new ValueError('Account is already open')
    }
    this.isOpen = true
  }

  close(): void {
    if (!this.isOpen) {
      throw new ValueError('Account is already closed')
    }
    this.isOpen = false
    this.#balance = 0
  }

  deposit(amount: number): void | never {
    if (!this.isOpen) {
      throw new ValueError('Account is closed')
    }
    if (amount < 0) {
      throw new ValueError('Cannot deposit negative amount')
    }
    this.#balance += amount
  }

  withdraw(amount: number): void | never {
    if (!this.isOpen) {
      throw new ValueError('Account is closed')
    }
    if (amount < 0) {
      throw new ValueError('Cannot withdraw negative amount')
    }
    if (this.#balance < amount) {
      throw new ValueError('Insufficient balance to withdraw from')
    }
    this.#balance -= amount
  }

  get balance(): number {
    if (!this.isOpen) {
      throw new ValueError('Account is closed')
    }
    return this.#balance
  }
}
