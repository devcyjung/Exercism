import java.util.List;

public class Satellite {
    public Node nodeFromTraversals(List<Character> preorderInput, List<Character> inorderInput) {
        if (preorderInput == null || preorderInput.isEmpty()) {
            if (inorderInput == null || inorderInput.isEmpty()) {
                return null;
            }
            throw new IllegalArgumentException("traversals must have the same length");
        }
        if (inorderInput.size() != preorderInput.size()) {
            throw new IllegalArgumentException("traversals must have the same length");
        }
        var root = preorderInput.getFirst();
        var inorderRootIndex = inorderInput.indexOf(root);
        if (inorderRootIndex == -1) {
            throw new IllegalArgumentException("traversals must have the same elements");
        }
        var leftInorder = inorderInput.subList(0, inorderRootIndex);
        var rightInorder = inorderInput.subList(inorderRootIndex + 1, inorderInput.size());
        var leftLength = leftInorder.size();
        var leftPreorder = preorderInput.subList(1, leftLength + 1);
        var rightPreorder = preorderInput.subList(leftLength + 1, preorderInput.size());
        if (rightInorder.contains(root) || leftPreorder.contains(root) || rightPreorder.contains(root)) {
            throw new IllegalArgumentException("traversals must contain unique items");
        }
        return new Node(root, nodeFromTraversals(leftPreorder, leftInorder), nodeFromTraversals(rightPreorder, rightInorder));
    }

    public Tree treeFromTraversals(List<Character> preorderInput, List<Character> inorderInput) {
        return new Tree(nodeFromTraversals(preorderInput, inorderInput));
    }
}
