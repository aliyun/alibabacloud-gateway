import { Readable } from 'stream';

export async function readableStreamToBuffer(stream: Readable): Promise<Buffer> {
    const chunks: Buffer[] = [];
    for await (const chunk of stream) {
        chunks.push(chunk as Buffer);
    }
    return Buffer.concat(chunks);
}

export async function readableStreamWithPrependBuffer(stream: Readable, prependBuffer: Buffer): Promise<Buffer> {
    const chunks: Buffer[] = [];
    for await (const chunk of stream) {
        chunks.push(chunk as Buffer);
    }
    return Buffer.concat([prependBuffer, ...chunks]);
}