<script lang="ts">
    import { store } from '../modules/spotify';

    let code: string | undefined;
    let player: Spotify.Player | undefined;
    let state: Spotify.WebPlaybackState | undefined;

    let track:
        | {
              playing: boolean;
              name: string;
              artists: string;
              album: {
                  previous?: string | undefined;
                  current?: string | undefined;
                  next?: string | undefined;
              };
          }
        | undefined;

    $: {
        if (state) {
            track = {
                playing: state.paused,
                name: state.track_window.current_track.name,
                artists:
                    state.track_window.current_track.artists
                        .map(a => a.name)
                        .join(', ') ?? 'No Artists',
                album: {
                    previous: state.track_window.previous_tracks
                        .at(0)
                        ?.album.images.at(0)?.url,
                    current:
                        state.track_window.current_track.album.images.at(0)
                            ?.url,
                    next: state.track_window.next_tracks
                        .at(0)
                        ?.album.images.at(0)?.url,
                },
            };
        }
    }

    function onStateChanged(newState: Spotify.WebPlaybackState | null) {
        state = newState ?? undefined;
    }

    function onPreviousClicked() {
        player
            ?.previousTrack()
            .then(() => console.log('playing previous track'))
            .catch(e => console.error(e));
    }

    function onToggleClicked() {
        player
            ?.togglePlay()
            .then(() => console.log('toggled play/pause'))
            .catch(e => console.error(e));
    }

    function onNextClicked() {
        player
            ?.nextTrack()
            .then(() => console.log('playing next track'))
            .catch(e => console.error(e));
    }

    store.subscribe(({ player: newPlayer }) => {
        if (!player && newPlayer) {
            player = newPlayer;
            player.addListener('player_state_changed', onStateChanged);

            player
                .getCurrentState()
                .then(state => onStateChanged(state))
                .catch(err => console.error(err));
        } else if (player && !newPlayer) {
            player = undefined;
            state = undefined;
        }
    });
</script>

<div class="flex w-full justify-between items-center gap-4 p-2 bg-base-100">
    {#if player}
        {#if !track}
            <div>No Track Found 😞</div>

            {#if code}
                <div>Room Code: {code}</div>
            {/if}
        {:else}
            <div class="flex items-center gap-2">
                {#if track.album.previous}
                    <link
                        rel="preload"
                        as="image"
                        href={track.album.previous}
                    />
                {/if}

                {#if track.album.next}
                    <link rel="preload" as="image" href={track.album.next} />
                {/if}

                {#if track.album.current}
                    <img
                        src={track.album.current}
                        alt="Album"
                        height="48"
                        width="48"
                        class="rounded-lg"
                    />
                {/if}

                <div class="flex items-center gap-1">
                    <h3 class="font-bold">{track.name}</h3>
                    <span>&mdash;</span>
                    <p class="text-sm">{track.artists}</p>
                </div>
            </div>

            <div class="flex items-center h-1/2 gap-2">
                <button class="btn btn-primary" on:click={onPreviousClicked}>
                    <i class="ti ti-player-track-prev" />
                </button>

                <button class="btn btn-primary" on:click={onToggleClicked}>
                    <i
                        class="ti"
                        class:ti-player-play={track.playing}
                        class:ti-player-pause={!track.playing}
                    />
                </button>

                <button class="btn btn-primary" on:click={onNextClicked}>
                    <i class="ti ti-player-track-next" />
                </button>
            </div>
        {/if}
    {/if}
</div>
