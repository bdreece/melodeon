import { writable } from 'svelte/store';
import { randomString } from './rand';
import { store as userStore } from './user';

export const code = randomString(6);
export const store = writable<{
    player?: Spotify.Player | undefined;
    deviceID?: string | undefined;
}>({});

userStore.subscribe(user => {
    if (!user) {
        store.update(s => {
            s.player?.disconnect();
            return {
                ...s,
                player: undefined,
            };
        });
    }
});

export function initialize(token: string) {
    window.onSpotifyWebPlaybackSDKReady = () => {
        const player = new Spotify.Player({
            name: `melodeon - ${code}`,
            volume: 0.5,
            getOAuthToken: cb => cb(token),
        });

        store.set({ player });

        player.addListener('ready', ({ device_id: deviceID }) => store.update(s => ({ ...s, deviceID })));
        player.addListener('not_ready', () => store.update(s => ({ ...s, deviceID: undefined })));

        player.on('initialization_error', onInitializationError);
        player.on('authentication_error', onAuthenticationError);
        player.on('playback_error', onPlaybackError);
        player.on('account_error', onAccountError);

        player.connect();
    }
}

function onInitializationError({ message }: Spotify.PlayerError) {
    console.error('spotify initialization failed:', message);
}

function onAuthenticationError({ message }: Spotify.PlayerError) {
    console.error('spotify authentication failed:', message);
}

function onAccountError({ message }: Spotify.PlayerError) {
    console.error('spotify oauth failed:', message);
}

function onPlaybackError({ message }: Spotify.PlayerError) {
    console.error('spotify playback failed:', message);
}

