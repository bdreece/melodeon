using System.Security.Claims;

using Melodeon.Domain;

namespace Melodeon.Web.Middleware;

public static class UserSessionMiddleware
{
    public static async Task InvokeAsync(HttpContext context, RequestDelegate next)
    {
        var id = context.User.FindFirstValue(ClaimTypes.NameIdentifier);
        var name = context.User.FindFirstValue(ClaimTypes.Name);
        var role = context.User.FindFirstValue(ClaimTypes.Role);

        if (id is not null)
        {
            context.Session.SetString(nameof(User.Id), id);
        }

        if (name is not null)
        {
            context.Session.SetString(nameof(User.Name), name);
        }

        if (role is not null)
        {
            context.Session.SetString(nameof(User.Role), role);
        }

        await next(context);
    }
}