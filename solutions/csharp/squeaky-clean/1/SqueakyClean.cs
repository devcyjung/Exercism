using System.Text;

public static class Identifier
{
    public static string Clean(string identifier)
    {
        bool capitalize = false;
        var b = new StringBuilder(identifier.Length);
        foreach (char c in identifier)
        {
            if (Char.IsWhiteSpace(c))
            {
                b.Append('_');
            }
            else if (Char.IsControl(c))
            {
                b.Append("CTRL");
            }
            else if (c == '-')
            {
                capitalize = true;
            }
            else if (!Char.IsLetter(c))
            {
                continue;
            }
            else if (Char.IsBetween(c, 'α', 'ω'))
            {
                continue;
            }
            else if (capitalize)
            {
                b.Append(Char.ToUpper(c));
                capitalize = false;
            }
            else
            {
                b.Append(c);
            }
        }
        return b.ToString();
    }
}