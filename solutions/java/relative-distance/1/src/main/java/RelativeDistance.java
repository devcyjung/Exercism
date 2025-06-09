import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class RelativeDistance {
    private final Map<String, List<String>> familyGraph;
    RelativeDistance(Map<String, List<String>> familyTree) {
        Map<String, List<String>> graph = new HashMap<>();
        familyTree.forEach((k, v) -> {
            graph.putIfAbsent(k, new ArrayList<>());
            v.forEach(e -> {
                graph.get(k).add(e);
                graph.putIfAbsent(e, new ArrayList<>());
                graph.get(e).add(k);
                v.stream().filter(sibling -> !sibling.equals(e))
                    .forEach(sibling -> graph.get(e).add(sibling));
            });
        });
        this.familyGraph = graph;
    }

    int degreeOfSeparation(String personA, String personB) {
        var visited = new HashSet<String>();
        var queue = new ArrayDeque<Pair>();
        queue.add(new Pair(personA, 0));
        visited.add(personA);
        while (!queue.isEmpty()) {
            var cur = queue.remove();
            var curName = cur.name();
            if (curName.equals(personB)) {
                return cur.degree();
            }
            var adjacent = familyGraph.get(curName);
            int nextDegree = cur.degree() + 1;
            for (var name: adjacent) {
                if (visited.contains(name)) {
                    continue;
                }
                if (name.equals(personB)) {
                    return nextDegree;
                }
                queue.add(new Pair(name, nextDegree));
                visited.add(name);
            }
        }
        return -1;
    }

    record Pair(String name, int degree) {}
}
