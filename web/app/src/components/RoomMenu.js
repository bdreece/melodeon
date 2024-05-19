import { Session, Styles } from '../lib';

const templates = {
    base: document.createElement('template'),
    profile: document.createElement('template'),
    login: document.createElement('template'),
};

templates.base.innerHTML = `
<nav class="nav">
    <ul></ul>
</nav>
`;

templates.profile.innerHTML = `
<li class="nav__item">
    <x-dropdown class="profile">
        <figure class="profile__badge" slot="toggle">
            <img alt="Profile" width="32" height="32">
            <figcaption></figcaption>
        </figure>
    
        <li slot="items">
            <a href="/settings">Settings</a>
        </li>
    
        <li slot="items">
            <a href="/logout">Logout</a>
        </li>
    </x-dropdown>
</li>
`;

templates.login.innerHTML = `
<li class="nav__item">
    <a class="nav__link" href="/join">Join Room</a>
</li>

<li class="nav__item">
    <a class="nav__link" href="/login">Create Room</a>
</li>
`;

export default class RoomMenu extends HTMLElement {
    /** @type {import('../lib/session').EventUnsubscriber} */
    #unsubscribe;
    #outlet;

    constructor() {
        super();
        const shadow = this.attachShadow({ mode: 'open' });
        shadow.append(...Styles.links(), templates.base.content.cloneNode(true));

        this.#outlet = shadow.querySelector('ul');
        this.#outlet.append(Session.authenticated ? renderProfile() : templates.login.content.cloneNode(true));
    }

    connectedCallback() {
        this.#unsubscribe = Session.subscribe(e => {
            this.#outlet.replaceChildren(e.detail ? renderProfile() : templates.login.content.cloneNode(true));
        });
    }

    disconnectedCallback() {
        this.#unsubscribe();
    }
}

function renderProfile() {
    /** @type {DocumentFragment} */
    const fragment = templates.profile.content.cloneNode(true);

    const img = fragment.querySelector('img');
    const caption = fragment.querySelector('figcaption');
    caption.textContent = Session.username;
    img.style.display = Session.image ? '' : 'none';
    img.src = Session.image;

    return fragment;
}

customElements.define('x-room-menu', RoomMenu);
