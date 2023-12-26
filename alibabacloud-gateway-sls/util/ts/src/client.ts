// This file is auto-generated, don't edit it
/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
import { Readable } from 'stream';
import * as $tea from '@alicloud/tea-typescript';
import * as lz4 from 'lz4';


export default class Client {

  static async readAndUncompressBlock(stream: Readable, compressType: string, bodyRawSize: string): Promise<Readable> {
    if (compressType !== 'lz4') {
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
      const decompressedData = Buffer.allocUnsafe(uncompressedBodySize);
      const decompressedDataSize = lz4.decodeBlock(compressedData, decompressedData);
      if (decompressedDataSize !== uncompressedBodySize) {
        throw new Error('Decompressed data size does not match the expected uncompressed body size');
      }
      const decompressedStream = new Readable({
        read() {
          this.push(decompressedData);
          this.push(null);
        }
      });
      return decompressedStream;
    } catch (error) {
      throw new Error(`Failed to decompress LZ4 block: ${error.message}`);
    }
  }

}
