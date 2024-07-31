using Melodeon.Domain.Abstractions;

using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace Microsoft.EntityFrameworkCore;

public static class MelodeonEntityTypeBuilderExtensions
{
    public static EntityTypeBuilder<T> ConfigureEntity<T>(this EntityTypeBuilder<T> builder)
        where T : Entity
    {
        builder
            .Property(x => x.Id)
            .HasColumnName("id");

        builder
            .Property(x => x.CreatedAt)
            .HasColumnName("created_at")
            .HasConversion<long>();

        builder
            .Property(x => x.UpdatedAt)
            .HasColumnName("updated_at")
            .HasConversion<long>();

        builder.HasKey(x => x.Id);

        return builder;
    }
}