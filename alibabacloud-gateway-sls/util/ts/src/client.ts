// This file is auto-generated, don't edit it
/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
import { Readable } from 'stream';
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
