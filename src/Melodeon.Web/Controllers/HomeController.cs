using Melodeon.Domain;

using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Controllers;

public sealed class HomeController : Controller
{
    public IActionResult Index()
    {
        if (User.IsInRole(Role.Host))
        {
            return Redirect("/host");
        }
        else if (User.IsInRole(Role.Guest))
        {
            return Redirect("/guest");
        }
        else
        {
            return View();
        }
    }
}