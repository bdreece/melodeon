using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Areas.Host.Controllers;

[Authorize(Roles = nameof(Host))]
[Area(nameof(Host))]
public sealed class HomeController : Controller
{
    public IActionResult Index() => View();
}