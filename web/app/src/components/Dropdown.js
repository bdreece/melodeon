import Backdrop from './Backdrop';
import { styles } from '../lib';

const template = document.createElement('template');
template.innerHTML = `
<div class="dropdown">
    <div class="dropdown__toggle">
        <slot name="toggle"></slot>
    </div>

    <div class="dropdown__items">
        <slot name="items"></slot>
    </div>
</div>
`;

export default class Dropdown extends HTMLElement {
    static observedAttributes = ['visible'];

    #toggle;
    #items;

    get visible() {
        return this.hasAttribute('visible');
    }
    set visible(value) {
        if (value) {
            this.setAttribute('visible', 'visible');
        } else {
            this.removeAttribute('visible');
        }
    }

    constructor() {
        super();

        const shadow = this.attachShadow({
            mode: 'closed',
            delegatesFocus: true,
        });

        shadow.append(
            styles.googleFonts.cloneNode(),
            styles.tablerIcons.cloneNode(),
            styles.custom.cloneNode(),
            template.content.cloneNode(true),
        );

        this.#toggle = shadow.querySelector('.dropdown__toggle');
        this.#items = shadow.querySelector('.dropdown__items');
    }

    connectedCallback() {
        this.#toggle.addEventListener('click', () => this.toggle());

        this.#render(this.visible);
    }

    /**
     * @param {string} name
     * @param {string} value
     */
    attributeChangedCallback(name, _, value) {
        if (name !== 'visible') {
            return;
        }

        this.#render(!!value);
    }

    show() {
        this.visible = true;
    }
    hide() {
        this.visible = false;
    }
    toggle() {
        this.visible = !this.visible;
    }

    /** @param {boolean} visible */
    #render(visible) {
        if (visible) {
            this.#items.classList.add('dropdown__items--visible');
            Backdrop.instance.tinted = false;
            Backdrop.instance.visible = true;
            Backdrop.instance.addEventListener('click', () => this.hide(), { once: true });
        } else {
            this.#items.classList.remove('dropdown__items--visible');
            Backdrop.instance.visible = false;
        }
    }
}

customElements.define('x-dropdown', Dropdown);
