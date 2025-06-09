/**
 * @classdesc Represents a window dimension
 * @class
 * @param {number} width - The width of the window in pixels
 * @param {number} height - The height of the window in pixels
 * @property {number} width - Default width is 80 pixels, minimum 1
 * @property {number} height - Default height is 60 pixels, minimum 1
 */
function Size(width = 80, height = 60) {
  let width_ = Math.max(1, sanitizeNumber(width, 80))
  let height_ = Math.max(1, sanitizeNumber(height, 60))

  Object.defineProperty(this, 'width', {
    get() {
      return width_
    },
    set(newWidth) {
      width_ = Math.max(sanitizeNumber(newWidth, 80), 1)
    },
    enumerable: true,
  })

  Object.defineProperty(this, 'height', {
    get() {
      return height_
    },
    set(newHeight) {
      height_ = Math.max(sanitizeNumber(newHeight, 60), 1)
    },
    enumerable: true,
  })
}

/**
 * Modifies the width and height of the window
 * @param {number} newWidth - The new width in pixels, minimum 1
 * @param {number} newHeight - The new height in pixels, minimum 1
 * @return {void}
 */
Size.prototype.resize = function(newWidth, newHeight) {
  this.width = sanitizeNumber(newWidth)
  this.height = sanitizeNumber(newHeight)
}

/**
 * @classdesc Represents the upper left corner position of a screen
 * @class
 * @param {number} x - The horizontal position in pixels
 * @param {number} y - The vertical position in pixels
 * @property {number} x - Default horizontal position is 0, minimum 0
 * @property {number} y - Default vertical position is 0, minimum 0
 */
function Position(x = 0, y = 0) {
  let x_ = Math.max(sanitizeNumber(x), 0)
  let y_ = Math.max(sanitizeNumber(y), 0)

  Object.defineProperty(this, 'x', {
    get() {
      return x_
    },
    set(newX) {
      x_ = Math.max(sanitizeNumber(newX), 0)
    },
    enumerable: true,
  })

  Object.defineProperty(this, 'y', {
    get() {
      return y_
    },
    set(newY) {
      y_ = Math.max(sanitizeNumber(newY), 0)
    },
    enumerable: true,
  })
}

/**
 * Moves the position of the upper left corner by modifying x and y
 * @param {number} newX - The new horizontal position in pixels, minimum 0
 * @param {number} newY - The new vertical position in pixels, minimum 0
 * @return {void}
 */
Position.prototype.move = function(newX, newY) {
  this.x = sanitizeNumber(newX)
  this.y = sanitizeNumber(newY)
}

/**
 * @classdesc Represents a program window within a screen
 * @class
 * @property {Size} screenSize - The screen size. Always (width, height) = (800, 600)
 * @property {Size} size - The window size, (width, height) = (80, 60) for new instance
 * @property {Position} position - The window's top left corner position, (x, y) = (0, 0) for new instance
 */
class ProgramWindow {
  static screenSize = Object.freeze(new Size(800, 600))
  #size= new Size()
  #position= new Position()

  get screenSize() {
    return ProgramWindow.screenSize
  }

  get size() {
    return new Size(this.#size.width, this.#size.height)
  }

  get position() {
    return new Position(this.#position.x, this.#position.y)
  }

  /**
   * Modifies the size of program window, as long as the screen size permits
   * @param {Size} newSize - The new Size object, minimum (1, 1)
   * @return {void}
   */
  resize(newSize) {
    if (!newSize instanceof Size) {
      throw new TypeError(`${newSize} is not of type Size`)
    }
    const newWidth = Math.min(newSize.width, this.screenSize.width - this.#position.x)
    const newHeight = Math.min(newSize.height, this.screenSize.height - this.#position.y)
    this.#size.resize(newWidth, newHeight)
  }

  /**
   * Modifies the position of top left corner, as long as it's within the screen
   * @param {Position} newPosition - The new Position object, minimum (0, 0)
   * @return {void}
   */
  move(newPosition) {
    if (!newPosition instanceof Position) {
      throw new TypeError(`${newPosition} is not of type Position`)
    }
    const newX = Math.min(newPosition.x, this.screenSize.width - this.#size.width)
    const newY = Math.min(newPosition.y, this.screenSize.height - this.#size.height)
    this.#position.move(newX, newY)
  }
}

/**
 * Changes the program window
 * @param {ProgramWindow} programWindow - The ProgramWindow object to modify
 * @return {ProgramWindow} - The modified result (width, height, x, y) = (400, 300, 100, 150)
 */
function changeWindow(programWindow) {
  if (!programWindow instanceof ProgramWindow) {
    throw new TypeError(`${programWindow} is not of type ProgramWindow`)
  }
  programWindow.resize(new Size(400, 300))
  programWindow.move(new Position(100, 150))
  return programWindow
}

/**
 * Sanitizes the arbitrary input into Number
 * param {*} input - Arbitrary input
 * param {number} [defaultNumber] - default number to use for sanitization
 * return {number} - The sanitized output
 */
function sanitizeNumber(input, defaultNumber = 0) {
  const number = Number(input)
  return Number.isFinite(number) ? number : defaultNumber
}

export {
  Size,
  Position,
  ProgramWindow,
  changeWindow,
}