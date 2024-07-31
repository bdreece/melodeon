using Melodeon.Domain.Abstractions;

namespace Melodeon.Domain;

public sealed class Room : Entity
{
    public string Code { get; init; } = Guid.NewGuid().ToString();

    public required string Name { get; set; }
    public required string HostName { get; set; }
    public required string HostAccessToken { get; set; }
    public ICollection<string> Guests { get; } = new List<string>();
}