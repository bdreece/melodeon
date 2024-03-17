const template = document.createElement('template');
template.innerHTML = `
<style>
    @import '/app/styles.css';
</style>

<div>
    <button type="button" class="btn btn-primary btn-prev">Prev</button>
    <button type="button" class="btn btn-primary btn-play-pause">Play/Pause</button>
    <button type="button" class="btn btn-primary btn-next">Next</button>
</div>
`;

export default class SpotifyPlayer extends HTMLElement {
    static ready = false;

    #player: Spotify.Player | undefined;

    get player() { return this.#player; }

    constructor() {
        super();

        const shadow = this.attachShadow({ mode: 'open' });
        shadow.append(template.content.cloneNode(true));
    }

    connectedCallback() {
        if (!SpotifyPlayer.ready) {
            throw new Error('spotify player not ready');
        }

        this.#player = new Spotify.Player({
            name: 'melodeon',
            getOAuthToken: cb => { cb(''); },
            volume: 0.5,
        });

        this.#player.addListener('ready', ({ device_id }) =>
            console.log('Spotify device available', device_id));

        this.#player.addListener('not_ready', ({ device_id }) =>
            console.log('Spotify device unavailable', device_id));

        const [prevTrack, playPause, nextTrack] = [
            this.shadowRoot!.querySelector<HTMLButtonElement>('.btn-prev-track'),
            this.shadowRoot!.querySelector<HTMLButtonElement>('.btn-play-pause'),
            this.shadowRoot!.querySelector<HTMLButtonElement>('.btn-next-track'),
        ];

        prevTrack?.addEventListener('click', () =>
            this.#player?.previousTrack()
                .then(() => console.log('started previous track'))
                .catch(e => console.error(e)));

        playPause?.addEventListener('click', () =>
            this.#player?.togglePlay()
                .then(() => console.log('toggled play/pause'))
                .catch(e => console.error(e)));

        nextTrack?.addEventListener('click', () =>
            this.#player?.nextTrack()
                .then(() => console.log('started next track'))
                .catch(e => console.error(e)));
    }

    static {
        window.onSpotifyWebPlaybackSDKReady = () => this.ready = true;
    }
}

customElements.define('spotify-player', SpotifyPlayer);

