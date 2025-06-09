export function toRna(dna:string):string {
  const rna:string[] = [];
  dna.split('').forEach((e:string) => {
    let s:string="";
    switch (e){
      case "G": s="C"; break;
      case "C": s="G"; break;
      case "T": s="A"; break;
      case "A": s="U"; break;
      default: throw new Error("Invalid input DNA.");
    }
    rna.push(s);
  });
  return rna.join("");
}
