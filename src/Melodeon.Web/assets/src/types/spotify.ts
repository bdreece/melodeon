declare global {
    interface Window {
        onSpotifyWebPlaybackSDKReady: () => void;
    }

    namespace Spotify {
        interface Options {
            name: string,
            volume: number,
            enableMediaSession?: boolean | undefined,
            getOAuthToken(cb: (token: string) => void): void
        }

        interface WebPlaybackTrack {
            uri: `spotify:track:${string}`,
            id: string,
            type: 'track' | 'episode' | 'ad',
            media_type: 'audio' | 'video',
            name: string,
            is_playable: boolean,
            album: {
                uri: `spotify:album:${string}`
                name: string
                images: {
                    url: string,
                }[]
            },
            artists: {
                uri: `spotify:artist:${string}`
                name: string
            }[]
        }

        interface WebPlaybackState {
            context: {
                uri: string,
                metadata: unknown
            },
            disallows: Record<'pausing' | 'peeking_next' | 'peeking_prev' | 'resuming' | 'seeking' | 'skipping_next' | 'skipping_prev', boolean>
            paused: boolean,
            position: number,
            repeat_mode: number,
            shuffle: boolean,
            track_window: {
                current_track: WebPlaybackTrack,
                previous_tracks: WebPlaybackTrack[],
                next_tracks: WebPlaybackTrack[],
            }
        }

        interface WebPlaybackPlayer {
            device_id: string;
            position: number;
            duration: number;
            track_window: {
                current_track: WebPlaybackTrack,
                previous_tracks: WebPlaybackTrack[],
                next_tracks: WebPlaybackTrack[],
            }
        }

        type WebPlaybackErrorName =
            'initialization_error' |
            'authentication_error' |
            'account_error' |
            'playback_error';

        interface WebPlaybackError {
            message: string
        }

        interface WebPlaybackEventNameMap {
            'ready': WebPlaybackPlayer;
            'not_ready': WebPlaybackPlayer;
            'player_state_changed': WebPlaybackPlayer;
            'autoplay_failed': void;
        }

        class Player {
            constructor(options: Options);
            connect(): Promise<boolean>
            disconnect(): void;
            on(errorName: WebPlaybackErrorName, callback: (e: WebPlaybackError) => void): void;
            addListener<K extends keyof WebPlaybackEventNameMap>(eventName: K, callback: (e: WebPlaybackEventNameMap[K]) => void): void;
            removeListener<K extends keyof WebPlaybackEventNameMap>(eventName: K): void;
            getCurrentState(): Promise<WebPlaybackState>;
            setName(name: string): Promise<void>;
            getVolume(): Promise<number>;
            setVolume(volume: number): Promise<void>;
            pause(): Promise<void>;
            resume(): Promise<void>;
            togglePlay(): Promise<void>;
            seek(positionMs: number): Promise<void>;
            previousTrack(): Promise<void>;
            nextTrack(): Promise<void>;
            activateElement(): Promise<void>;
        }
    }
}

export { }
