#load "local:?path=scripts/common.cake"
#load "local:?path=scripts/tasks.cake"

#nullable enable

using Spectre.Console;

const string solution = "./Melodeon.sln";

static readonly string[] apps =
{
    "./src/Melodeon.Web",
};

static readonly string[] tests = 
{
    "./src/Melodeon.UnitTests",
    "./src/Melodeon.IntegrationTests",
};

var args = Arguments().ToDictionary(
    p => p.Key,
    p => p.Value.First());

Setup(ctx => BuildState.Setup(ctx, args));

var clean = Task(Tasks.Clean)
    .Does<BuildState>(Tasks.CleanDotNet)
    .Does(Tasks.CleanNpm);

var restore = Task(Tasks.Restore)
    .Does<BuildState>(Tasks.RestoreDotNet)
    .Does(Tasks.RestoreNpm);

var build = Task(nameof(Tasks.Build))
    .IsDependentOn(clean)
    .IsDependentOn(restore)
    .Does<BuildState>(Tasks.BuildDotNet);

var tag = Task(Tasks.Tag)
    .WithCriteria<BuildState>(Criteria.IsRelease)
    .Does<BuildState>(Tasks.TagDotNet)
    .Does(Tasks.TagNpm);

var publish = Task(Tasks.Publish)
    .IsDependentOn(clean)
    .IsDependentOn(tag)
    .IsDependentOn(build)
    .WithCriteria<BuildState>(Criteria.IsRelease)
    .Does(Tasks.PrePublish)
    .DoesForEach<BuildState, string>(apps, Tasks.PublishEach);

var test = Task(Tasks.Test)
    .IsDependentOn(build)
    .Does(Tasks.PreTestDotNet)
    .DoesForEach<BuildState, string>(tests, Tasks.TestEachDotNet)
    .Does(Tasks.TestNpm);

var watch = Task(Tasks.Watch)
    .WithCriteria<BuildState>(Criteria.IsDebug)
    .Does<BuildState>(Tasks.WatchDotNet);

string? target;
if (!args.TryGetValue(nameof(target), out target) &&
    !args.TryGetValue("t", out target))
{
    target = watch.Task.Name;
}

RunTarget(target);
