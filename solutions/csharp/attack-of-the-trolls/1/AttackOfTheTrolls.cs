enum AccountType {
    Guest, User, Moderator,
}

[Flags]
enum Permission : byte {
    Read = 0b001,
    Write = 0b010,
    Delete = 0b100,
    All = Read | Write | Delete,
    None = 0b000,
}

static class Permissions
{
    public static Permission Default(AccountType accountType) => accountType switch
    {
        AccountType.Guest => Permission.Read,
        AccountType.User => Permission.Read | Permission.Write,
        AccountType.Moderator => Permission.All,
        _ => Permission.None,
    };

    public static Permission Grant(Permission current, Permission grant) => current | grant;

    public static Permission Revoke(Permission current, Permission revoke) => current & ~revoke;

    public static bool Check(Permission current, Permission check) => (current & check ^ check) == Permission.None;
}