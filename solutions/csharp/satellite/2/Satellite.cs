public record Tree(char Value, Tree? Left, Tree? Right);

public static class Satellite
{
    public static Tree? TreeFromTraversals(char[] preOrder, char[] inOrder) => (preOrder, inOrder) switch
    {
        ([], []) => null,
        ([var p], [var i]) when p == i => new Tree(p, null, null),
        ([var root, .. var pre], _) when preOrder.Length == inOrder.Length && inOrder.Count(e => e == root) == 1 && pre.All(e => e != root) => new Lazy<Tree?>(() =>
        {
            var leftInOrder = inOrder.TakeWhile(e => e != root).ToArray();
            var rightInOrder = inOrder.SkipWhile(e => e != root).Skip(1).ToArray();
            var leftLength = leftInOrder.Length;
            var leftPreOrder = pre.Take(leftLength).ToArray();
            var rightPreOrder = pre.Skip(leftLength).ToArray();
            return new Tree(root, TreeFromTraversals(leftPreOrder, leftInOrder), TreeFromTraversals(rightPreOrder, rightInOrder));
        }).Value,
        _ => throw new ArgumentException($"{preOrder} and {inOrder} don't form a binary tree.")
    };
}
