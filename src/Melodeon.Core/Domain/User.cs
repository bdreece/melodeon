namespace Melodeon.Domain;

public sealed class User
{
    public required string Id { get; set; }
    public required string Name { get; set; }
    public required string Role { get; set; }
}