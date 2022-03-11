// snowprints

function b64EncodeUnicode(str) {
    return btoa(encodeURIComponent(str));
};

function UnicodeDecodeB64(str) {
    return decodeURIComponent(atob(str));
};

function getTimeSince(timestampDelta: number, contextDelta: number) {
    return performance.now() - contextDelta + timestampDelta;
}


class Snowprint {
    private rotation: number = 0;

    private id: number;
    private modulo: number;
    private timestampDelta: number;
    private contextDelta: number;

    constructor(id: number, date: number, modulo: number) {
        this.id = id;
        this.modulo = modulo;
        this.timestampDelta = Date.now() - date;
        this.contextDelta = performance.now();
    }

    createSnowprint(): string {
        this.rotation += 1;
        this.rotation = this.rotation % this.modulo;
        const timestamp = getTimeSince(this.timestampDelta, this.contextDelta);

        return `${timestamp}_${this.id}_${this.rotation}`;
    }
}