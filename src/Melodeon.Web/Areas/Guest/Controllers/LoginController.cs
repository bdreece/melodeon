using System.Security.Claims;

using FluentResults.Extensions.FluentAssertions;

using MediatR;

using Melodeon.Domain;
using Melodeon.UseCases;
using Melodeon.Web.Areas.Guest.Models;

using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Melodeon.Web.Areas.Guest.Controllers;

[AllowAnonymous]
[Area(nameof(Guest))]
public sealed class LoginController : Controller
{
    private readonly ISender _mediator;

    public LoginController(ISender mediator)
    {
        _mediator = mediator;
    }

    public IActionResult Index([FromQuery] string code) => View(new LoginModel { Code = code });

    [HttpPost]
    public async Task<IActionResult> IndexAsync([FromForm] LoginModel form)
    {
        var request = new JoinRoom.Request(form.Code, form.DisplayName);
        var result = await _mediator.Send(request);
        result.Should().BeSuccess();

        var identity = new ClaimsIdentity(new[]
        {
            new Claim(ClaimTypes.NameIdentifier, Guid.NewGuid().ToString()),
            new Claim(ClaimTypes.Name, form.DisplayName),
            new Claim(ClaimTypes.Role, Role.Guest),
        });

        await HttpContext.SignInAsync(
            scheme: MelodeonAuthenticationDefaults.GuestScheme,
            principal: new ClaimsPrincipal(identity));

        return Redirect("/guest");
    }
}