/**
 * Read data from a readable stream, and parse it by JSON format
 * @param stream the readable stream
 * @return the parsed result
 */
// This file is auto-generated, don't edit it. Thanks.

using System;
using System.IO;
using System.Threading.Tasks;
using Ionic.Zlib;

namespace AlibabaCloud.GatewaySls_Util
{
    public class Client 
    {

        public static Stream ReadAndUncompressBlock(Stream stream, string compressType, string bodyRawSize)
        {
            if (!string.Equals(compressType, "gzip", StringComparison.OrdinalIgnoreCase))
            {
                throw new NotSupportedException(string.Format("{0} compression type is not implemented.", compressType));
            }
            long expectedSize;
            if (!long.TryParse(bodyRawSize, out expectedSize))
            {
                throw new ArgumentException("Invalid bodyRawSize value. It must be a valid numeric string.");
            }

            stream.Position = 0;
            MemoryStream decompressedStream = new MemoryStream();
            using (var decompressionStream = new ZlibStream(stream, Ionic.Zlib.CompressionMode.Decompress))
            {
                decompressionStream.CopyTo(decompressedStream);
            }
            decompressedStream.Position = 0;
            if (decompressedStream.Length != expectedSize)
            {
                throw new InvalidDataException(
                    string.Format("Decompressed data size is {0}, which does not match the expected size of {1}.",
                    decompressedStream.Length, expectedSize));
            }
            return decompressedStream;
        }

        public static async Task<Stream> ReadAndUncompressBlockAsync(Stream stream, string compressType, string bodyRawSize)
        {
            throw new NotImplementedException();
        }

    }
}
