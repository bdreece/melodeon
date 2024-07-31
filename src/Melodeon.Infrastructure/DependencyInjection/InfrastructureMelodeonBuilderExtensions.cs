using Microsoft.Extensions.Configuration;
using Serilog;


using Microsoft.Extensions.Http;
using Microsoft.Extensions.DependencyInjection;
using Melodeon.Domain;
using Melodeon.Data;
using Melodeon.Infrastructure;
using Melodeon.Infrastructure.Data;
using Melodeon.Infrastructure.Security;
using Melodeon.Infrastructure.Spotify;
using Melodeon.Spotify;
using Microsoft.Extensions.Options;
using Serilog.Events;

namespace Microsoft.AspNetCore.Builder;

public static class InfrastructureMelodeonBuilderExtensions
{
    public static IMelodeonBuilder ConfigureInfrastructureDefaults(this IMelodeonBuilder builder)
    {
        return builder
            .ConfigureCookies()
            .AddSqliteRepository()
            .AddSpotifyClients();
    }

    public static IMelodeonBuilder AddSerilog(
        this IMelodeonBuilder builder, 
        Action<LoggerConfiguration>? configureAction = default)
    {
        LoggingOptions options;
        using (var sp = builder.Services.BuildServiceProvider())
        {
            options = sp.GetRequiredService<IOptions<LoggingOptions>>().Value;
        }

        var configuration = new LoggerConfiguration()
            .MinimumLevel.Override("Microsoft.EntityFrameworkCore", LogEventLevel.Debug)
            .ReadFrom.Configuration(builder.Configuration)
            .Enrich.FromLogContext()
            .WriteTo.Console()
            .WriteTo.File(path: options.FilePath);

        configureAction?.Invoke(configuration);

        Log.Logger = configuration.CreateLogger();

        builder.Logging.AddSerilog();

        return builder;
    }

    public static IMelodeonBuilder AddSqliteRepository(this IMelodeonBuilder builder)
    {
        builder.Services
            .AddDbContext<ApplicationDbContext>()
            .AddScoped<IRepository<Room>>(sp => sp.GetRequiredService<ApplicationDbContext>())
            .AddScoped<IReadOnlyRepository<Room>>(sp => sp.GetRequiredService<IRepository<Room>>());

        using (var sp = builder.Services.BuildServiceProvider())
        {
            var db = sp.GetRequiredService<ApplicationDbContext>();
            db.Database.EnsureCreated();
        }

        return builder;
    }

    public static IMelodeonBuilder AddSpotifyClients(this IMelodeonBuilder builder)
    {
        static HttpClient resolveHttpClient(IServiceProvider sp, object? key) => sp
            .GetRequiredService<IHttpClientFactory>()
            .CreateClient((string)key!);

        builder.Services
            .Configure<SpotifyOptions>(options => builder.Configuration
                .GetRequiredSection(SpotifyOptions.Spotify)
                .Bind(options));

        builder.Services
            .AddHttpClient(nameof(SpotifyServerHttpClient))
            .AddTypedClient<SpotifyServerHttpClient>();

        builder.Services
            .AddHttpClient(nameof(SpotifyUserHttpClient))
            .AddTypedClient<SpotifyUserHttpClient>();

        builder.Services
            .AddKeyedSingleton(nameof(SpotifyServerHttpClient), resolveHttpClient)
            .AddSingleton<ISpotifyServerClient>(static sp =>
            {
                var httpClient = sp.GetRequiredKeyedService<HttpClient>(nameof(SpotifyServerHttpClient));

                return sp
                    .GetRequiredService<ITypedHttpClientFactory<SpotifyServerHttpClient>>()
                    .CreateClient(httpClient);
            });

        builder.Services
            .AddKeyedScoped(nameof(SpotifyUserHttpClient), resolveHttpClient)
            .AddScoped<ISpotifyUserClient>(static sp =>
            {
                var httpClient = sp.GetRequiredKeyedService<HttpClient>(nameof(SpotifyUserHttpClient));

                return sp
                    .GetRequiredService<ITypedHttpClientFactory<SpotifyUserHttpClient>>()
                    .CreateClient(httpClient);
            });

        return builder;
    }

    public static IMelodeonBuilder ConfigureCookies(this IMelodeonBuilder builder)
    {
        builder.Services
            .Configure<CookieOptions>(CookieOptions.Host, options => builder.Configuration
                .GetRequiredSection(CookieOptions.Cookie)
                .GetRequiredSection(CookieOptions.Host)
                .Bind(options))
            .Configure<CookieOptions>(CookieOptions.Guest, options => builder.Configuration
                .GetRequiredSection(CookieOptions.Cookie)
                .GetRequiredSection(CookieOptions.Guest)
                .Bind(options));

        return builder;
    }
}