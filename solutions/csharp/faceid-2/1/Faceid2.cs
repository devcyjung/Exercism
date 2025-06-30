public sealed record class FacialFeatures(string EyeColor, decimal PhiltrumWidth);

public sealed record class Identity(string Email, FacialFeatures FacialFeatures);

public class Authenticator
{
    private static readonly Identity adminIdentity =
        new Identity("admin@exerc.ism", new FacialFeatures("green", 0.9m));

    private HashSet<Identity> identityRegistry = new();
    
    public static bool AreSameFace(FacialFeatures faceA, FacialFeatures faceB) => faceA == faceB;

    public bool IsAdmin(Identity identity) => identity == adminIdentity;

    public bool Register(Identity identity) => identityRegistry.Add(identity);

    public bool IsRegistered(Identity identity) => identityRegistry.Contains(identity);

    public static bool AreSameObject(Identity identityA, Identity identityB) =>
        ReferenceEquals(identityA, identityB);
}