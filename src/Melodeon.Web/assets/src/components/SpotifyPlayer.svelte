<svelte:options customElement={{ tag: 'spotify-player', shadow: 'none' }} />

<script lang="ts">
  import * as stores from "../stores";

  let volume: number;
  let player: Spotify.Player | undefined;
  let state: (Spotify.WebPlaybackPlayer & Spotify.WebPlaybackState) | undefined;
  stores.spotify.state.subscribe((s) => (state = s));
  stores.spotify.player.subscribe((p) =>
    p?.activateElement().then(() => (player = p)),
  );

  $: track = state?.track_window.current_track;
  $: name = track?.name;
  $: artists = track?.artists.map((a) => a.name);
  $: albumArt = track?.album.images.at(0)?.url;
  $: position = state?.position;
  $: duration = state?.duration;
  $: progress = duration && position && 100 * (position / duration);

  function onPlayPauseClicked() {
    player?.togglePlay().catch((e) => console.error(e));
  }

  function onPrevClicked() {
    player?.previousTrack().catch((e) => console.error(e));
  }

  function onNextClicked() {
    player?.nextTrack().catch((e) => console.error(e));
  }

  function setVolume(value) {
  }
</script>

<div class="grid grid-cols-[auto_auto_auto] lg:grid-cols-[auto_auto_auto_auto] gap-2 w-full">
  <div class="flex-none grid grid-cols-2 gap-2">
    <img
      class="row-span-2 place-self-center"
      src={albumArt}
      alt=""
      width="32"
      height="32"
    />
    <div class="grid grid-rows-subgrid row-span-2 truncate">
      <h6 class="text-sm font-bold">{name}</h6>
      <p class="text-xs">{artists}</p>
    </div>
  </div>

  <div class="flex-1 hidden lg:flex items-center">
    <progress class="progress w-full" value={progress ?? 0} max="100"
    ></progress>
  </div>

  <div class="flex-none w-fit">
    <ul class="menu menu-horizontal">
      <li>
        <button on:click={onPrevClicked}>
          <i class="iconify tabler--player-track-prev"></i>
        </button>
      </li>

      <li>
        <button on:click={onPlayPauseClicked}>
          {#if state?.paused}
            <i class="iconify tabler--player-play"></i>
          {:else}
            <i class="iconify tabler--player-pause"></i>
          {/if}
        </button>
      </li>

      <li>
        <button on:click={onNextClicked}>
          <i class="iconify tabler--player-track-next"></i>
        </button>
      </li>
    </ul>
  </div>
</div>
