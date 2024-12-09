import { Readable } from 'stream';
import { readableStreamToBuffer } from './common';
import lz4 from 'lz4';
import zlib from 'zlib';

const supportedDecompressor = ['lz4', 'gzip', 'deflate']

export async function isDecompressorAvailable(compressType: string): Promise<boolean> {
    return supportedDecompressor.includes(compressType);
}

export async function lz4Decompress(input: Readable, bodyRawSize: number): Promise<Buffer> {
    const output = Buffer.alloc(bodyRawSize);
    const n = lz4.decodeBlock(await readableStreamToBuffer(input), output);
    if (n !== bodyRawSize) {
        return Buffer.from(output.slice(0, n));
    }
    return output;
}

export async function gzipDecompress(input: Readable, bodyRawSize: number): Promise<Buffer> {
    const bytes = await readableStreamToBuffer(input);
    return new Promise((resolve, reject) => {
        zlib.inflate(bytes, (err, result) => {
            if (err) {
                reject(err);
            } else {
                resolve(result as Buffer);
            }
        });
    });
}

export type DecompressFunc = (input: Readable, bodyRawSize: number) => Promise<Buffer>;