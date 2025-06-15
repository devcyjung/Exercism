class Crumb {
  constructor(value, tree) {
    this.value = value
    this.tree = tree
  }
}

class LeftCrumb extends Crumb {}

class RightCrumb extends Crumb {}

export class Zipper {
  constructor(tree, crumbs) {
    this.tree = tree
    this.crumbs = crumbs
  }

  static fromTree(tree) {
    return new Zipper(tree, [])
  }

  toTree() {
    return this.crumbs.reduce((acc, crumb) => {
      const { value, tree } = crumb
      if (crumb instanceof LeftCrumb) {
        return { value, left: acc, right: tree }
      } else {
        return { value, left: tree, right: acc }
      }
    }, this.tree)
  }

  value() {
    return this.tree.value
  }

  left() {
    const { value, left, right } = this.tree
    if (left === undefined || left === null) {
      return null
    }
    return new Zipper(left, [new LeftCrumb(value, right), ...this.crumbs])
  }

  right() {
    const { value, left, right } = this.tree
    if (right === undefined || right === null) {
      return null
    }
    return new Zipper(right, [new RightCrumb(value, left), ...this.crumbs])
  }

  up() {
    if (this.crumbs.length === 0) {
      return null
    }
    const [crumb, ...rest] = this.crumbs
    const { value, tree } = crumb
    if (crumb instanceof LeftCrumb) {
      return new Zipper({ value, left: this.tree, right: tree }, rest)
    } else {
      return new Zipper({ value, left: tree, right: this.tree }, rest)
    }
  }

  setValue(value) {
    const { left, right } = this.tree
    return new Zipper({ value, left, right }, this.crumbs)
  }

  setLeft(left) {
    const { value, right } = this.tree
    return new Zipper({ value, left, right }, this.crumbs)
  }

  setRight(right) {
    const { value, left } = this.tree
    return new Zipper({ value, left, right }, this.crumbs)
  }
}