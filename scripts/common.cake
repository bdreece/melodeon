using Spectre.Console;

#nullable enable

public sealed record BuildState(
    string Configuration,
    string Framework,
    string Environment)
{
    public const string Release = nameof(Release);
    public const string Debug = nameof(Debug);

    public const string Net8 = "net8.0";

    public const string Development = nameof(Development);
    public const string Production = nameof(Production);

    public bool IsRelease => Configuration == Release;
    public bool IsDebug => Configuration == Debug;

    public bool IsNet8 => Framework == Net8;

    public bool IsDevelopment => Environment == Development;
    public bool IsProduction => Environment == Production;

    public static BuildState Setup(ICakeContext context, IDictionary<string, string> args)
    {
        AnsiConsole.MarkupLine("[green]Loading build state...[/]");
        
        var table = new Table();
        table.HideHeaders();
        table.AddColumn(string.Empty);
        table.AddColumn(string.Empty);

        string? configuration;
        if (!args.TryGetValue(nameof(configuration), out configuration) &&
            !args.TryGetValue("c", out configuration))
        {
            configuration = Debug;
        }

        table.AddRow("[yellow]Configuration:[/]", configuration);

        string? framework;
        if (!args.TryGetValue(nameof(framework), out framework) &&
            !args.TryGetValue("f", out framework))
        {
            framework = Net8;
        }

        table.AddRow("[yellow]Framework:[/]", framework);

        string? environment;
        if (!args.TryGetValue(nameof(environment), out environment) &&
            !args.TryGetValue("E", out environment))
        {
            environment = Development;
        }

        table.AddRow("[yellow]Environment:[/]", environment);

        AnsiConsole.Write(table);

        return new(configuration, framework, environment);
    }
}

#nullable restore
