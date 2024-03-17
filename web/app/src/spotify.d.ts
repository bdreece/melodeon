declare global {
    interface Window {
        onSpotifyWebPlaybackSDKReady: () => void;
    }

    namespace Spotify {
        interface PlayerOptions {
            name: string;
            volume?: number;
            enableMediaSession?: boolean | undefined;
            getOAuthToken(cb: (token: string) => void): void;
        }

        interface WebPlaybackPlayer {
            device_id: string;
        }

        interface WebPlaybackState {
            context: {
                uri: string,
                metadata: unknown,
            },
            disallows: {
                pausing: boolean,
                peeking_next: boolean,
                peeking_prev: boolean,
                resuming: boolean,
                seeking: boolean,
                skipping_next: boolean,
                skipping_prev: boolean
            },
            paused: boolean,
            position: number,
            repeat_mode: number,
            shuffle: boolean,
            track_window: {
                current_track: WebPlaybackTrack,
                previous_tracks: WebPlaybackTrack[],
                next_tracks: WebPlaybackTrack[]
            }
        }

        interface WebPlaybackTrack {
            uri: string,
            id: string | null,
            type: 'track' | 'episode' | 'ad',
            media_type: 'audio' | 'video',
            name: string,
            is_playable: boolean,
            album: {
                uri: string,
                name: string,
                images: {
                    url: string
                }[];
            },
            artists: {
                uri: string,
                name: string
            }[]
        }

        type PlayerEventNameMap = {
            ready: WebPlaybackPlayer;
            not_ready: WebPlaybackPlayer
            player_state_changed: WebPlaybackState
        }

        declare class Player {
            constructor(options: PlayerOptions);

            addListener<
                K extends keyof PlayerEventNameMap,
                E extends PlayerEventNameMap[K]
            >(name: K, cb: (evt: E) => void): void;

            connect(): void;
            disconnect(): void;
            getCurrentState(): Promise<WebPlaybackState | null>;
            setName(name: string): Promise<void>;
            getVolume(): Promise<number>;
            setVolume(volume: number): Promise<void>;
            pause(): Promise<void>;
            resume(): Promise<void>;
            togglePlay(): Promise<void>;
            seek(ms: number): Promise<void>;
            previousTrack(): Promise<void>;
            nextTrack(): Promise<void>;
        }
    }
}

export { }
