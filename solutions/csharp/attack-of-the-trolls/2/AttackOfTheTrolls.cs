enum AccountType {
    Guest = Permission.Read,
    User = Permission.Read | Permission.Write,
    Moderator = Permission.All,
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
    public static Permission Default(AccountType accountType) =>
        Enum.IsDefined(accountType) ? (Permission) accountType : Permission.None;

    public static Permission Grant(Permission current, Permission grant) => current | grant;

    public static Permission Revoke(Permission current, Permission revoke) => current & ~revoke;

    public static bool Check(Permission current, Permission check) => (current & check ^ check) == Permission.None;
}