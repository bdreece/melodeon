<svelte:options customElement={{ tag: "spotify-volume", shadow: "none" }} />

<script lang="ts">
    import { onDestroy } from "svelte";
    import { get } from "svelte/store";
    import { spotify } from "../stores";

    let volume: number = 0;
    let player: Spotify.Player | undefined;
    const unsubscribe = spotify.player.subscribe(p => {
        player = p;
        player?.addListener("ready", () =>
            player
                ?.getVolume()
                .then(v => (volume = v * 100))
                .catch(e => console.error(e)),
        );
    });

    function changeVolume(e: ChangeEvent) {
        volume = +(e.target as HTMLInputElement).value;
        player?.setVolume(volume / 100).catch(e => console.error(e));
    }

    onDestroy(() => {
        unsubscribeState();
        unsubscribePlayer();
    });
</script>

<input
    class="range range-xs"
    type="range"
    min="0"
    max="100"
    value={volume}
    on:change={changeVolume}
/>
