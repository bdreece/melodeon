using Melodeon.Web.Middleware;

namespace Microsoft.AspNetCore.Builder;

public static class MelodeonWebApplicationExtensions
{
    public static WebApplication UseMelodeonDefaults(this WebApplication app)
    {
        app.UseStaticFiles();
        app.UseSession();
        app.UseAuthentication();
        app.UseAuthorization();

        app.Use(UserSessionMiddleware.InvokeAsync);

        app.MapAreaControllerRoute(
            name: "guest",
            areaName: nameof(Melodeon.Web.Areas.Guest),
            pattern: "guest/{controller=Home}/{action=Index}/{id?}");

        app.MapAreaControllerRoute(
            name: "host",
            areaName: nameof(Melodeon.Web.Areas.Host),
            pattern: "host/{controller=Home}/{action=Index}/{id?}");

        app.MapDefaultControllerRoute();

        return app;
    }
}