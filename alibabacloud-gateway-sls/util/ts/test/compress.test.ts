import Client from '../src/client';
import { readableStreamToBuffer } from '../src/common';
import { Readable } from 'stream';
import assert from 'assert';


async function testOnData(data: string) {
    const originalData = Buffer.from('Hello, world!');
    for (const compressType of ['lz4', 'gzip', 'deflate']) {

        const compressedData = await Client.compress(originalData, compressType);
        const readableStream = Readable.from([compressedData]);

        const uncompressResult = await Client.readAndUncompressBlock(readableStream, compressType, String(originalData.length));
        const output = await readableStreamToBuffer(uncompressResult);

        assert.strictEqual(originalData.length, output.length);
        assert.deepStrictEqual(originalData, output, 'The uncompress data should match the original data');
    }
}

function generateRamdonString(length: number): string {
    let result = '';
    for (let i = 0; i < length; i++) {
        result += String.fromCharCode(Math.floor(Math.random() * 256));
    }
    return result;
}

async function runTest() {
    await testOnData('Hello, world!');
    await testOnData('');
    await testOnData('Hello, world!你好');
    for (let i = 0; i < 100; i++) {
        await testOnData(generateRamdonString(Math.floor(Math.random() * 1000)));
    }
    console.log('All tests passed!');
}

runTest();