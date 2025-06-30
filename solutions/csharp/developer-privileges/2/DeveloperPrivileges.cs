public class Authenticator
{
    private static readonly Identity admin = new("admin@ex.ism", new("green", 0.9m), ["Chanakya", "Mumbai", "India"]); 
    
    public Identity Admin => admin;

    private static readonly Dictionary<string, Identity> devs = new()
    {
        ["Bertrand"] = new("bert@ex.ism", new("blue", 0.8m), ["Bertrand", "Paris", "France"]),
        ["Anders"] = new("anders@ex.ism", new("brown", 0.85m), ["Anders", "Redmond", "USA"]),
    };
    
    public IDictionary<string, Identity> Developers => devs;
}

public record class FacialFeatures(string EyeColor, decimal PhiltrumWidth);

public record class Identity(string Email, FacialFeatures FacialFeatures, IList<string> NameAndAddress);