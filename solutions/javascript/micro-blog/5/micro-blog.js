const segmenter = new Intl.Segmenter()

export const truncate = input => Iterator.from(segmenter.segment(input)).take(5).map(s => s.segment).toArray().join("")

// WRONG : produces "👍🏽👍🏽👍"
console.log(Array.from("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽").slice(0, 5).join(""))
// CORRECT: produces "👍🏽👍🏽👍🏽👍🏽👍🏽"
console.log(segmenter.segment("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽")[Symbol.iterator]().take(5).map(s => s.segment).toArray().join(""))