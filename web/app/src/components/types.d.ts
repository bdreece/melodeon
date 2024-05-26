import type { default as Backdrop } from './Backdrop'
import type { default as Dropdown } from './Dropdown';
import type { default as RoomMenu } from './RoomMenu';

declare global {
    interface HTMLElementTagNameMap {
        'x-backdrop': Backdrop,
        'x-dropdown': Dropdown,
        'x-room-menu': RoomMenu,
    }
}

export { }
