using System.Reflection;

using FluentValidation;

using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Diagnostics.Metrics;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

namespace Microsoft.AspNetCore.Builder;

public static class MelodeonHostApplicationBuilderExtensions
{
    private static readonly Assembly s_assembly = typeof(MelodeonHostApplicationBuilderExtensions).Assembly;

    public static IMelodeonBuilder AddMelodeon(this IHostApplicationBuilder builder)
    {
        builder.Services
            .AddValidatorsFromAssembly(s_assembly, includeInternalTypes: true)
            .AddMediatR(x => x
                .RegisterServicesFromAssembly(s_assembly));

        return new MelodeonBuilder(
            Services: builder.Services,
            Configuration: builder.Configuration,
            Environment: builder.Environment,
            Logging: builder.Logging,
            Metrics: builder.Metrics);
    }

    private sealed record MelodeonBuilder(
        IServiceCollection Services,
        IConfiguration Configuration,
        IHostEnvironment Environment,
        ILoggingBuilder Logging,
        IMetricsBuilder Metrics) : IMelodeonBuilder;
}