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
const zlib = require('zlib');
import * as po_protobuf from 'pomelo-protobuf';
import * as $tea from '@alicloud/tea-typescript';

// 定义 proto_json 对象类型
interface ProtoJson {
  [key: string]: any;
}

const proto_json: ProtoJson = {
  "message Log": {
    "required uInt32 time": 1,
    "message Content": {
      "required string key": 1,
      "required string value": 2,
    },
    "repeated Content contents": 2,
  },
  "message LogTag": {
    "required string Key": 1,
    "required string Value": 2,
  },
  "LogGroup": {
    "repeated Log logs": 1,
    "optional string category": 2,
    "optional string topic": 3,
    "optional string source": 4,
    "optional string MachineUUID": 5,
    "repeated LogTag LogTags": 6,
  },
  "LogStore": {
    "repeated Log logs": 1,
    "optional string category": 2,
    "optional string topic": 3,
    "optional string source": 4,
  },
  "LogGroupList": {
    "message LogGroup": {
      "repeated Log logs": 1,
      "optional string category": 2,
      "optional string topic": 3,
      "optional string source": 4,
      "optional string MachineUUID": 5,
      "repeated LogTag LogTags": 6,
    },
    "repeated LogGroup logGroupList": 1,
  },
};

const protos = po_protobuf.parse(proto_json);

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

  static serializeToPbBytes(request: any): Buffer {
    po_protobuf.init({
      encoderProtos: protos,
      decoderProtos: protos,
    });
    const logGroup = JSON.parse(request.body);

    if (!logGroup.logs) {
      throw $tea.newError({
        code: 'ContentError',
        message: 'Logitems is empty.',
      });
    }

    if (logGroup.logs.length > 4096) {
      throw $tea.newError({
        code: 'ContentError',
        message: 'Logitems length exceed 4096.',
      });
    }

    const pb = po_protobuf.encode('LogGroup', logGroup);

    if (pb.length > 3 * 1024 * 1024) {
      throw $tea.newError({
        code: 'ContentError',
        message: 'Logitems size exceed 5MB',
      });
    }

    return pb;
  }

  static deserializeFromPbBytes(data: Buffer, statusCode: number, headers: { [key: string]: string }): any {
    try {
      po_protobuf.init({
        encoderProtos: protos,
        decoderProtos: protos,
      });
  
      const pb = po_protobuf.decode("LogGroupList", data);
      return pb;
    } catch (error) {
      var requestId = headers['x-log-requestid'];
      throw $tea.newError({
        code: 'ContentError',
        message: 'Protobuf parse failed.',
        data: {
          requestId: requestId,
          statusCode: statusCode,
        },
      });
    }
  }

}
