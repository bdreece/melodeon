export default class Styles {
    /** @type {HTMLLinkElement} */
    static #custom = document.getElementById('custom-styles');
    /** @type {HTMLLinkElement} */
    static #googleFonts = document.getElementById('google-fonts');
    /** @type {HTMLLinkElement} */
    static #tablerIcons = document.getElementById('tabler-icons');

    static get custom() {
        return cloneLink(this.#custom);
    }
    static get googleFonts() {
        return cloneLink(this.#googleFonts);
    }
    static get tablerIcons() {
        return cloneLink(this.#tablerIcons);
    }

    static *links() {
        yield this.custom;
        yield this.googleFonts;
        yield this.tablerIcons;
    }
}

/** @param {HTMLLinkElement} link */
function cloneLink(link) {
    /** @type {HTMLLinkElement} */
    const newLink = link.cloneNode();
    newLink.removeAttribute('id');
    return newLink;
}
