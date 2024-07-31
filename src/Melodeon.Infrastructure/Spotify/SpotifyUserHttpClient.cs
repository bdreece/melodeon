using Melodeon.Spotify;

namespace Melodeon.Infrastructure.Spotify;

internal sealed class SpotifyUserHttpClient : ISpotifyUserClient
{
    private readonly HttpClient _httpClient;

    public SpotifyUserHttpClient(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }
}