// This file is auto-generated, don't edit it. Thanks.

using System.IO;
using Ionic.Zlib;

namespace AlibabaCloud.GatewaySls_Util
{
    public class Compressor
    {
        public static bool IsCompressorAvailable(string compressType)
        {
            return compressType == "deflate" || compressType == "gzip";
        }

        public static byte[] GzipCompress(byte[] src)
        {
            // ZlibStream.CompressBuffer(src) uses CompressionLevel.Best
            // here we use CompressionLevel.Default
            using (MemoryStream memoryStream = new MemoryStream())
            {
                using (Stream compressor = new ZlibStream(memoryStream, CompressionMode.Compress, CompressionLevel.Default))
                {
                    compressor.Write(src, 0, src.Length);
                }
                return memoryStream.ToArray();
            }
        }
    }
}
