import lz4 from 'lz4';
import zlib from 'zlib';

const supportedCompressor = ['lz4', 'gzip', 'deflate']

export async function lz4Compress(data: Buffer): Promise<Buffer> {
    const output = Buffer.alloc(lz4.encodeBound(data.length));
    const n = lz4.encodeBlock(data, output);
    return Buffer.from(output.slice(0, n));
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
