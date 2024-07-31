using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Areas.Guest.Controllers;

[Area(nameof(Guest))]
[Authorize(Roles = nameof(Guest))]
public sealed class HomeController : Controller
{
    public IActionResult Index() => View();
}