public static class Diamond
{
    public static string Make(char target)
    {
        int order = target - 'A';
        if (order < 0 || order > 25)
            throw new Exception("target must be between A and Z");
        int side = 2 * order + 1;
        int left = order, right = order, top = 0, bottom = side - 1;
        char[] buf = new char[side];
        Array.Fill(buf, ' ');
        string[] diamond = new string[side];
        for (char ch = 'A'; ch <= target; ++ch)
        {
            buf[left] = buf[right] = ch;
            diamond[top++] = diamond[bottom--] = new String(buf);
            buf[left--] = buf[right++] = ' ';
        }
        return String.Join('\n', diamond);
    }
}