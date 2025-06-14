export function degreesOfSeparation(
  familyTree: Record<string, string[]>, personA: string, personB: string
): number {
  const adjacency: Map<string, Set<string>> = new Map()
  for (const [parent, children] of Object.entries(familyTree)) {
    if (!adjacency.has(parent)) {
      adjacency.set(parent, new Set())
    }
    children.forEach(child => {
      adjacency.get(parent)!.add(child)
      if (!adjacency.has(child)) {
        adjacency.set(child, new Set())
      }
      adjacency.get(child)!.add(parent)
      children.forEach(sibling => {
        adjacency.get(child)!.add(sibling)
      })
    })
  }
  const queue = [{ name: personA, distance: 0 }]
  const visited: Set<string> = new Set([personA])
  while (queue.length > 0) {
    const current = queue.shift()!
    const adjacent = adjacency.get(current.name)
    if (adjacent === undefined) {
      return -1
    }
    if (current.name === personB) {
      return current.distance
    }
    for (const immediate of adjacent.values()) {
      if (visited.has(immediate)) {
        continue
      }
      visited.add(immediate)
      queue.push({ name: immediate, distance: current.distance + 1 })
    }
  }
  return -1
}