using System.Security.Claims;
using System.Text.Json;

using Melodeon.Domain;

namespace Melodeon.Infrastructure.Spotify;

public sealed class SpotifyOptions
{
    public const string Spotify = nameof(Spotify);

    public static readonly Uri AccountsUri = new("https://accounts.spotify.com");
    public static readonly Uri ApiUri = new("https://api.spotify.com");
    public static readonly string[] Scopes =
    {
        "playlist-read-private",
        "playlist-read-collaborative",
        "streaming",
        "user-library-read",
        "user-modify-playback-state",
        "user-read-currently-playing",
        "user-read-playback-state",
        "user-read-email",
        "user-read-private",
    };

    public static readonly IReadOnlyDictionary<string, string> JsonKeys =
        new Dictionary<string, string>
        {
            [ClaimTypes.NameIdentifier] = "id",
            [ClaimTypes.Name] = "display_name",
            [ClaimTypes.Uri] = "href",
        };

    public static readonly IReadOnlyDictionary<string, Func<JsonElement, string?>> CustomJson =
        new Dictionary<string, Func<JsonElement, string?>>
        {
            [ClaimTypes.Role] = static _ => Role.Host,
            [ClaimTypes.UserData] = static el => el
                .GetProperty("images")
                .EnumerateArray()
                .FirstOrDefault()
                .GetProperty("url")
                .GetString(),
        };

    public required string ClientId { get; set; }
    public required string ClientSecret { get; set; }
    public required string CallbackPath { get; set; }
}