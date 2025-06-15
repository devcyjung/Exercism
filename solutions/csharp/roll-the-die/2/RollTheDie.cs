public class Player
{
    private Random rand = Random.Shared;
    
    public int RollDie() => rand.Next(1, 19);

    public double GenerateSpellStrength() => rand.NextDouble() * 100;
}