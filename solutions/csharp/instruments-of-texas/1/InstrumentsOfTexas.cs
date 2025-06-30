public class CalculationException : Exception
{
    public CalculationException(int operand1, int operand2, string message, Exception inner)
        : base(message, inner)
    {
        Operand1 = operand1;
        Operand2 = operand2;
    }

    public int Operand1 { get; }
    public int Operand2 { get; }
}

public class CalculatorTestHarness
{
    private Calculator calculator;

    public CalculatorTestHarness(Calculator calculator)
    {
        this.calculator = calculator;
    }

    public string TestMultiplication(int x, int y)
    {
        try
        {
            Multiply(x, y);
            return "Multiply succeeded";
        }
        catch (CalculationException ex) when (ex.Operand1 < 0 && ex.Operand2 < 0)
        {
            return $"Multiply failed for negative operands. {ex.InnerException.Message}";
        }
        catch (CalculationException ex) when (ex.Operand1 >= 0 || ex.Operand2 >= 0)
        {
            return $"Multiply failed for mixed or positive operands. {ex.InnerException.Message}";
        }
    }

    public int Multiply(int x, int y)
    {
        try
        {
            return calculator.Multiply(x, y);
        }
        catch (Exception ex)
        {
            throw new CalculationException(x, y, "", ex);
        }
    }
}

public class Calculator
{
    public int Multiply(int x, int y) => checked(x * y);
}