<script lang="ts">
  export let player: Spotify.Player;
  let state: Spotify.WebPlaybackState | undefined;

  player.addListener('player_state_changed', () => {
      player.getCurrentState().then(s => state = s);
  });

  player.getCurrentState().then(s => state = s);
  $: name = state?.track_window.current_track.name;
  $: artists = state?.track_window.current_track.artists.map(a => a.name).join(', ');
  $: albumArt = state?.track_window.current_track.album.images[0].url;
</script>

<div class="flex"></div>
