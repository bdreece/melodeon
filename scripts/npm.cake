#addin "nuget:?package=Cake.Npm&version=4.0.0"

using Cake.Npm.BumpVersion;
using Cake.Npm.Ci;
using Cake.Npm.RunScript;

public sealed class NpmWorkspaceBumpVersionSettings : NpmBumpVersionSettings
{
    public bool AllWorkspaces { get; init; }
    public bool IncludeWorkspaceRoot { get; init; }
    public IList<string> Workspaces { get; init; } = new List<string>();

    protected override void EvaluateCore(ProcessArgumentBuilder args)
    {
        base.EvaluateCore(args);

        if (IncludeWorkspaceRoot)
        {
            args.Append("--include-workspace-root");
        }

        if (AllWorkspaces)
        {
            args.Append("--workspaces");
        }
        else if (Workspaces.Any())
        {
            foreach (var ws in Workspaces)
            {
                args.AppendSwitch("--workspace", ws);
            }
        }
    }
}

public sealed class NpmWorkspaceCiSettings : NpmCiSettings
{
    public bool AllWorkspaces { get; init; }
    public bool IncludeWorkspaceRoot { get; init; }
    public IList<string> Workspaces { get; init; } = new List<string>();

    protected override void EvaluateCore(ProcessArgumentBuilder args)
    {
        base.EvaluateCore(args);

        if (IncludeWorkspaceRoot)
        {
            args.Append("--include-workspace-root");
        }

        if (AllWorkspaces)
        {
            args.Append("--workspaces");
        }
        else if (Workspaces.Any())
        {
            foreach (var ws in Workspaces)
            {
                args.AppendSwitch("--workspace", ws);
            }
        }
    }
}

public sealed class NpmWorkspaceRunScriptSettings : NpmRunScriptSettings
{
    public bool IfPresent { get; init; }
    public bool AllWorkspaces { get; init; }
    public bool IncludeWorkspaceRoot { get; init; }
    public IList<string> Workspaces { get; init; } = new List<string>();

    protected override void EvaluateCore(ProcessArgumentBuilder args)
    {
        if (string.IsNullOrEmpty(ScriptName))
        {
            throw new ArgumentNullException(nameof(ScriptName), "Must provide script name.");
        }

        if (IncludeWorkspaceRoot)
        {
            args.Append("--include-workspace-root");
        }

        if (AllWorkspaces)
        {
            args.Append("--workspaces");
        }
        else if (Workspaces.Any())
        {
            foreach (var ws in Workspaces)
            {
                args.AppendSwitch("--workspace", ws);
            }
        }

        if (IfPresent)
        {
            args.Append("--if-present");
        }

        args.AppendQuoted(ScriptName);

        if (Arguments.Any())
        {
            args.Append("--");
            foreach (var arg in Arguments)
            {
                args.Append(arg);
            }
        }
    }
}
