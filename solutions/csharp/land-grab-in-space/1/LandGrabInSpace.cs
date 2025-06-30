public sealed record Coord(int x, int y);

public sealed record Plot(Coord c1, Coord c2, Coord c3, Coord c4);

public class ClaimsHandler
{
    private List<Plot> claimRegistry = new();
    
    public void StakeClaim(Plot plot) => claimRegistry.Add(plot);

    public bool IsClaimStaked(Plot plot) => claimRegistry.Contains(plot);

    public bool IsLastClaim(Plot plot) => claimRegistry.Count > 0 && claimRegistry[^1] == plot;

    public Plot GetClaimWithLongestSide() => claimRegistry
        .MaxBy(c => Math.Max(Math.Abs(c.c2.x - c.c1.x), Math.Abs(c.c3.y - c.c2.y)));
}