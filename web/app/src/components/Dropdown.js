import { Styles } from '../lib';

const template = document.createElement('template');
template.innerHTML = `
<div class="dropdown">
    <button type="button" class="dropdown__toggle">
        <slot name="toggle"></slot>
    </button>

    <ul class="dropdown__items">
        <slot name="items"></slot>
    </ul>
</div>
`;

export default class Dropdown extends HTMLElement {
    static observedAttributes = ['open'];

    #toggle;
    #items;

    get open() {
        return !!this.getAttribute('open');
    }
    set open(value) {
        this.setAttribute('open', value ? 'true' : '');
    }

    constructor() {
        super();

        const shadow = this.attachShadow({
            mode: 'closed',
            slotAssignment: 'named',
            delegatesFocus: true,
        });

        shadow.append(...Styles.links(), template.content.cloneNode(true));

        this.#toggle = shadow.querySelector('button');
        this.#items = shadow.querySelector('.dropdown__items');
    }

    connectedCallback() {
        if (this.open) {
            this.#items.classList.add('dropdown__items--open');
        }

        this.#toggle.addEventListener('click', e => {
            this.#items.classList.toggle('dropdown__items--open');
            this.dispatchEvent(new PointerEvent('click', e));
        });
    }

    /**
     * @param {string} name
     * @param {string} value
     */
    attributeChangedCallback(name, _, value) {
        if (name !== 'open') {
            return;
        } else if (value) {
            this.#items.classList.add('dropdown__items--open');
        } else {
            this.#items.classList.remove('dropdown__items--open');
        }
    }
}

customElements.define('x-dropdown', Dropdown);
