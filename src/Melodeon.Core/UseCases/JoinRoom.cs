using FluentAssertions;

using FluentResults;

using FluentValidation;

using MediatR;

using Melodeon.Data;
using Melodeon.Domain;

namespace Melodeon.UseCases;

public static class JoinRoom
{
    public sealed record Request(string Code, string DisplayName) : IRequest<Result>
    {
        internal sealed class Validator : AbstractValidator<Request>
        {
            public Validator()
            {
                RuleFor(x => x.Code).NotEmpty().Length(8);
                RuleFor(x => x.DisplayName).NotEmpty();
            }
        }
    }

    internal sealed class Handler : IRequestHandler<Request, Result>
    {
        private readonly IRepository<Room> _rooms;

        public Handler(IRepository<Room> rooms)
        {
            _rooms = rooms;
        }

        public async Task<Result> Handle(Request request, CancellationToken ct = default)
        {
            var room = await _rooms
                .Where(r => r.Code.StartsWith(request.Code))
                .FirstOrDefaultAsync(ct);

            Result.FailIf(room is null, "room not found");

            room!.Guests.Add(request.DisplayName);

            await _rooms.SaveChangesAsync(ct);
            return Result.Ok();
        }
    }
}