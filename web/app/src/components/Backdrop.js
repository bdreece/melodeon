import { styles } from '../lib';

const template = document.createElement('template');
template.innerHTML = `<div class="backdrop"></div>`;

export default class Backdrop extends HTMLElement {
    static observedAttributes = ['visible', 'tinted'];

    #element;

    static get instance() {
        return document.querySelector('x-backdrop');
    }

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

    get tinted() {
        return this.hasAttribute('tinted');
    }
    set tinted(value) {
        if (value) {
            this.setAttribute('tinted', 'tinted');
        } else {
            this.removeAttribute('tinted');
        }
    }

    constructor() {
        super();

        const shadow = this.attachShadow({ mode: 'closed' });
        shadow.append(
            styles.googleFonts.cloneNode(),
            styles.tablerIcons.cloneNode(),
            styles.custom.cloneNode(),
            template.content.cloneNode(true),
        );

        this.#element = shadow.querySelector('div');
    }

    connectedCallback() {
        this.#updateVisible(this.visible);
        this.#updateTinted(this.tinted);
    }

    /**
     * @param {string} name
     * @param {string} value
     */
    attributeChangedCallback(name, _, value) {
        switch (name) {
            case 'visible':
                this.#updateVisible(!!value);
                break;
            case 'tinted':
                this.#updateTinted(!!value);
                break;
            default:
                break;
        }
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
    #updateVisible(visible) {
        if (visible) {
            this.#element.classList.add('backdrop--visible');
            this.#element.addEventListener('click', () => this.hide(), { once: true });
        } else {
            this.#element.classList.remove('backdrop--visible');
        }
    }

    /** @param {boolean} tinted */
    #updateTinted(tinted) {
        if (tinted) {
            this.#element.classList.add('backdrop--tinted');
        } else {
            this.#element.classList.remove('backdrop--tinted');
        }
    }
}

customElements.define('x-backdrop', Backdrop);
