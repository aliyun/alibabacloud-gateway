import { Readable } from 'stream';
import { readableStreamToBuffer, readableStreamWithPrependBuffer } from './common';
import {uncompress as napiLz4Uncompress} from 'lz4-napi';
import zlib from 'zlib';
import fs from 'fs';
const supportedDecompressor = ['lz4', 'gzip', 'deflate']

export async function isDecompressorAvailable(compressType: string): Promise<boolean> {
    return supportedDecompressor.includes(compressType);
}

export async function lz4Decompress(input: Readable, bodyRawSize: number): Promise<Buffer> {
    // prepend 4 bytes size to the buffer
    const sizeBuffer = Buffer.alloc(4);
    sizeBuffer.writeUInt32LE(bodyRawSize, 0);
    let compressedBuffer = await readableStreamWithPrependBuffer(input, sizeBuffer);
    const napiOutput = await napiLz4Uncompress(compressedBuffer);
    if (napiOutput.length !== bodyRawSize) {
        throw new Error(`unexpected uncompressed size: ${napiOutput.length}, expected: ${bodyRawSize}`);
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