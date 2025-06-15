public static class TelemetryBuffer
{
    public static byte[] ToBuffer(long reading) => reading switch
    {
        > UInt32.MaxValue or < Int32.MinValue => BufPrefix(256 - 8, BitConverter.GetBytes(reading)),
        > Int32.MaxValue => BufPrefix(4, BitConverter.GetBytes((uint) reading)),
        > UInt16.MaxValue or < Int16.MinValue => BufPrefix(256 - 4, BitConverter.GetBytes((int) reading)),
        < 0 => BufPrefix(256 - 2, BitConverter.GetBytes((short) reading)),
        _ => BufPrefix(2, BitConverter.GetBytes((ushort) reading)),
    };

    private static byte[] BufPrefix(byte prefix, byte[] src)
    {
        byte[] result = new byte[9];
        result[0] = prefix;
        Buffer.BlockCopy(src, 0, result, 1, src.Length);
        return result;
    }
        
    public static long FromBuffer(byte[] buffer) => (buffer.Length, buffer[0]) switch
    {
        (9, 256 - 8) => BitConverter.ToInt64(buffer, 1),
        (9, 4) => (long) BitConverter.ToUInt32(buffer, 1),
        (9, 256 - 4) => (long) BitConverter.ToInt32(buffer, 1),
        (9, 256 - 2) => (long) BitConverter.ToInt16(buffer, 1),
        (9, 2) => (long) BitConverter.ToUInt16(buffer, 1),
        _ => 0,
    };
}