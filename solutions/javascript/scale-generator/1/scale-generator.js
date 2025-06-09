export class Scale {
  sharps = ['A', 'A#', 'B', 'C', 'C#', 'D', 'D#', 'E', 'F', 'F#', 'G', 'G#']
  flats = ['A', 'Bb', 'B', 'C', 'Db', 'D', 'Eb', 'E', 'F', 'Gb', 'G', 'Ab']
  
  constructor(tonic) {
    this.isFlat = (tonic.length === 2 && tonic[1] === "b")
      || ['F', 'd', 'g', 'c', 'f'].includes(tonic)
    this.isSharp = !this.isFlat
    this.series = this.isSharp ? this.sharps : this.flats
    this.tonic = tonic.slice(0, 1).toUpperCase() + tonic.slice(1)
    this.tonicIndex = this.series.indexOf(this.tonic)
  }

  chromatic() {
    const chr = Array.from({ length: 12 })
    chr.forEach((_, i, a) => {
      a[i] = this.series[(this.tonicIndex + i) % 12]
    })
    return chr
  }

  interval(intervals) {
    const chr = Array.from({ length: intervals.length + 1 })
    chr[0] = this.series[this.tonicIndex]
    let current = this.tonicIndex
    intervals.split("").forEach((v, i) => {
      switch (v) {
        case "M":
          current += 2
          break
        case "m":
          ++current
          break
        case "A":
          current += 3
          break
      }
      current %= 12
      chr[i + 1] = this.series[current]
    })
    return chr
  }
}
