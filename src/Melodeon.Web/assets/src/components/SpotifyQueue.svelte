<svelte:options customElement={{ tag: "spotify-queue", shadow: "none" }} />

<script lang="ts">
  import type {
    PlaybackState,
    SpotifyApi,
    Track,
  } from "@spotify/web-api-ts-sdk";

  import { onDestroy } from "svelte";
  import { spotify } from "../stores";

  let client: SpotifyApi;
  let state: PlaybackState;
  let tracks: Track[] = [];

  function getData() {
    Promise.all([
      client?.player.getPlaybackState().then((s) => (state = s)),

      client?.player
        .getUsersQueue()
        .then((result) => (tracks = result.queue as Track[])),
    ]).catch((e) => console.error(e));
  }

  const unsubscribe = spotify.client.subscribe((c) => {
    client = c;
  });

  const interval = setInterval(getData, 5000);

  onDestroy(() => {
    clearInterval(interval);
    unsubscribe();
  });

  function playTrack(uri: string) {
    state &&
      client?.player
        .startResumePlayback(state.device.id ?? "", state.context?.uri, undefined, { uri })
        .catch((e) => console.error(e));
  }
</script>

<ul class="menu bg-base-200 rounded-box m-2">
  {#each tracks as track}
    <li>
      <button class="w-full" type="button" on:click={() => playTrack(track.uri)}>
        <img
          src={track.album.images.at(0)?.url}
          alt=""
          width="32"
          height="32"
        />
        <h6 class="text-sm truncate font-bold">{track.name}</h6>
        <p class="text-xs hidden xl:inline">{track.artists.map((a) => a.name).join(", ")}</p>
      </button>
    </li>
  {/each}
</ul>
