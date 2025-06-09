pub struct Triangle(u64, u64, u64);

impl Triangle {
    pub fn build(sides: [u64; 3]) -> Option<Self> {
        let mut sides = sides;
        sides.sort();
        if sides[2] <= sides[0] + sides[1] && sides[0] > 0 {
            return Some(Self(sides[2], sides[1], sides[0]));
        }
        None
    }

    pub fn is_equilateral(&self) -> bool {
        self.0 == self.2
    }

    pub fn is_scalene(&self) -> bool {
        !self.is_equilateral() && !self.is_isosceles()
    }

    pub fn is_isosceles(&self) -> bool {
        self.0 == self.1 || self.1 == self.2
    }
}
