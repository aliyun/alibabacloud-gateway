// This file is auto-generated, don't edit it
/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
import { Readable } from 'stream';
const zlib = require('zlib');

export default class Client {

  static async readAndUncompressBlock(stream: Readable, compressType: string, bodyRawSize: string): Promise<Readable> {
    if (compressType !== 'gzip') {
      throw new Error(`Unsupported compress type: ${compressType}`);
    }
    const uncompressedBodySize = parseInt(bodyRawSize, 10);

    const compressedData = await new Promise<Buffer>((resolve, reject) => {
      const chunks: Buffer[] = [];
      stream.on('data', (chunk: Buffer) => chunks.push(chunk));
      stream.on('error', reject);
      stream.on('end', () => {
        resolve(Buffer.concat(chunks));
      });
    });
    try {
      const decompressedData = await new Promise<Buffer>((resolve, reject) => {
        zlib.inflate(compressedData, (err, result) => {
          if (err) reject(err);
          else resolve(result);
        });
      });
      if (decompressedData.length !== uncompressedBodySize) {
        throw new Error('Decompressed data size does not match the expected uncompressed body size');
      }
      return Readable.from([decompressedData]);
    } catch (error) {
      throw new Error(`Failed to decompress gzip block: ${error.message}`);
    }
  }

}
