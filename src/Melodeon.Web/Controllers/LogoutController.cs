using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Controllers;

public sealed class LogoutController : Controller
{
    public async Task<IActionResult> IndexAsync()
    {
        await HttpContext.SignOutAsync();
        HttpContext.User = new();

        return View();
    }
}