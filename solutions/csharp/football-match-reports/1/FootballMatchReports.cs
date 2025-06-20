public static class PlayAnalyzer
{
    public static string AnalyzeOnField(int shirtNum) => shirtNum switch
    {
        1 => "goalie",
        2 => "left back",
        3 or 4 => "center back",
        5 => "right back",
        >= 6 and <= 8 => "midfielder",
        9 => "left wing",
        10 => "striker",
        11 => "right wing",
        _ => "UNKNOWN"
    };

    public static string AnalyzeOffField(object report) => report switch
    {
        int supporterCount => $"There are {supporterCount} supporters at the match.",
        string announcement => announcement,
        Injury injuryIncident => $"Oh no! {injuryIncident.GetDescription()} Medics are on the field.",
        Incident incident => incident.GetDescription(),
        Manager { Name: var name, Club: null } => name,
        Manager { Name: var name, Club: var club } => $"{name} ({club})",
        _ => string.Empty
    };
}