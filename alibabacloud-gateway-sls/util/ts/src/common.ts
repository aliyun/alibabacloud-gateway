import { Readable } from 'stream';

export async function readableStreamToBuffer(stream: Readable): Promise<Buffer> {
    const chunks: Buffer[] = [];
    return new Promise((resolve, reject) => {
        stream.on('data', (chunk: Buffer) => chunks.push(chunk));
        stream.on('error', reject);
        stream.on('end', () => {
            resolve(Buffer.concat(chunks));
        });
    });
}