namespace Melodeon.Data;

public interface IReadOnlyRepository<out T> : IAsyncDisposable, IQueryable<T> { }