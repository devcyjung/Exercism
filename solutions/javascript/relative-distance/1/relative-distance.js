export const degreesOfSeparation = (familyTree, personA, personB) => {
  const graph = makeFamilyGraph(familyTree)
  const queue = []
  const visited = new Set()
  queue.push({name: personA, degree: 0})
  visited.add(personA)
  while (queue.length > 0) {
    const cur = queue.shift()
    const curName = cur.name
    if (curName === personB) {
      return cur.degree
    }
    const nextDegree = cur.degree + 1
    for (const adjacent of graph.get(curName)) {
      if (adjacent === personB) {
        return nextDegree
      }
      if (visited.has(adjacent)) {
        continue
      }
      queue.push({name: adjacent, degree: nextDegree})
      visited.add(adjacent)
    }
  }
  return -1
}

function makeFamilyGraph(familyTree) {
  const graph = new Map()
  for (const [parent, children] of Object.entries(familyTree)) {
    if (!graph.has(parent)) {
      graph.set(parent, new Set())
    }
    const parentSet = graph.get(parent)
    children.forEach(child => {
      parentSet.add(child)
    })
    children.forEach(child => {
      if (!graph.has(child)) {
        graph.set(child, new Set())
      }
      const childSet = graph.get(child)
      childSet.add(parent)
      children.forEach(sibling => {
        childSet.add(sibling)
      })
    })
  }
  return graph
}