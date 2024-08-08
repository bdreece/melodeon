using System.Text.Json.Serialization;

using Melodeon.Domain;

using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Areas.Host.Controllers;

[Area(nameof(Host))]
public sealed class AuthController : Controller
{
    [Authorize(Roles = Role.Host)]
    public async Task<IActionResult> Token()
    {
        var accessToken = await HttpContext.GetTokenAsync("access_token") ??
            throw new ArgumentNullException("access_token");

        var refreshToken = await HttpContext.GetTokenAsync("refresh_token") ??
            throw new ArgumentNullException("refresh_token");

        var expiresAt = await HttpContext.GetTokenAsync("expires_at") ??
            throw new ArgumentNullException("expires_at");

        var expiresIn = DateTime.UtcNow - DateTime.Parse(expiresAt);

        var response = new TokenResponse(
            AccessToken: accessToken,
            RefreshToken: refreshToken,
            ExpiresIn: (int)expiresIn.TotalSeconds);

        return Json(response);
    }

    private record TokenResponse(
        [property: JsonPropertyName("access_token")] string AccessToken,
        [property: JsonPropertyName("refresh_token")] string RefreshToken,
        [property: JsonPropertyName("expires_in")] int ExpiresIn);
}