import { writable } from 'svelte/store';

export type User = {
    href: string;
    id: string;
    type: string;
    uri: string;
    country: string;
    display_name: string;
    images: {
        url: string;
        height: number;
        width: number;
    }[];
}

export const store = writable<User>();
