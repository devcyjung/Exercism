public static class SimpleCalculator
{
    public static string Calculate(int operand1, int operand2, string? operation) => operation switch
    {
        null => throw new ArgumentNullException("Operation cannot be null"),
        "" => throw new ArgumentException("Operation cannot be empty"),
        "+" => $"{operand1} + {operand2} = {operand1 + operand2}",
        "*" => $"{operand1} * {operand2} = {operand1 * operand2}",
        "/" => Divide(operand1, operand2),
        _ => throw new ArgumentOutOfRangeException($"Operation {operation} unsupported"),
    };

    private static string Divide(int operand1, int operand2)
    {
        try
        {
            return $"{operand1} / {operand2} = {operand1 / operand2}";
        }
        catch (DivideByZeroException)
        {
            return "Division by zero is not allowed.";
        }
    }
}
