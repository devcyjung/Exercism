pub fn translate(rna: &str) -> Option<Vec<&str>> {
    rna
        .as_bytes()
        .chunks(3)
        .map(std::str::from_utf8)
        .map(Result::ok)
        .map(|opt| opt.and_then(codon))
        .take_while(|opt| !matches!(opt, Some(Codon::Stop)))
        .map(|opt| match opt {
            Some(Codon::Continue(name)) => Some(name),
            _ => None,
        })
        .collect()
}

fn codon(rna: &str) -> Option<Codon> {
    match rna {
        "AUG" => Some(Codon::Continue("Methionine")),
        "UUU" | "UUC" => Some(Codon::Continue("Phenylalanine")),
        "UUA" | "UUG" => Some(Codon::Continue("Leucine")),
        "UCU" | "UCC" | "UCA" | "UCG" => Some(Codon::Continue("Serine")),
        "UAU" | "UAC" => Some(Codon::Continue("Tyrosine")),
        "UGU" | "UGC" => Some(Codon::Continue("Cysteine")),
        "UGG" => Some(Codon::Continue("Tryptophan")),
        "UAA" | "UAG" | "UGA" => Some(Codon::Stop),
        _ => None,
    }
}

enum Codon<'a> {
    Continue(&'a str),
    Stop,
}