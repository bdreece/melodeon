const alpha = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'

export function randomString(n: number) {
    let s = '';
    for (let i = 0; i < n; i++) {
        let r = Math.floor(Math.random() * alpha.length);
        s += alpha.at(r) ?? '';
    }

    return s;
}
