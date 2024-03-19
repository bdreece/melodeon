<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    export let token: string;
    export let name: string = 'melodeon';
    export let volume: number = 0.5;

    let state: Spotify.WebPlaybackState | undefined;
    let player: Spotify.Player | undefined;
    const dispatch = createEventDispatcher();

    window.onSpotifyWebPlaybackSDKReady = onSpotifyWebPlaybackSDKReady;

    function onSpotifyWebPlaybackSDKReady() {
        player = new Spotify.Player({
            name,
            volume,
            getOAuthToken: cb => cb(token),
        });

        player.addListener('ready', onReady);
        player.addListener('not_ready', onNotReady);
        player.addListener('player_state_changed', onStateChanged);
        player.on('initialization_error', onInitializationError);
        player.on('authentication_error', onAuthenticationError);
        player.on('account_error', onAccountError);
        player.on('playback_error', onPlaybackError);

        player
            .getCurrentState()
            .then(state => onStateChanged(state))
            .catch(err => console.error(err));

        player
            .activateElement()
            .then(() => console.log('activated'))
            .catch(err => console.error(err));
    }

    function onStateChanged(newState: Spotify.WebPlaybackState | null) {
        state = newState ?? undefined;
    }

    function onReady({ device_id }: Spotify.WebPlaybackPlayer) {
        dispatch('device-changed', device_id);
    }

    function onNotReady() {
        dispatch('device-changed', null);
    }

    function onInitializationError({ message }: Spotify.PlayerError) {
        console.error('initialization failed', message);
    }

    function onAuthenticationError({ message }: Spotify.PlayerError) {
        console.error('authentication failed', message);
    }

    function onAccountError({ message }: Spotify.PlayerError) {
        console.error('oauth failed', message);
    }

    function onPlaybackError({ message }: Spotify.PlayerError) {
        console.error('playback failed', message);
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
</script>

<div class="card w-96 bg-base-100 shadow-xl">
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
        <h2 class="card-title">{state?.track_window.current_track.name ?? 'No Track'}</h2>

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
                    class:ti-player-play={!state?.paused}
                    class:ti-player-pause={state?.paused}
                />
            </button>

            <button class="btn btn-primary" on:click={onNextClicked}>
                <i class="ti ti-player-track-next" />
            </button>
        </div>
    </div>
</div>
