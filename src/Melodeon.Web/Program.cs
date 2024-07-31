var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.AddMelodeon()
    .ConfigureInfrastructureDefaults()
    .ConfigureWebDefaults();

var app = builder.Build();

app.UseMelodeonDefaults();

app.Run();