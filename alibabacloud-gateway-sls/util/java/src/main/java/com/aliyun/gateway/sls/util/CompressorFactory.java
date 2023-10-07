package com.aliyun.gateway.sls.util;

import com.aliyun.tea.TeaConverter;
import com.aliyun.tea.TeaException;
import com.aliyun.tea.TeaPair;
import net.jpountz.lz4.LZ4Compressor;
import net.jpountz.lz4.LZ4Exception;
import net.jpountz.lz4.LZ4Factory;

import java.io.ByteArrayOutputStream;
import java.util.zip.Deflater;

interface Compressor {
    byte[] compress(byte[] data) throws Exception;
}

class Lz4Compressor implements Compressor {
    @Override
    public byte[] compress(byte[] data) {
        final int rawSize = data.length;
        LZ4Factory factory = LZ4Factory.fastestInstance();

        // compress data
        LZ4Compressor compressor = factory.fastCompressor();

        int maxCompressedLength = compressor.maxCompressedLength(rawSize);
        int encodingSize = 0;
        byte[] rawCompressed = new byte[maxCompressedLength];
        try {
            encodingSize = compressor.compress(data, 0, rawSize, rawCompressed, 0, maxCompressedLength);
        } catch (LZ4Exception e) {
            throw new TeaException(TeaConverter.buildMap(new TeaPair("CompressException", e.getMessage())));
        }

        if (encodingSize <= 0) {
            throw new TeaException(TeaConverter.buildMap(new TeaPair("CompressException", "Invalid encoding size")));
        }

        byte[] ret = new byte[encodingSize];
        System.arraycopy(rawCompressed, 0, ret, 0, encodingSize);

        return ret;
    }
}

class GzipCompressor implements Compressor {
    @Override
    public byte[] compress(byte[] data) {
        ByteArrayOutputStream out = new ByteArrayOutputStream(data.length);
        Deflater compressor = new Deflater();
        try {
            compressor.setInput(data);
            compressor.finish();
            byte[] buf = new byte[10240];
            while (!compressor.finished()) {
                int count = compressor.deflate(buf);
                out.write(buf, 0, count);
            }
            return out.toByteArray();
        } finally {
            compressor.end();
        }
    }
}

public class CompressorFactory {
    public static Compressor getCompressor(String compressType) {
        if ("lz4".equals(compressType)) {
            return new Lz4Compressor();
        }

        throw new IllegalArgumentException("Invalid compressType: " + compressType);
    }
}
