namespace System.Linq;

public static class QueryableExtensions
{
    public static ValueTask<T?> FirstOrDefaultAsync<T>(this IQueryable<T> source, CancellationToken ct = default)
    {
        return source.Take(1).ToAsyncEnumerable().FirstOrDefaultAsync(ct);
    }

    public static ValueTask<T[]> ToArrayAsync<T>(this IQueryable<T> source, CancellationToken ct = default)
    {
        return source.ToAsyncEnumerable().ToArrayAsync();
    }
}