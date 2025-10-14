import {compress as napiLz4Compress} from 'lz4-napi';
import zlib from 'zlib';

const supportedCompressor = ['lz4', 'gzip', 'deflate']

export async function lz4Compress(data: Buffer): Promise<Buffer> {
    return await napiLz4Compress(data);
}

export async function gzipCompress(data: Buffer): Promise<Buffer> {
    return new Promise((resolve, reject) => {
        zlib.deflate(data, (err, result) => {
            if (err) {
                reject(err);
            } else {
                resolve(result as Buffer);
            }
        });
    });
}

export async function isCompressorAvailable(compressType: string): Promise<boolean> {
    return supportedCompressor.includes(compressType);
}
