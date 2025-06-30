using System.Collections.Immutable;

public class Authenticator
{
    public Authenticator(Identity admin) => this.admin = admin;

    private readonly Identity admin;

    private Dictionary<string, Identity> developers = new()
        {
            ["Bertrand"] = new("bert@ex.ism", EyeColorKind.Blue),
            ["Anders"] = new("anders@ex.ism", EyeColorKind.Brown)
        };

    public Identity Admin => new(admin.Email, admin.EyeColor);

    public IDictionary<string, Identity> GetDevelopers() => developers.ToImmutableDictionary();
}

public enum EyeColorKind
{
    Blue, Green, Brown, Hazel, Grey,
}

public readonly record struct EyeColor(EyeColorKind Color)
{   
    public static implicit operator EyeColor(string color) =>
        new(Enum.Parse<EyeColorKind>(color, ignoreCase: true));

    public static implicit operator string(EyeColor color) =>
        color.Color.ToString().ToLower();

    public static implicit operator EyeColor(EyeColorKind color) => new(color);

    public static implicit operator EyeColorKind(EyeColor color) => color.Color;
}

public record struct Identity(string Email, EyeColor EyeColor);