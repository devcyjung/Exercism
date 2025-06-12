export class Crypto {
  private cipher: string
  
  constructor(plainText: string) {
    const sanitized = Array.from(plainText.toLowerCase())
      .filter(c => ('a' <= c && c <= 'z') || ('0' <= c && c <= '9'))
    let r: number
    let c: number
    const size = sanitized.length
    const root = Math.floor(Math.sqrt(size))
    if (root * root === size) {
      [r, c] = [root, root]
    } else if (root * (root + 1) >= size) {
      [r, c] = [root, root + 1]
    } else {
      [r, c] = [root + 1, root + 1]
    }
    const result = []
    for (let ci = 0; ci < c; ++ci) {
      for (let ri = 0; ri < r; ++ri) {
        const textIdx = ri * c + ci
        if (textIdx < size) {
          result.push(sanitized[textIdx])
        } else {
          result.push(' ')
        }
      }
      if (ci != c - 1) {
        result.push(' ')
      }
    }
    this.cipher = result.join('')
  }

  get ciphertext(): string {
    return this.cipher
  }
}