public static class TelemetryBuffer
{
    private const byte P_LONG = 256 - sizeof(long);
    private const byte P_UINT = sizeof(uint);
    private const byte P_INT = 256 - sizeof(int);
    private const byte P_SHORT = 256 - sizeof(short);
    private const byte P_USHORT = sizeof(ushort);
    private const int BUF_LEN = 9;
    
    public static byte[] ToBuffer(long reading) => reading switch
    {
        > UInt32.MaxValue or < Int32.MinValue => BufPrefix(P_LONG, BitConverter.GetBytes(reading)),
        > Int32.MaxValue => BufPrefix(P_UINT, BitConverter.GetBytes((uint) reading)),
        > UInt16.MaxValue or < Int16.MinValue => BufPrefix(P_INT, BitConverter.GetBytes((int) reading)),
        < 0 => BufPrefix(P_SHORT, BitConverter.GetBytes((short) reading)),
        _ => BufPrefix(P_USHORT, BitConverter.GetBytes((ushort) reading)),
    };

    private static byte[] BufPrefix(byte prefix, byte[] src)
    {
        byte[] result = new byte[BUF_LEN];
        result[0] = prefix;
        Buffer.BlockCopy(src, 0, result, 1, src.Length);
        return result;
    }
        
    public static long FromBuffer(byte[] buffer)
    {
        if (buffer.Length != BUF_LEN)
        {
            return 0;
        }
        return buffer[0] switch
        {
            P_LONG => BitConverter.ToInt64(buffer, 1),
            P_UINT => (long) BitConverter.ToUInt32(buffer, 1),
            P_INT => (long) BitConverter.ToInt32(buffer, 1),
            P_SHORT => (long) BitConverter.ToInt16(buffer, 1),
            P_USHORT => (long) BitConverter.ToUInt16(buffer, 1),
            _ => 0,
        };
    }
}