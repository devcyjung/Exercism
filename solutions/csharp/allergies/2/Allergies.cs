public enum Allergen
{
    Eggs,
    Peanuts,
    Shellfish,
    Strawberries,
    Tomatoes,
    Chocolate,
    Pollen,
    Cats
}

public record class Allergies(int mask)
{
    private static readonly Allergen[] ALLERGENS = Enum.GetValues<Allergen>();
    
    public bool IsAllergicTo(Allergen allergen) => (mask & 1 << (int) allergen) > 0;

    public Allergen[] List() => ALLERGENS.Where(allergen => IsAllergicTo(allergen)).ToArray();
}