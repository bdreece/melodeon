import { Session, styles } from '../lib';

/** @satisfies {Record<string, HTMLTemplateElement>} */
const templates = {
    base: document.createElement('template'),
    profile: document.querySelector('#profile-template'),
    login: document.querySelector('#login-template'),
};

templates.base.innerHTML = `
<nav class="nav">
    <ul></ul>
</nav>
`;

export default class AuthMenu extends HTMLElement {
    /** @type {import('../lib/session').EventUnsubscriber} */
    #unsubscribe;
    #outlet;

    constructor() {
        super();
        const shadow = this.attachShadow({
            mode: 'closed',
        });

        shadow.append(
            styles.googleFonts.cloneNode(),
            styles.tablerIcons.cloneNode(),
            styles.custom.cloneNode(),
            templates.base.content.cloneNode(true),
        );

        this.#outlet = shadow.querySelector('ul');
        this.#outlet.append(Session.authenticated ? renderProfile() : renderLogin());
    }

    connectedCallback() {
        this.#unsubscribe = Session.subscribe(e => {
            this.#outlet.replaceChildren(e.detail ? renderProfile() : renderLogin());
        });
    }

    disconnectedCallback() {
        this.#unsubscribe();
    }
}

function renderLogin() {
    return templates.login.content.cloneNode(true);
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

customElements.define('x-auth-menu', AuthMenu);
