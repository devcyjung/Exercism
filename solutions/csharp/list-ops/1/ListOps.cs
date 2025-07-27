public static class ListOps
{
    public static int Length<T>(List<T> input) =>
        input.Count;

    public static List<T> Reverse<T>(List<T> input) =>
        input.AsEnumerable().Reverse().ToList();

    public static List<U> Map<T, U>(List<T> input, Func<T, U> map) =>
        input.Select(map).ToList();

    public static List<T> Filter<T>(List<T> input, Func<T, bool> predicate) =>
        input.Where(predicate).ToList();

    public static U Foldl<T, U>(List<T> input, U start, Func<U, T, U> func) =>
        input.Aggregate(start, func);

    public static U Foldr<T, U>(List<T> input, U start, Func<T, U, U> func) =>
        Reverse(input).Aggregate(start, (el, acc) => func(acc, el));

    public static List<T> Concat<T>(List<List<T>> input) =>
        input.SelectMany(e => e).ToList();

    public static List<T> Append<T>(List<T> left, List<T> right) =>
        left.Concat(right).ToList();
}