static class GameMaster
{
    public static string Describe(Character character) =>
        $"You're a level {character.Level} {character.Class} with {character.HitPoints} hit points.";

    public static string Describe(Destination destination) =>
        $"You've arrived at {destination.Name}, which has {destination.Inhabitants} inhabitants.";

    public static string Describe(TravelMethod travelMethod) =>
        $"You're traveling to your destination {travelMethod.Description()}.";

    public static string Describe(
        Character character, Destination destination, TravelMethod travelMethod = TravelMethod.Walking
    ) =>
        $"{Describe(character)} {Describe(travelMethod)} {Describe(destination)}";
}

record struct Character(string Class, int Level, int HitPoints);

record struct Destination(string Name, int Inhabitants);

enum TravelMethod { Walking, Horseback }

static class TravelMethodExtension
{
    public static string Description(this TravelMethod method) => method switch
    {
        TravelMethod.Walking => "by walking",
        TravelMethod.Horseback => "on horseback",
    };
}