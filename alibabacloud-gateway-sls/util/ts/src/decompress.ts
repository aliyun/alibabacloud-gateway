import { Readable } from 'stream';
import { readableStreamToBuffer } from './common';
import {uncompress as napiLz4Uncompress} from 'lz4-napi';
import zlib from 'zlib';

const supportedDecompressor = ['lz4', 'gzip', 'deflate']

export async function isDecompressorAvailable(compressType: string): Promise<boolean> {
    return supportedDecompressor.includes(compressType);
}

export async function lz4Decompress(input: Readable, bodyRawSize: number): Promise<Buffer> {
    const napiOutput = await napiLz4Uncompress(await readableStreamToBuffer(input));
    if (napiOutput.length !== bodyRawSize) {
        return Buffer.from(napiOutput.slice(0, bodyRawSize));
    }
    return napiOutput;
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