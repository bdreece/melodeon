public sealed class DotNetWatchSettings : DotNetRunSettings
{
    public bool NoHotReload { get; set; }
}

public sealed class DotNetWatcher : DotNetTool<DotNetWatchSettings>
{
    private readonly ICakeEnvironment _environment;

    public DotNetWatcher(
        IFileSystem fileSystem,
        ICakeEnvironment environment,
        IProcessRunner processRunner,
        IToolLocator tools)
            : base(fileSystem, environment, processRunner, tools)
    {
        _environment = environment;
    }

    public void Watch(string project, ProcessArgumentBuilder arguments, DotNetWatchSettings settings)
    {
        ArgumentNullException.ThrowIfNull(settings);
        RunCommand(settings, GetArguments(project, arguments, settings));
    }

    private ProcessArgumentBuilder GetArguments(string project, ProcessArgumentBuilder arguments, DotNetWatchSettings settings)
    {
        var builder = CreateArgumentBuilder(settings);

        builder.Append("watch");

        if (settings.NoHotReload)
        {
            builder.Append("--no-hot-reload");
        }
        
        builder.Append("run");

        // Specific path?
        if (project != null)
        {
            builder.Append("--project");
            builder.AppendQuoted(project);
        }

        // Framework
        if (!string.IsNullOrEmpty(settings.Framework))
        {
            builder.Append("--framework");
            builder.Append(settings.Framework);
        }

        // Configuration
        if (!string.IsNullOrEmpty(settings.Configuration))
        {
            builder.Append("--configuration");
            builder.Append(settings.Configuration);
        }

        // No Restore
        if (settings.NoRestore)
        {
            builder.Append("--no-restore");
        }

        // No Build
        if (settings.NoBuild)
        {
            builder.Append("--no-build");
        }

        // Runtime
        if (!string.IsNullOrEmpty(settings.Runtime))
        {
            builder.Append("--runtime");
            builder.Append(settings.Runtime);
        }

        // Sources
        if (settings.Sources != null)
        {
            foreach (var source in settings.Sources)
            {
                builder.Append("--source");
                builder.AppendQuoted(source);
            }
        }

        // Roll Forward Policy
        if (!(settings.RollForward is null))
        {
            builder.Append("--roll-forward");
            builder.Append(settings.RollForward.Value.ToString("F"));
        }

        // MSBuild Settings
        if (settings.MSBuildSettings != null)
        {
            builder.AppendMSBuildSettings(settings.MSBuildSettings, _environment);
        }

        // Arguments
        if (!arguments.IsNullOrEmpty())
        {
            builder.Append("--");
            arguments.CopyTo(builder);
        }

        return builder;
    }
} 

public static void DotNetWatch(ICakeContext context, DotNetWatchSettings settings)
{
    DotNetWatch(context, null, settings);
}

public static void DotNetWatch(
    ICakeContext context,
    string project,
    DotNetWatchSettings settings)
{
    DotNetWatch(context, project, null, settings);
}

public static void DotNetWatch(
    ICakeContext context,
    string project,
    ProcessArgumentBuilder arguments,
    DotNetWatchSettings settings)
{
    var watcher = new DotNetWatcher(
        context.FileSystem,
        context.Environment,
        context.ProcessRunner,
        context.Tools);

    watcher.Watch(project, arguments, settings);
}
