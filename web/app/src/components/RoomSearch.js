const template = document.createElement('template');
template.innerHTML = `
<form class="form">
    <label class="form__control">
        Join Room:

        <input
            type="search"
            name="code"
            placeholder="Code"
            maxlength="255"
            required
        >
    </label>

    <button type="submit" class="button">
        Submit
    </button>
</form>
`;

export default class RoomSearch extends HTMLElement {
    #form;

    constructor() {
        super();
        const shadow = this.attachShadow({
            mode: 'closed',
            delegatesFocus: true,
        });

        shadow.append(template.content.cloneNode(true));
        this.#form = shadow.querySelector('form');
    }

    connectedCallback() {
        this.#form.addEventListener('submit', () => {
            this.#search().catch(e => console.error(e));
        });
    }

    async #search() {
        const data = new FormData(this.#form);
        const params = new URLSearchParams(data);
        const res = await fetch(`/join?${params}`, {
            credentials: 'include',
        });

        if (!res.ok) {
            throw new Error('Room not found!');
        }

        const code = data.get('code');
        location.replace(`/room/${code}`);
    }
}
