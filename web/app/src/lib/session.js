/**
 * @callback EventListener
 * @param {SessionChangeEvent} e
 * @returns {void}
 *
 * @callback EventUnsubscriber
 * @param {boolean|EventListenerOptions|undefined} [options=]
 * @returns {void}
 */

/**
 * @typedef {object} SessionData
 * @prop {string} username
 * @prop {string} room
 * @prop {string|null|undefined} [image=]
 */

export default class Session {
    /** @type {SessionData|undefined} */
    static #data;
    static #observer;
    static #target = new EventTarget();

    static get authenticated() {
        return !!this.#data;
    }

    static get username() {
        return this.#data?.username;
    }
    static get image() {
        return this.#data?.image;
    }
    static get room() {
        return this.#data?.room;
    }

    static {
        const el = document.getElementById('session');
        if (!el) {
            throw new Error('session element not found!');
        }

        if (el.textContent) {
            this.#data = JSON.parse(el.textContent);
            console.dir(this.#data);
        }

        this.#observer = new MutationObserver(mutations => {
            for (const mutation of mutations) {
                if (mutation.type !== 'characterData' || !mutation.target.textContent) {
                    continue;
                }

                this.#data = JSON.parse(mutation.target.textContent);
                console.dir(this.#data);

                this.#target.dispatchEvent(new SessionChangeEvent(this.#data));
            }
        });

        this.#observer.observe(el, {
            characterData: true,
        });
    }

    /**
     * @param {EventListener} listener
     * @param {boolean|AddEventListenerOptions|undefined} [options=]
     * @returns {EventUnsubscriber}
     */
    static subscribe(listener, options) {
        this.#target.addEventListener(SessionChangeEvent.name, listener, options);

        return options => this.#target.removeEventListener(SessionChangeEvent.name, listener, options);
    }
}

/** @extends {CustomEvent<SessionData|undefined>} */
export class SessionChangeEvent extends CustomEvent {
    static name = 'session-change';

    /** @param {SessionData|undefined} [session=] */
    constructor(session) {
        super(SessionChangeEvent.name, { detail: session });
    }
}
