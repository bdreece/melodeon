namespace Melodeon.Web.Areas.Guest.Models;

public sealed class LoginModel
{
    public required string Code { get; set; }

    public string DisplayName { get; set; } = string.Empty;
}