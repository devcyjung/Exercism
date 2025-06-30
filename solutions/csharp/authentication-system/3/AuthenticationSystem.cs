using System.Collections.Immutable;

public class Authenticator
{
    private static class EyeColor
    {
        public const string Blue = "blue";
        public const string Green = "green";
        public const string Brown = "brown";
        public const string Hazel = "hazel";
        public const string Grey = "grey";
    }

    public Authenticator(Identity admin) => this.admin = admin;

    private readonly Identity admin;

    private readonly Dictionary<string, Identity> developers = new()
        {
            ["Bertrand"] = new("bert@ex.ism", "blue"),
            ["Anders"] = new("anders@ex.ism", "brown"),
        };

    public Identity Admin => new(admin.Email, admin.EyeColor);

    public IDictionary<string, Identity> GetDevelopers() => developers.ToImmutableDictionary();
}

public record struct Identity(string Email, string EyeColor);