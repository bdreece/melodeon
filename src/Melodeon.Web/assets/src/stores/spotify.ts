import type { AccessToken } from "@spotify/web-api-ts-sdk";

import { derived, get, readable, writable } from "svelte/store";
import { SpotifyApi } from "@spotify/web-api-ts-sdk";

export const token = readable<AccessToken>(undefined, set => {
    fetch("/host/auth/token", { credentials: "include" })
        .then(res => res.json())
        .then(token =>
            set({
                ...token,
                token_type: "Bearer",
            }),
        )
        .catch(e => console.error(e));
});

export const client = derived<typeof token, SpotifyApi>(
    token,
    ($token, set) => {
        if ($token) {
            const sdk = SpotifyApi.withAccessToken(
                import.meta.env.VITE_SPOTIFY_CLIENT_ID,
                $token,
            );
            set(sdk);
        }
    },
);

export const player = writable<Spotify.Player | undefined>();

Object.assign(window, {
    onSpotifyWebPlaybackSDKReady() {
        const $player = new Spotify.Player({
            name: "melodeon",
            volume: 0.6,
            getOAuthToken(cb) {
                const $token = get(token);
                cb($token?.access_token);
            },
        });

        $player
            .connect()
            .then(connected => connected && console.log("player connected"))
            .catch(e => console.error(e));

        player.set($player);
    },
});

export const device = derived<typeof player, string | undefined>(
    player,
    ($player, set) => {
        $player?.addListener("ready", ({ device_id }) => set(device_id));
    },
);

export const state = derived<
    typeof player,
    (Spotify.WebPlaybackState & Spotify.WebPlaybackPlayer) | undefined
>(player, ($player, set) => {
    $player?.addListener("player_state_changed", playerState => {
        $player.getCurrentState().then(playbackState =>
            set({
                ...playerState,
                ...playbackState,
            }),
        );
    });
});

derived<
    [typeof client, typeof device],
    {
        client: SpotifyApi;
        device: string;
    }
>([client, device], ([$client, $device], set) =>
    set({
        client: $client,
        device: $device,
    }),
).subscribe(({ client, device }) => {
    device &&
        client?.player.transferPlayback([device]).catch(e => console.error(e));
});
