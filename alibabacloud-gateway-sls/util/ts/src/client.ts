// This file is auto-generated, don't edit it
/**
 * @remarks
 * Read data from a readable stream, and parse it by JSON format
 * 
 * @param stream - the readable stream
 * @returns the parsed result
 */
import { Readable } from 'stream';

import {gzipCompress, isCompressorAvailable, lz4Compress} from './compress';
import {lz4Decompress, gzipDecompress, DecompressFunc, isDecompressorAvailable} from './decompress';

export default class Client {

  static async readAndUncompressBlock(stream: Readable, compressType: string, bodyRawSize: string): Promise<Readable> {
    const rawSize = parseInt(bodyRawSize, 10);
    if (rawSize === 0) {
      return Readable.from([]);
    }

    let decompressFunc:DecompressFunc = undefined;
    if (compressType === 'lz4') {
      decompressFunc = lz4Decompress;
    } else if (compressType === 'gzip' || compressType === 'deflate') {
      decompressFunc = gzipDecompress;
    } else {
      throw new Error(`Unsupported decompress type: ${compressType}`)
    }

    let uncompressed: Buffer | undefined = undefined;
    try {
      uncompressed = await decompressFunc(stream, rawSize);
    }
    catch (error) {
      throw new Error(`Failed to decompress using compresstype ${compressType}, error: ${error.message}`);
    }

    if (uncompressed.length !== rawSize) {
      throw new Error(`unexpected uncompressed size: ${uncompressed.length}, expected: ${rawSize}, compressType: ${compressType}`);
    }
    return Readable.from([uncompressed]);
  }

  /**
   * @remarks
   * Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.
   * 
   * @param src - the data to be compressed
   * @param compressType - the compress type
   * @returns the compressed data
   * 
   * @throws error if the compress type is not supported or the compress failed
   */
  static async compress(src: Buffer, compressType: string): Promise<Buffer> {
    if (compressType === 'lz4') {
      return await lz4Compress(src)
    }
    if (compressType === 'gzip' || compressType === 'deflate') {
      return await gzipCompress(src)
    }
    throw new Error(`Unsupported compress type: ${compressType}`);
  }

  static async isCompressorAvailable(compressType: string): Promise<boolean> {
    return await isCompressorAvailable(compressType);
  }

  static async isDecompressorAvailable(compressType: string): Promise<boolean> {
    return await isDecompressorAvailable(compressType);
  }

  static async bytesLength(src: Buffer): Promise<number> {
    return src.length;
  }

}
