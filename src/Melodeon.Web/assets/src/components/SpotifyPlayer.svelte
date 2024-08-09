<svelte:options customElement={{ tag: "spotify-player", shadow: "none" }} />

<script lang="ts">
  import { onDestroy } from "svelte";
  import { get } from "svelte/store";
  import { spotify } from "../stores";

  let device: string | undefined;
  let state: (Spotify.WebPlaybackPlayer & Spotify.WebPlaybackState) | undefined;

  const unsubscribe = {
    device: spotify.device.subscribe((d) => (device = d)),
    state: spotify.state.subscribe((s) => (state = s)),
  };

  onDestroy(() => {
    unsubscribe.device();
    unsubscribe.state();
  });

  $: track = state?.track_window.current_track;
  $: shuffle = state?.shuffle;
  $: repeatMode = state?.repeat_mode;
  $: paused = state?.paused ?? true;
  $: name = track?.name;
  $: artists = track?.artists.map((a) => a.name).join(", ");
  $: albumArt = track?.album.images.at(0)?.url;
  $: position = state?.position;
  $: duration = state?.duration;
  $: progress = duration && position && 100 * (position / duration);

  function onPlayPauseClicked() {
    get(spotify.player)
      ?.togglePlay()
      .catch((e) => console.error(e));
  }

  function onPrevClicked() {
    get(spotify.player)
      ?.previousTrack()
      .catch((e) => console.error(e));
  }

  function onNextClicked() {
    get(spotify.player)
      ?.nextTrack()
      .catch((e) => console.error(e));
  }

  function onShuffleClicked() {
    device &&
      get(spotify.client)
        ?.player.togglePlaybackShuffle(!shuffle, device)
        .catch((e) => console.error(e));
  }

  function onRepeatClicked() {
    let mode: "track" | "context" | "off";
    if (repeatMode === 0) {
      mode = "context";
    } else if (repeatMode === 1) {
      mode = "track";
    } else {
      mode = "off";
    }

    device &&
      get(spotify.client)
        ?.player.setRepeatMode(mode, device)
        .catch((e) => console.error(e));
  }
</script>

<div
  class="grid grid-cols-[auto_auto_auto] lg:grid-cols-[auto_auto_auto_auto_auto] gap-2 w-full"
>
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
          {#if paused}
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

  <div class="flex-none hidden lg:block">
    <ul class="menu menu-horizontal items-center h-full">
      <li>
        {#if shuffle}
          <button class="btn btn-accent btn-xs" on:click={onShuffleClicked}>
            <i class="iconify tabler--arrows-shuffle"></i>
          </button>
        {:else}
          <button class="btn btn-ghost btn-xs" on:click={onShuffleClicked}>
            <i class="iconify tabler--arrows-shuffle"></i>
          </button>
        {/if}
      </li>

      <li>
        {#if state?.repeat_mode === 0}
          <button class="btn btn-ghost btn-xs" on:click={onRepeatClicked}>
            <i class="iconify tabler--repeat-off"></i>
          </button>
        {:else if state?.repeat_mode === 1}
          <button class="btn btn-accent btn-xs" on:click={onRepeatClicked}>
            <i class="iconify tabler--repeat"></i>
          </button>
        {:else}
          <button class="btn btn-accent btn-xs" on:click={onRepeatClicked}>
            <i class="iconify tabler--repeat-once"></i>
          </button>
        {/if}
      </li>
    </ul>
  </div>
</div>
