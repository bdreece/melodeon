using System.Net.Http.Headers;
using System.Text.Json;

using Melodeon.Infrastructure.Spotify;

using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Authentication.OAuth;
using Microsoft.Extensions.Options;

namespace Microsoft.AspNetCore.Authentication;

public static class MelodeonAuthenticationBuilderExtensions
{
    private const string HostDefaultScheme = "Host Cookies";
    private const string HostChallengeScheme = "Host OAuth";

    public static AuthenticationBuilder AddMelodeonRoot(this AuthenticationBuilder builder) => builder
        .AddCookie(CookieAuthenticationDefaults.AuthenticationScheme)
        .AddPolicyScheme(
            MelodeonAuthenticationDefaults.PolicyScheme,
            MelodeonAuthenticationDefaults.PolicyScheme, static x =>
        {
            x.ForwardDefaultSelector = static context =>
            {
                if (context.Request.Path.StartsWithSegments("/host"))
                    return MelodeonAuthenticationDefaults.HostScheme;
                else
                    return CookieAuthenticationDefaults.AuthenticationScheme;
            };
        });

    public static AuthenticationBuilder AddMelodeonHost(this AuthenticationBuilder builder) => builder
        .AddPolicyScheme(
            MelodeonAuthenticationDefaults.HostScheme,
            MelodeonAuthenticationDefaults.HostScheme, static x =>
        {
            x.ForwardDefault = CookieAuthenticationDefaults.AuthenticationScheme;
            x.ForwardChallenge = HostChallengeScheme;
        })
        .AddOAuth(HostChallengeScheme, x =>
        {
            SpotifyOptions options;
            using (var sp = builder.Services.BuildServiceProvider())
            {
                options = sp.GetRequiredService<IOptions<SpotifyOptions>>().Value;
            }

            x.ClientId = options.ClientId;
            x.ClientSecret = options.ClientSecret;
            x.CallbackPath = options.CallbackPath;
            x.AuthorizationEndpoint = new Uri(SpotifyOptions.AccountsUri, "/authorize").ToString();
            x.TokenEndpoint = new Uri(SpotifyOptions.AccountsUri, "/api/token").ToString();
            x.UserInformationEndpoint = new Uri(SpotifyOptions.ApiUri, "/v1/me").ToString();

            foreach (var scope in SpotifyOptions.Scopes)
            {
                x.Scope.Add(scope);
            }

            foreach (var (claimType, key) in SpotifyOptions.JsonKeys)
            {
                x.ClaimActions.MapJsonKey(claimType, key);
            }

            foreach (var (claimType, func) in SpotifyOptions.CustomJson)
            {
                x.ClaimActions.MapCustomJson(claimType, func);
            }

            x.Events = new OAuthEvents
            {
                OnCreatingTicket = static async context =>
                {
                    using var request = new HttpRequestMessage(
                        HttpMethod.Get, context.Options.UserInformationEndpoint);

                    request.Headers.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
                    request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", context.AccessToken);

                    using var response = await context.Backchannel.SendAsync(
                        request, HttpCompletionOption.ResponseHeadersRead);

                    response.EnsureSuccessStatusCode();

                    using var stream = await response.Content.ReadAsStreamAsync();
                    using (var document = await JsonDocument.ParseAsync(stream))
                    {
                        context.RunClaimActions(document.RootElement);
                    }
                },
            };
        });
}

public static class MelodeonAuthenticationDefaults
{
    public const string PolicyScheme = "Melodeon Root Policy";
    public const string HostScheme = "Melodeon Host Policy";
    public const string GuestScheme = "Melodeon Guest Cookies";
}