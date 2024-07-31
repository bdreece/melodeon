namespace Melodeon.Data;

public interface IMutable<in T>
{
    void Add(T entity);
    void AddRange(IEnumerable<T> entities);

    void Update(T entity);
    void UpdateRange(IEnumerable<T> entities);

    void Remove(T entity);
    void RemoveRange(IEnumerable<T> entities);
}