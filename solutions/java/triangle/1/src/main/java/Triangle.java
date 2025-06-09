enum TriangleType {
    ISOSCELES, EQUILATERAL, SCALENE
}

class Triangle {
    private final TriangleType triangleType;
    
    Triangle(double side1, double side2, double side3) throws TriangleException {
        long s1 = Double.doubleToLongBits(side1);
        long s2 = Double.doubleToLongBits(side2);
        long s3 = Double.doubleToLongBits(side3);
        if (s1 > s2) {
            s1 ^= s2;
            s2 ^= s1;
            s1 ^= s2;
        }
        if (s1 > s3) {
            s1 ^= s3;
            s3 ^= s1;
            s1 ^= s3;
        }
        if (s2 > s3) {
            s2 ^= s3;
            s3 ^= s2;
            s2 ^= s3;
        }
        double d1 = Double.longBitsToDouble(s1);
        double d2 = Double.longBitsToDouble(s2);
        double d3 = Double.longBitsToDouble(s3);
        if (d3 >= d1 + d2) {
            throw new TriangleException();
        }
        if (d1 == d3) {
            triangleType = TriangleType.EQUILATERAL;
        } else if (d2 == d1 || d2 == d3) {
            triangleType = TriangleType.ISOSCELES;
        } else {
            triangleType = TriangleType.SCALENE;
        }
    }

    boolean isEquilateral() {
        return triangleType == TriangleType.EQUILATERAL; 
    }

    boolean isIsosceles() {
        return triangleType == TriangleType.ISOSCELES
            || triangleType == TriangleType.EQUILATERAL;
    }

    boolean isScalene() {
        return triangleType == TriangleType.SCALENE;
    }

}