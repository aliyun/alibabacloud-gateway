package com.aliyun.gateway.sls.util;

import com.github.luben.zstd.Zstd;
import net.jpountz.lz4.LZ4Compressor;
import net.jpountz.lz4.LZ4Factory;

import java.io.ByteArrayOutputStream;
import java.util.Arrays;
import java.util.List;

import java.util.zip.Deflater;

interface Compressor {
    byte[] compress(byte[] data) throws Exception;
}

class CompressorFactory {
    public static final List<String> supportedCompressTypes = Arrays.asList("lz4", "zstd", "gzip", "deflate");

    public static boolean isCompressorAvailable(String compressType) {
        return supportedCompressTypes.contains(compressType);
    }

    public static Compressor getCompressor(String compressType) {
        if ("lz4".equals(compressType)) {
            return new Lz4Compressor();
        }
        if ("gzip".equals(compressType) || "deflate".equals(compressType)) {
            return new GzipCompressor();
        }
        if ("zstd".equals(compressType)) {
            return new ZstdCompressor();
        }

        throw new IllegalArgumentException("Invalid compressType: " + compressType);
    }
}

class Lz4Compressor implements Compressor {
    @Override
    public byte[] compress(byte[] data) throws Exception {
        LZ4Factory factory = LZ4Factory.fastestInstance();
        LZ4Compressor compressor = factory.fastCompressor();
        int maxCompressedLength = compressor.maxCompressedLength(data.length);
        byte[] compressed = new byte[maxCompressedLength];
        int compressedLength = compressor.compress(data, 0, data.length, compressed, 0, maxCompressedLength);
        // todo: eliminate bytes copy
        byte[] finalCompressed = new byte[compressedLength];
        System.arraycopy(compressed, 0, finalCompressed, 0, compressedLength);

        return finalCompressed;
    }
}

class GzipCompressor implements Compressor {
    @Override
    public byte[] compress(byte[] data) throws Exception {
        Deflater compressor = new Deflater();
        compressor.setInput(data);
        compressor.finish();

        ByteArrayOutputStream bos = new ByteArrayOutputStream(data.length);
        byte[] buf = new byte[1024];
        while (!compressor.finished()) {
            int count = compressor.deflate(buf);
            bos.write(buf, 0, count);
        }
        compressor.end();

        return bos.toByteArray();
    }
}

class ZstdCompressor implements Compressor {
    @Override
    public byte[] compress(byte[] data) throws Exception {
        return Zstd.compress(data);
    }
}