using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Areas.Host.Controllers;

[Authorize(Roles = nameof(Host))]
[Area(nameof(Host))]
public sealed class HomeController : Controller
{
    public IActionResult Index()
    {
        return View();
    }

    public async Task<IActionResult> Search([FromQuery(Name = "q")] string query)
    {
        return await Task.FromResult(Ok());
    }
}