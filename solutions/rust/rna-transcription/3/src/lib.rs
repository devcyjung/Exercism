#[derive(Debug, PartialEq, Eq)]
pub enum DNA {
    A, C, G, T,
}

#[derive(Debug, PartialEq, Eq)]
pub enum RNA {
    U, G, C, A,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Dna {
    dna: Vec<DNA>,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna {
    rna: Vec<RNA>,
}

impl TryFrom<char> for DNA {
    type Error = ();
    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            'A' => Ok(Self::A),
            'C' => Ok(Self::C),
            'G' => Ok(Self::G),
            'T' => Ok(Self::T),
            _   => Err(()),
        }
    }
}

impl TryFrom<char> for RNA {
    type Error = ();
    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            'C' => Ok(Self::C),
            'G' => Ok(Self::G),
            'A' => Ok(Self::A),
            'U' => Ok(Self::U),
            _   => Err(()),
        }
    }
}

impl From<DNA> for RNA {
    fn from(value: DNA) -> Self {
        match value {
            DNA::G => Self::C,
            DNA::C => Self::G,
            DNA::T => Self::A,
            DNA::A => Self::U,
        }
    }
}

impl Dna {
    pub fn new(dna: &str) -> Result<Self, usize> {
        dna.chars().enumerate()
            .map(|(i, c)| DNA::try_from(c).map_err(|_| i))
            .collect::<Result<Vec<_>, _>>()
            .map(|dna| Self{ dna })
    }

    pub fn into_rna(self) -> Rna {
        let rna = self.dna.into_iter().map(Into::into).collect();
        Rna{ rna }
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        rna.chars().enumerate()
            .map(|(i, c)| RNA::try_from(c).map_err(|_| i))
            .collect::<Result<Vec<_>, _>>()
            .map(|rna| Self{ rna })
    }
}