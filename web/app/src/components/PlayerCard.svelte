<script lang="ts">
    import { store } from '../modules/spotify';

    let player: Spotify.Player | undefined;
    let state: Spotify.WebPlaybackState | undefined;

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
        }
    });
</script>

<div class="card w-96 mx-auto bg-base-100 shadow-xl">
    <figure class="px-10 pt-10">
        {#if state}
            <img
                src={state.track_window.current_track.album.images.at(0)?.url ??
                    ''}
                alt="Album Art"
                class="rounded-lg"
            />
        {:else}
            <i class="ti ti-album-off" />
        {/if}
    </figure>
    <div class="card-body items-center text-center">
        <h2 class="card-title">
            {state?.track_window.current_track.name ?? 'No Track'}
        </h2>

        <p>
            {state?.track_window.current_track.artists.at(0)?.name ?? ''}
        </p>

        <div class="card-actions">
            <button class="btn btn-primary" on:click={onPreviousClicked}>
                <i class="ti ti-player-track-prev" />
            </button>

            <button class="btn btn-primary" on:click={onToggleClicked}>
                <i
                    class="ti"
                    class:ti-player-play={state?.paused}
                    class:ti-player-pause={!state?.paused}
                />
            </button>

            <button class="btn btn-primary" on:click={onNextClicked}>
                <i class="ti ti-player-track-next" />
            </button>
        </div>
    </div>
</div>
