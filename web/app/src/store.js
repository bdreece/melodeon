/**
 * @template [T=unknown]
 * @extends {CustomEvent<T>}
 */
export class StoreChangeEvent extends CustomEvent {
    static type = 'store-change';

    constructor(value) {
        super(StoreChangeEvent.type, { detail: value });
    }
}

/** @template [T=unknown] */
export default class Store {
    #target = new EventTarget();
    #value;

    /** @param {T} [initialValue=] */
    constructor(initialValue) {
        this.#value = initialValue;
    }

    get value() {
        return this.#value;
    }

    /** @param {StoreChangeEventListenerOrStoreChangeEventListenerObject<T>} listener */
    subscribe(listener) {
        this.#target.addEventListener(StoreChangeEvent.type, listener);
    }

    /** @param {T} value */
    update(value) {
        this.#value = value;
        this.#target.dispatchEvent(new StoreChangeEvent(value));
    }
}

/**
 * @template [T=unknown]
 * @callback StoreChangeEventListener
 * @param {StoreChangeEvent<T>} e
 * @returns {void}
 */

/**
 * @template [T=unknown]
 * @typedef {object} StoreChangeEventListenerObject
 * @prop {StoreChangeEventListener<T>} handleEvent
 */

/**
 * @template [T=unknown]
 * @typedef {StoreChangeEventListener<T> | StoreChangeEventListenerObject<T>} StoreChangeEventListenerOrStoreChangeEventListenerObject
 */
