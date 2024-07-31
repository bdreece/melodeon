using Microsoft.AspNetCore.Authentication;

namespace Microsoft.AspNetCore.Builder;

public static class WebMelodeonBuilderExtensions
{
    public static IMelodeonBuilder ConfigureWebDefaults(this IMelodeonBuilder builder)
    {
        return builder
            .AddAuthentication()
            .AddAuthorization()
            .AddSession()
            .AddControllers();
    }

    public static IMelodeonBuilder AddAuthentication(
        this IMelodeonBuilder builder, Action<AuthenticationOptions>? configureAction = default)
    {
        builder.Services
            .AddAuthentication(options =>
            {
                options.DefaultScheme = MelodeonAuthenticationDefaults.PolicyScheme;
                options.RequireAuthenticatedSignIn = false;

                configureAction?.Invoke(options);
            })
            .AddMelodeonRoot()
            .AddMelodeonHost();

        return builder;
    }

    public static IMelodeonBuilder AddAuthorization(this IMelodeonBuilder builder)
    {
        builder.Services.AddAuthorization();

        return builder;
    }

    public static IMelodeonBuilder AddSession(this IMelodeonBuilder builder)
    {
        builder.Services
            .AddDistributedMemoryCache()
            .AddSession(x =>
        {
            x.IdleTimeout = TimeSpan.FromHours(6);
            x.Cookie.HttpOnly = true;
            x.Cookie.IsEssential = true;
        });

        return builder;
    }

    public static IMelodeonBuilder AddControllers(this IMelodeonBuilder builder)
    {
        builder.Services.AddRouting(static x =>
        {
            x.LowercaseUrls = true;
            x.LowercaseQueryStrings = true;
        });

        builder.Services.AddControllersWithViews();

        return builder;
    }
}