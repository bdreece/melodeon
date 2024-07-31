using Microsoft.Extensions.Options;

namespace Melodeon.Infrastructure;

public sealed class LoggingOptions : IOptions<LoggingOptions>
{
    public const string Logging = nameof(Logging);

    public required string FilePath { get; init; }

    LoggingOptions IOptions<LoggingOptions>.Value => this;
}