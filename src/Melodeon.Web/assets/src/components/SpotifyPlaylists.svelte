<svelte:options customElement={{ tag: "spotify-playlists", shadow: "none" }} />

<script lang="ts">
  import type { SimplifiedPlaylist, SpotifyApi } from "@spotify/web-api-ts-sdk";

  import { onDestroy } from "svelte";
  import { spotify } from "../stores";

  let client: SpotifyApi;
  let playlists: SimplifiedPlaylist[] = [];
  const unsubscribe = spotify.client.subscribe(($client) => {
    client = $client;
    client?.currentUser.playlists
      .playlists(10)
      .then((p) => (playlists = p.items))
      .catch((e) => console.error(e));
  });

  onDestroy(unsubscribe);

  function playPlaylist(uri: string) {
    client?.player
      .getPlaybackState()
      .then((state) => client.player.startResumePlayback(state.device.id!, uri))
      .catch((e) => console.error(e));
  }
</script>

<ul class="menu bg-base-200 rounded-box m-2">
  {#each playlists as playlist}
    <li>
      <button type="button" on:click={() => playPlaylist(playlist.uri)}>
        <h6 class="text-sm font-bold">{playlist.name}</h6>
      </button>
    </li>
  {/each}
</ul>
