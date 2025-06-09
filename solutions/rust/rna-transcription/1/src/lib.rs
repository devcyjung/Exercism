#[derive(Debug, PartialEq, Eq)]
pub enum DNA {
    A, C, G, T,
}

#[derive(Debug, PartialEq, Eq)]
pub enum RNA {
    U, G, C, A,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Dna{
    dna: Vec<DNA>,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna{
    rna: Vec<RNA>,
}

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        let dna = dna.chars().enumerate()
            .try_fold(Vec::new(), |mut acc, (i, ch)| match ch {
                'A' => {
                    acc.push(DNA::A);
                    Ok(acc)
                },
                'C' => {
                    acc.push(DNA::C);
                    Ok(acc)
                }
                'G' => {
                    acc.push(DNA::G);
                    Ok(acc)
                },
                'T' => {
                    acc.push(DNA::T);
                    Ok(acc)
                },
                _   => Err(i),
            })?;
        Ok(Self{ dna })
    }

    pub fn into_rna(self) -> Rna {
        Rna {
            rna: self.dna.iter().map(|d| {
                match d {
                    DNA::A => RNA::U,
                    DNA::C => RNA::G,
                    DNA::G => RNA::C,
                    DNA::T => RNA::A,
                }
            }).collect(),
        }
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        let rna = rna.chars().enumerate()
            .try_fold(Vec::new(), |mut acc, (i, ch)| match ch {
                'C' => {
                    acc.push(RNA::C);
                    Ok(acc)
                },
                'G' => {
                    acc.push(RNA::G);
                    Ok(acc)
                }
                'A' => {
                    acc.push(RNA::A);
                    Ok(acc)
                },
                'U' => {
                    acc.push(RNA::U);
                    Ok(acc)
                },
                _   => Err(i),
            })?;
        Ok(Self { rna })
    }
}