#load "local:?path=common.cake"
#load "local:?path=npm.cake"
#load "local:?path=watch.cake"

#tool "dotnet:?package=GitVersion.Tool&version=5.12.0"

using Cake.Npm;
using Spectre.Console;

#nullable enable

public static class Tasks
{
    public const string Clean = nameof(Clean);
    public const string Restore = nameof(Restore);
    public const string Build = nameof(Build);
    public const string Tag = nameof(Tag);
    public const string Publish = nameof(Publish);
    public const string Test = nameof(Test);
    public const string Watch = nameof(Watch);

    public static readonly string[] All =
    {
        Clean, Restore, Build, Tag, Publish, Test, Watch
    };

    public static void CleanDotNet(ICakeContext context, BuildState state)
    {
        AnsiConsole.MarkupLine("[green]Cleaning .NET solution...[/]");

        context.DotNetClean(solution, new DotNetCleanSettings
        {
            Configuration = state.Configuration,
            Framework = state.Framework,
        });
    }

    public static void CleanNpm(ICakeContext context)
    {
        AnsiConsole.MarkupLine("[green]Cleaning NPM workspaces...[/]");

        context.NpmRunScript(new NpmWorkspaceRunScriptSettings
        {
            ScriptName = "clean",
            IfPresent = true,
            AllWorkspaces = true,
            IncludeWorkspaceRoot = true,
        });
    }
    
    public static void RestoreDotNet(ICakeContext context, BuildState state)
    {
        AnsiConsole.MarkupLine("[green]Restoring .NET solution...[/]");

        context.DotNetRestore(solution);
    }

    public static void RestoreNpm(ICakeContext context)
    {
        AnsiConsole.MarkupLine("[green]Restoring NPM workspaces...[/]");

        context.NpmCi(new NpmWorkspaceCiSettings
        {
            AllWorkspaces = true,
            IncludeWorkspaceRoot = true,
        });
    }

    public static void BuildDotNet(ICakeContext context, BuildState state)
    {
        AnsiConsole.MarkupLine("[green]Building .NET solution...[/]");

        context.DotNetBuild(solution, new DotNetBuildSettings
        {
            Configuration = state.Configuration,
            Framework = state.Framework,
            NoRestore = true,
        });
    }

    public static void TagDotNet(ICakeContext context, BuildState state)
    {
        AnsiConsole.MarkupLine($"[green]Tagging .NET solution...[/]");
        var version = context.GitVersion(
            settings: new GitVersionSettings
        {
            EnsureAssemblyInfo = true,
            UpdateAssemblyInfo = true,
        });

        var table = new Table();
        table.HideHeaders();
        table.AddColumn(string.Empty);
        table.AddColumn(string.Empty);

        table.AddRow("[yellow]Branch Name:[/]", version.BranchName);
        table.AddRow("[yellow]Commit Date:[/]", version.CommitDate);

        table.AddRow("[yellow]SHA:[/]", version.Sha);
        table.AddRow("[yellow]Short SHA:[/]", version.ShortSha);
        table.AddRow("[yellow]Uncommitted Changes:[/]",
            version.UncommittedChanges?.ToString() ?? string.Empty);

        table.AddRow("[yellow]Version:[/]", version.SemVer);
        table.AddRow("[yellow]Full Version:[/]", version.FullSemVer);
        table.AddRow("[yellow]Informational Version:[/]", version.InformationalVersion);

        AnsiConsole.Write(table);
    }

    public static void TagNpm(ICakeContext context)
    {
        AnsiConsole.MarkupLine($"[green]Tagging NPM workspaces...[/]");

        context.NpmBumpVersion(new NpmWorkspaceBumpVersionSettings
        {
            AllWorkspaces = true,
            IncludeWorkspaceRoot = true,
        });
    }

    public static void PrePublish() =>
        AnsiConsole.MarkupLine("[green]Publishing .NET solution...[/]");

    public static void PublishEach(BuildState state, string project, ICakeContext context)
    {
        var markup = new Markup($"Publishing .NET project: [yellow]{project}[/]");
        AnsiConsole.Write(new Padder(markup).PadLeft(2));

        context.DotNetPublish(project, new DotNetPublishSettings
        {
            Configuration = state.Configuration,
            Framework = state.Framework,
            NoRestore = true,
            NoBuild = true,
        });
    }

    public static void PreTestDotNet() =>
        AnsiConsole.MarkupLine("[green]Testing .NET solution...[/]");

    public static void TestEachDotNet(BuildState state, string project, ICakeContext context)
    {
        var markup = new Markup($"Testing .NET project: [yellow]{project}[/]");
        AnsiConsole.Write(new Padder(markup).PadLeft(2));

        context.DotNetTest(project, new DotNetTestSettings
        {
            Configuration = state.Configuration,
            Framework = state.Framework,
            NoRestore = true,
            NoBuild = true,
        });
    }

    public static void TestNpm(ICakeContext context)
    {
        AnsiConsole.MarkupLine("[green]Testing NPM workspaces...[/]");

        context.NpmRunScript(new NpmWorkspaceRunScriptSettings
        {
            ScriptName = "test",
            IfPresent = true,
            AllWorkspaces = true,
            IncludeWorkspaceRoot = true,
        });
    }

    public static void WatchDotNet(ICakeContext context, BuildState state)
    {
        var args = context.Arguments().ToDictionary(
            p => p.Key, p => p.Value.First());
        
        string? project;
        if (!args.TryGetValue(nameof(project), out project) &&
            !args.TryGetValue("p", out project))
        {
            project = apps[0];
        }

        AnsiConsole.MarkupLine($"[green]Watching .NET project: [yellow]{project}[/]...[/]");

        DotNetWatch(context, project, new DotNetWatchSettings
        {
            Configuration = state.Configuration,
            Framework = state.Framework,
            NoRestore = true,
        });
    }
}

public static class Criteria
{
    public static bool IsDebug(ICakeContext _, BuildState state) => state.IsDebug;
    public static bool IsRelease(ICakeContext _, BuildState state) => state.IsRelease;
}

#nullable restore

