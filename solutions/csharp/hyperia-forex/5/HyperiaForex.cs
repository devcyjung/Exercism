public readonly struct CurrencyAmount
{
    private readonly decimal amount;
    private readonly string currency;

    public CurrencyAmount(decimal amount, string currency)
    {
        this.amount = amount;
        this.currency = currency;
    }

    public static bool operator ==(CurrencyAmount one, CurrencyAmount other) => one - other == 0;

    public static bool operator !=(CurrencyAmount one, CurrencyAmount other) => one - other != 0;

    public static bool operator >(CurrencyAmount one, CurrencyAmount other) => one - other > 0;

    public static bool operator <(CurrencyAmount one, CurrencyAmount other) => one - other < 0;

    public static CurrencyAmount operator +(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount + other.amount, one.currency)
            : throw new ArgumentException();

    public static CurrencyAmount operator -(CurrencyAmount one) => new(-one.amount, one.currency);

    public static CurrencyAmount operator -(CurrencyAmount one, CurrencyAmount other) => one + (-other);

    public static CurrencyAmount operator *(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount * other.amount, one.currency)
            : throw new ArgumentException();

    public static CurrencyAmount operator ~(CurrencyAmount one) => new(1 / one.amount, one.currency);

    public static CurrencyAmount operator /(CurrencyAmount one, CurrencyAmount other) => one * (~other);

    public static explicit operator double(CurrencyAmount one) => (double) one.amount;

    public static implicit operator decimal(CurrencyAmount one) => one.amount;
}