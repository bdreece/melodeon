using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Diagnostics.Metrics;
using Microsoft.Extensions.Logging;

namespace Microsoft.AspNetCore.Builder;

public interface IMelodeonBuilder
{
    IServiceCollection Services { get; }
    IConfiguration Configuration { get; }
    ILoggingBuilder Logging { get; }
    IMetricsBuilder Metrics { get; }
}