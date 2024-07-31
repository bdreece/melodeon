namespace Melodeon.Data;

public interface IRepository<T> : IReadOnlyRepository<T>, IMutable<T>
{
    Task<int> SaveChangesAsync(CancellationToken ct);
}