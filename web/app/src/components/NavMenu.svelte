<script lang="ts">
    import type { User } from '../modules/user';
    import { store } from '../modules/user';

    let user: User | undefined;

    store.subscribe(u => user = u);
</script>

{#if user}
    <div class="flex items-center justify-end gap-2 px-1">
        <a class="btn btn-primary" href="/host/wizard">Create Room</a>
    </div>

    <div class="dropdown dropdown-end">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
            <div class="w-10 rounded-full">
                {#each user.images as img}
                    <img
                        src={img.url}
                        alt="Profile"
                        height={img.height}
                        width={img.width}
                    />
                {:else}
                    <p>{user.display_name}</p>
                {/each}
            </div>
        </div>

        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <ul
            tabindex="0"
            class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52"
        >
            <li><a href="/host/settings">Settings</a></li>

            <li><a href="/logout">Logout</a></li>
        </ul>
    </div>
{:else}
    <div class="flex items-center justify-end gap-4 px-1">
        <form
            hx-get="/guest/room"
            hx-push-url="true"
            class="flex place-items-center gap-2"
        >
            <input
                class="input input-bordered"
                name="code"
                type="search"
                placeholder="Code"
            />
            <button class="btn btn-ghost" type="submit">Join Room</button>
        </form>

        <a hx-boost="false" href="/login" class="btn btn-primary">
            <i class="ti ti-brand-spotify text-xl"></i>
            Login with Spotify
        </a>
    </div>
{/if}
