static class Badge
{
    public static string Print(int? id, string name, string? department) => (id, department) switch
    {
        (int empId, string empDep and not null) => $"[{empId}] - {name} - {empDep.ToUpper()}",
        (null, string empDep and not null) => $"{name} - {empDep.ToUpper()}",
        (int ownerId, null) => $"[{ownerId}] - {name} - OWNER",
        (null, null) => $"{name} - OWNER",
    };
}