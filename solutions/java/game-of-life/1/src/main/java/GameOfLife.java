import java.util.ArrayList;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class GameOfLife {
    
    public static int[][] tick(int[][] matrix){
        final int nrow = matrix.length;
        final int ncol = matrix.length == 0 ? 0 : matrix[0].length;
        var toZero = new ArrayList<Pair>(ncol * nrow / 2);
        var toOne = new ArrayList<Pair>(ncol * nrow / 2);
        IntStream.range(0, nrow).boxed()
            .flatMap(i -> IntStream.range(0, ncol).mapToObj(j -> new Pair(i, j)))
            .forEach(p -> {
                switch (aliveNeighbors(p.i, p.j, nrow, ncol, matrix)) {
                case 2:
                    if (matrix[p.i][p.j] == 0) {
                        toZero.add(p);
                    }
                    break;
                case 3:
                    toOne.add(p);
                    break;
                default:
                    toZero.add(p);                
                }
            });
        toZero.stream().forEach(p -> {
            matrix[p.i][p.j] = 0;
        });
        toOne.stream().forEach(p -> {
            matrix[p.i][p.j] = 1;
        });
        return matrix;
    }

    private static int aliveNeighbors(int i, int j, int nrow, int ncol, int[][] matrix) {
        var count = IntStream.range(-1, 2).boxed()
            .flatMap(di -> IntStream.range(-1, 2).mapToObj(dj -> new Pair(i + di, j + dj)))
            .filter(
                p -> p.i >= 0 && p.j >= 0 && p.i < nrow && p.j < ncol 
                && matrix[p.i][p.j] == 1 && (p.i != i || p.j != j)
            )
            .count();
        return Math.toIntExact(count);
    }

    record Pair(int i, int j) {}
}