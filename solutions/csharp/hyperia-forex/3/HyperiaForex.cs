public struct CurrencyAmount
{
    private decimal amount;
    private string currency;

    public CurrencyAmount(decimal amount, string currency)
    {
        this.amount = amount;
        this.currency = currency;
    }

    public static bool operator ==(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? one.amount == other.amount : throw new ArgumentException();

    public static bool operator !=(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? one.amount != other.amount : throw new ArgumentException();

    public static bool operator >(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? one.amount > other.amount : throw new ArgumentException();

    public static bool operator <(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? one.amount < other.amount : throw new ArgumentException();

    public static CurrencyAmount operator +(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount + other.amount, one.currency)
            : throw new ArgumentException();

    public static CurrencyAmount operator -(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount - other.amount, one.currency)
            : throw new ArgumentException();

    public static CurrencyAmount operator *(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount * other.amount, one.currency)
            : throw new ArgumentException();

    public static CurrencyAmount operator /(CurrencyAmount one, CurrencyAmount other) =>
        one.currency == other.currency ? new(one.amount / other.amount, one.currency)
            : throw new ArgumentException();

    public static explicit operator double(CurrencyAmount one) => (double) one.amount;

    public static implicit operator decimal(CurrencyAmount one) => one.amount;
}