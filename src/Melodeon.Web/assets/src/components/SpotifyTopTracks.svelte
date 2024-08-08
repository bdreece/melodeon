<svelte:options customElement={{ tag: "spotify-top-tracks", shadow: "none" }} />

<script lang="ts">
  import type { MaxInt, SpotifyApi, Track } from "@spotify/web-api-ts-sdk";

  import { onDestroy } from "svelte";
  import { spotify } from "../stores";
  import { derived } from "svelte/store";

  let tracks: Track[] = [];
  let client: SpotifyApi;
  let device: string;

  const unsubscribe = derived<
    [typeof spotify.client, typeof spotify.device],
    { client: SpotifyApi; device: string }
  >([spotify.client, spotify.device], ([client, device], set) =>
    set({ client, device }),
  ).subscribe((state) => {
    client = state.client;
    device = state.device;
    state.client?.currentUser
      .topItems("tracks", "short_term")
      .then((t) => (tracks = t.items))
      .catch((e) => console.error(e));
  });

  onDestroy(unsubscribe);

  function playTrack(uri: string) {
    client?.currentUser.profile
      .then((state) =>
        client.player.startResumePlayback(
          device,
          state.context?.uri ?? undefined,
          undefined,
          { uri },
        ),
      )
      .catch((e) => console.error(e));
  }
</script>

<ul class="menu bg-base-200 rounded-box m-2">
  {#each tracks as track}
    <li>
      <button type="button" on:click={() => playTrack(track.uri)}>
        <img
          src={track.album.images.at(0)?.url}
          alt=""
          width="32"
          height="32"
        />

        <h6 class="text-sm font-bold truncate">{track.name}</h6>
        <p class="text-xs hidden xl:inline">{track.artists.map((a) => a.name).join(", ")}</p>
      </button>
    </li>
  {/each}
</ul>
