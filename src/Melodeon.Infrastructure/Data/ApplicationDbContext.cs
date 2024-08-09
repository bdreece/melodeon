using System.Collections;
using System.Linq.Expressions;

using Melodeon.Data;
using Melodeon.Domain;

using Microsoft.EntityFrameworkCore;

namespace Melodeon.Infrastructure.Data;

internal sealed class ApplicationDbContext : DbContext, IRepository<Room>
{
    public required DbSet<Room> Rooms { get; set; }

    public Type ElementType => Rooms.AsQueryable().ElementType;
    public Expression Expression => Rooms.AsQueryable().Expression;
    public IQueryProvider Provider => Rooms.AsQueryable().Provider;

    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
        : base(options) { }

    protected override void OnModelCreating(ModelBuilder builder)
    {
        var room = builder.Entity<Room>();
        room.ConfigureEntity();

        room.Property(x => x.Name)
            .HasColumnName("name");

        room.Property(x => x.HostName)
            .HasColumnName("host_name");

        room.Property(x => x.HostAccessToken)
            .HasColumnName("host_access_token");

        room.ToTable("rooms");
    }

    protected override void OnConfiguring(DbContextOptionsBuilder builder)
    {
        builder.UseSqlite();
    }

    public IEnumerator<Room> GetEnumerator() => Rooms.AsQueryable().GetEnumerator();

    IEnumerator IEnumerable.GetEnumerator() => Rooms.AsQueryable().GetEnumerator();

    public void Add(Room entity) => base.Add(entity);
    public void AddRange(IEnumerable<Room> entities) => base.AddRange(entities);

    public void Update(Room entity) => base.Update(entity);
    public void UpdateRange(IEnumerable<Room> entities) => base.UpdateRange(entities);

    public void Remove(Room entity) => base.Remove(entity);
    public void RemoveRange(IEnumerable<Room> entities) => base.RemoveRange(entities);
}