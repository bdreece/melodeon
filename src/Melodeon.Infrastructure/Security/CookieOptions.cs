using Microsoft.Extensions.Options;

namespace Melodeon.Infrastructure.Security;

public sealed class CookieOptions : IOptions<CookieOptions>
{
    public const string Cookie = nameof(Cookie);
    public const string Host = nameof(Host);
    public const string Guest = nameof(Guest);

    public const string Path = "/";
    public const bool HttpOnly = false;
    public static readonly TimeSpan Expiration = TimeSpan.FromHours(8);

    public required string Name { get; set; }
    public required string LoginPath { get; set; }
    public required string LogoutPath { get; set; }

    CookieOptions IOptions<CookieOptions>.Value => this;
}