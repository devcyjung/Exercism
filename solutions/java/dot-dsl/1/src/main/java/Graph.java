import java.util.Collection;
import java.util.Collections;
import java.util.Map;
import java.util.ArrayList;

public class Graph {
    private final Map<String, String> attributes;
    private final Collection<Node> nodes;
    private final Collection<Edge> edges;
    
    public Graph() {
        this(Collections.emptyMap());
    }

    public Graph(Map<String, String> attributes) {
        this.attributes = attributes;
        this.nodes = new ArrayList<>();
        this.edges = new ArrayList<>();
    }

    public Collection<Node> getNodes() {
        return nodes;
    }

    public Collection<Edge> getEdges() {
        return edges;
    }

    public Graph node(String name) {
        nodes.add(new Node(name));
        return this;
    }

    public Graph node(String name, Map<String, String> attributes) {
        nodes.add(new Node(name, attributes));
        return this;
    }

    public Graph edge(String start, String end) {
        edges.add(new Edge(start, end));
        return this;
    }

    public Graph edge(String start, String end, Map<String, String> attributes) {
        edges.add(new Edge(start, end, attributes));
        return this;
    }

    public Map<String, String> getAttributes() {
        return attributes;
    }
}
