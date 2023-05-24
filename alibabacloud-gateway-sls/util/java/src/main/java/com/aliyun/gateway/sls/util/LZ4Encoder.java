package com.aliyun.gateway.sls.util;

import net.jpountz.lz4.LZ4Compressor;
import net.jpountz.lz4.LZ4Factory;
import net.jpountz.lz4.LZ4FastDecompressor;

import com.aliyun.tea.*;

/**
 * LZ4Encoder is used to compress and decompress the log data
 *
 * @author sls_dev
 */
public class LZ4Encoder {

    public static byte[] compressToLhLz4Chunk(byte[] data) throws Exception {
        final int rawSize = data.length;
        LZ4Factory factory = LZ4Factory.fastestInstance();

        // compress data
        LZ4Compressor compressor = factory.fastCompressor();

        int maxCompressedLength = compressor.maxCompressedLength(rawSize);
        int encodingSize = 0;
        byte[] rawCompressed = new byte[maxCompressedLength];
        encodingSize = compressor.compress(data, 0, rawSize, rawCompressed, 0, maxCompressedLength);

        if (encodingSize <= 0) {
            throw new TeaException(TeaConverter.buildMap(new TeaPair("message", "Invalid encoding size")));
        }

        byte[] ret = new byte[encodingSize];
        System.arraycopy(rawCompressed, 0, ret, 0, encodingSize);

        return ret;
    }

    public static byte[] decompressFromLhLz4Chunk(byte[] compressedData, int rawSize) throws Exception {

        LZ4Factory factory = LZ4Factory.fastestInstance();

        // decompress data
        LZ4FastDecompressor decompressor = factory.fastDecompressor();

        byte[] restored = new byte[rawSize];
        decompressor.decompress(compressedData, 0, restored, 0, rawSize);

        return restored;
    }
}
