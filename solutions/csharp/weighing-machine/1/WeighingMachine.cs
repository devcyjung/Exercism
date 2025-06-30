using System.Diagnostics.CodeAnalysis;

class WeighingMachine
{
    [SetsRequiredMembers]
    public WeighingMachine(int precision)
    {
        Precision = precision;
    }

    public required int Precision { get; init; }

    private double weight;
    public double Weight
    {
        get => weight;
        set => weight = value < 0
            ? throw new ArgumentOutOfRangeException(nameof(value), $"{nameof(value)} cannot be less negative")
            : value;
    }

    public double TareAdjustment { get; set; } = 5;

    public string DisplayWeight { get => $"{(Weight - TareAdjustment).ToString($"F{Precision}")} kg"; }
}