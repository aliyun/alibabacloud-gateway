// This file is auto-generated, don't edit it. Thanks.

using System;
using System.IO;
using System.Threading.Tasks;

namespace AlibabaCloud.GatewaySls_Util
{

    public class Client 
    {
        /// <term><b>Description:</b></term>
        /// <description>
        /// <para>Read data from a readable stream, and parse it by JSON format</para>
        /// </description>
        /// 
        /// <param name="stream">
        /// the readable stream
        /// </param>
        /// 
        /// <returns>
        /// the parsed result
        /// </returns>
        public static Stream ReadAndUncompressBlock(Stream stream, string compressType, string bodyRawSize)
        {
            long expectedSize;
            if (!long.TryParse(bodyRawSize, out expectedSize))
            {
                throw new ArgumentException("Invalid bodyRawSize value. It must be a valid numeric string.");
            }
            stream.Position = 0;

            Stream output;
            if (compressType == "deflate" || compressType == "gzip")
            {
                output = Decompressor.GzipDecompress(stream, expectedSize);
            }
            else if (compressType == "zstd")
            {
                output = Decompressor.ZstdDecompress(stream, expectedSize);
            }
            else
            {
                throw new NotSupportedException(string.Format("unsupported decompression type {0}.", compressType));
            }
            output.Position = 0;
            if (output.Length != expectedSize)
            {
                throw new InvalidDataException(
                    string.Format("unexpected uncompressed size: {0}, expected: {1}, compressType: {2}.", output.Length, expectedSize, compressType));
            }
            return output;
        }

        #pragma warning disable 1998
        public static async Task<Stream> ReadAndUncompressBlockAsync(Stream stream, string compressType, string bodyRawSize)
        {
            throw new NotImplementedException();
        }

        /// <term><b>Description:</b></term>
        /// <description>
        /// <para>Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.</para>
        /// </description>
        /// 
        /// <param name="src">
        /// the data to be compressed
        /// </param>
        /// <param name="compressType">
        /// the compress type
        /// </param>
        /// 
        /// <returns>
        /// the compressed data
        /// </returns>
        /// 
        /// <term><b>Exception:</b></term>
        /// error if the compress type is not supported or the compress failed
        public static byte[] Compress(byte[] src, string compressType)
        {
            if (compressType == "deflate" || compressType == "gzip")
            {
                return Compressor.GzipCompress(src);
            }
            else if (compressType == "zstd")
            {
                return Compressor.ZstdCompress(src);
            }
            else
            {
                throw new NotSupportedException(string.Format("unsupported compression type {0}.", compressType));
            }
        }

        /// <term><b>Description:</b></term>
        /// <description>
        /// <para>Compress data by specified compress type, use isCompressorAvailable to check if the compress type is supported.</para>
        /// </description>
        /// 
        /// <param name="src">
        /// the data to be compressed
        /// </param>
        /// <param name="compressType">
        /// the compress type
        /// </param>
        /// 
        /// <returns>
        /// the compressed data
        /// </returns>
        /// 
        /// <term><b>Exception:</b></term>
        /// error if the compress type is not supported or the compress failed
        public static async Task<byte[]> CompressAsync(byte[] src, string compressType)
        {
            throw new NotImplementedException();
        }

        public static bool IsCompressorAvailable(string compressType)
        {
            return Compressor.IsCompressorAvailable(compressType);
        }

        public static async Task<bool> IsCompressorAvailableAsync(string compressType)
        {
            throw new NotImplementedException();
        }

        public static bool IsDecompressorAvailable(string compressType)
        {
            return Decompressor.IsDecompressorAvailable(compressType);
        }

        public static async Task<bool> IsDecompressorAvailableAsync(string compressType)
        {
            throw new NotImplementedException();
        }

        public static long BytesLength(byte[] src)
        {
            return src.Length;
        }

        public static async Task<long> BytesLengthAsync(byte[] src)
        {
            return await Task.FromResult<long>(src.Length);
        }
        #pragma warning restore 1998
    }
}
