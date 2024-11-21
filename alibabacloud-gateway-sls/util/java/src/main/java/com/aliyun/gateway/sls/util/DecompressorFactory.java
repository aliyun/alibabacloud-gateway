package com.aliyun.gateway.sls.util;

import com.github.luben.zstd.Zstd;
import net.jpountz.lz4.LZ4Factory;
import net.jpountz.lz4.LZ4FastDecompressor;

import java.io.ByteArrayOutputStream;
import java.util.Arrays;
import java.util.List;
import java.util.zip.DataFormatException;
import java.util.zip.Inflater;

interface Decompressor {
    byte[]  decompress(byte[] data, int rawSize) throws Exception;
}

class DecompressorFactory {
    public static final List<String> supportedCompressTypes = Arrays.asList("lz4", "gzip", "zstd", "deflate");

    public static boolean isDecompressorAvailable(String compressType) {
        return supportedCompressTypes.contains(compressType);
    }

    public static Decompressor getDecompressor(String compressType) {
        if ("lz4".equals(compressType)) {
            return new Lz4Decompressor();
        }
        if ("gzip".equals(compressType) || "deflate".equals(compressType)) {
            return new GzipDecompressor();
        }
        if ("zstd".equals(compressType)) {
            return new ZstdDecompressor();
        }

        throw new IllegalArgumentException("Invalid compressType: " + compressType);
    }
    
}

class Lz4Decompressor implements Decompressor {
    @Override
    public byte[] decompress(byte[] data, int rawSize) {
        LZ4Factory factory = LZ4Factory.fastestInstance();

        // decompress data
        LZ4FastDecompressor decompressor = factory.fastDecompressor();

        byte[] restored = new byte[rawSize];
        decompressor.decompress(data, 0, restored, 0, rawSize);

        return restored;
    }
}

class GzipDecompressor implements Decompressor {
    @Override
    public byte[] decompress(byte[] data, int rawSize) throws DataFormatException {
        ByteArrayOutputStream bos = new ByteArrayOutputStream(data.length);
        Inflater decompressor = new Inflater();
        try {
            decompressor.setInput(data);
            final byte[] buf = new byte[1024];
            while (!decompressor.finished()) {
                int count = decompressor.inflate(buf);
                bos.write(buf, 0, count);
            }
            return bos.toByteArray();
        } finally {
            decompressor.end();
        }
    }
}

class ZstdDecompressor implements Decompressor {
    @Override
    public byte[] decompress(byte[] data, int rawSize) {
        return Zstd.decompress(data, rawSize);
    }
}