public class Orm : IDisposable
{
    private Database database;

    public Orm(Database database) => this.database = database;

    public void Begin() => database.BeginTransaction();

    public void Write(string data) => ExecuteWithRecovery(() => database.Write(data));

    public void Commit() => ExecuteWithRecovery(() => database.EndTransaction());

    private void ExecuteWithRecovery(Action action, Action? recovery = null)
    {
        try
        {
            action();
        }
        catch (InvalidOperationException)
        {
            (recovery ?? Dispose)();
        }
    }

    public void Dispose() => database.Dispose();
}