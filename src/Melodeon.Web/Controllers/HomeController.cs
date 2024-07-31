using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Controllers;

public sealed class HomeController : Controller
{
    public IActionResult Index() => View();
}