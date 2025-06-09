#[derive(Debug, PartialEq, Eq)]
pub struct Dna {
    dna: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna {
    rna: String,
}

const DNA: [char; 4] = ['G', 'C', 'T', 'A'];
const RNA: [char; 4] = ['C', 'G', 'A', 'U'];

impl Dna {
    pub fn new(dna: &str) -> Result<Self, usize> {
        dna.chars().enumerate()
            .map(|(i, c)| if DNA.contains(&c) { Ok(c) } else { Err(i) })
            .collect::<Result<_, _>>().map(|dna| Self{ dna })
    }

    pub fn into_rna(self) -> Rna {
        let rna = self.dna.chars()
            .map(|c| RNA[
                DNA.iter().position(|&e| c == e)
                    .expect("Invalid character in Dna")
            ])
            .collect();
        Rna { rna }
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        rna.chars().enumerate()
            .map(|(i, c)| if RNA.contains(&c) { Ok(c) } else { Err(i) })
            .collect::<Result<_, _>>().map(|rna| Self{ rna })
    }
}