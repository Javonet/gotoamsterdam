

using System.Text.Json;
using System.Text.Json.Serialization;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();



app.MapGet("/compute", () =>
{
    int result = robotConnector.Robot.Compute(4,5);
    Result r = new Result();
    r.result=result;
    return JsonSerializer.Serialize(r);
})
.WithName("compute")
.WithOpenApi();

app.Run();

class Result {
    public int result;
}
