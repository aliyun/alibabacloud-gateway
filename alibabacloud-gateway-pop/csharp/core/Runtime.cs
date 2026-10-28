using System;
using System.Collections;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using Darabonba;
using Darabonba.Utils;
using Org.BouncyCastle.Crypto.Digests;
using Org.BouncyCastle.Crypto.Macs;
using Org.BouncyCastle.Crypto.Parameters;
using Tea;

namespace AlibabaCloud.GatewayPop
{
    /// <summary>
    /// Gateway-private helpers. Not part of the supported public API.
    /// </summary>
    internal static class Runtime
    {
        internal static bool IsUnset(object obj)
        {
            return Extensions.IsNull(obj);
        }

        internal static string GetNonce()
        {
            return Guid.NewGuid().ToString();
        }

        internal static bool Empty(string val)
        {
            return string.IsNullOrEmpty(val);
        }

        internal static string DefaultString(string real, string defaultStr)
        {
            if (string.IsNullOrWhiteSpace(real))
            {
                return defaultStr;
            }
            return real;
        }

        internal static byte[] ToBytes(string val)
        {
            return BytesUtils.From(val, "utf-8");
        }

        internal static byte[] ReadAsBytes(System.IO.Stream stream)
        {
            return StreamUtils.ReadAsBytes(stream);
        }

        internal static string ReadAsString(System.IO.Stream stream)
        {
            return StreamUtils.ReadAsString(stream);
        }

        internal static object ReadAsJSON(System.IO.Stream stream)
        {
            return StreamUtils.ReadAsJSON(stream);
        }

        internal static byte[] AssertAsBytes(object obj)
        {
            if (obj is byte[])
            {
                return (byte[])obj;
            }
            throw new ArgumentException("The value is not byte[]");
        }

        internal static Dictionary<string, object> AssertAsMap(object value)
        {
            if (value != null && value is IDictionary)
            {
                IDictionary dic = (IDictionary)value;
                Dictionary<string, object> dicObj = new Dictionary<string, object>();
                foreach (object key in dic.Keys)
                {
                    dicObj.Add((string)key, dic[key]);
                }
                return dicObj;
            }
            if (value is Newtonsoft.Json.Linq.JObject)
            {
                return (Dictionary<string, object>)JSONUtils.Deserialize(value);
            }
            throw new ArgumentException("The value is not Dictionary");
        }

        internal static string ToJSONString(object value)
        {
            return JSONUtils.SerializeObject(value);
        }

        internal static bool EqualString(string val1, string val2)
        {
            return val1 == val2;
        }

        internal static bool EqualNumber(int? val1, int val2)
        {
            return val1 == val2;
        }

        internal static bool Is4xx(int? code)
        {
            return code >= 400 && code < 500;
        }

        internal static bool Is5xx(int? code)
        {
            return code >= 500 && code < 600;
        }

        internal static string GetTimestamp()
        {
            return DateTime.UtcNow.ToString("yyyy-MM-dd'T'HH:mm:ss'Z'", CultureInfo.InvariantCulture);
        }

        internal static string GetEndpointRules(string product, string regionId, string endpointType, string network, string suffix)
        {
            string result;
            product = product == null ? "" : product;
            if (endpointType == "regional")
            {
                if (string.IsNullOrEmpty(regionId))
                {
                    throw new ArgumentException("RegionId is empty, please set a valid RegionId", "RegionId");
                }
                result = "<product><suffix><network>.<region_id>.aliyuncs.com".Replace("<region_id>", regionId);
            }
            else
            {
                result = "<product><suffix><network>.aliyuncs.com";
            }
            result = result.Replace("<product>", product.ToLower());
            if (string.IsNullOrEmpty(network) || network == "public")
            {
                result = result.Replace("<network>", "");
            }
            else
            {
                result = result.Replace("<network>", "-" + network);
            }
            if (string.IsNullOrEmpty(suffix))
            {
                result = result.Replace("<suffix>", "");
            }
            else
            {
                result = result.Replace("<suffix>", "-" + suffix);
            }
            return result;
        }

        internal static string ToForm(IDictionary filter)
        {
            if (filter == null)
            {
                return string.Empty;
            }
            Dictionary<string, object> dict = filter.Keys.Cast<string>().ToDictionary(key => key, key => filter[key]);
            Dictionary<string, string> outDict = new Dictionary<string, string>();
            TileDict(outDict, dict);
            List<string> listStr = new List<string>();
            foreach (KeyValuePair<string, string> keypair in outDict)
            {
                if (string.IsNullOrWhiteSpace(keypair.Value))
                {
                    continue;
                }
                listStr.Add(PercentEncode(keypair.Key) + "=" + PercentEncode(keypair.Value));
            }
            return string.Join("&", listStr);
        }

        internal static string PercentEncode(string value)
        {
            return URL.PercentEncode(value);
        }

        internal static string HexEncode(byte[] raw)
        {
            return BytesUtils.ToHex(raw);
        }

        internal static byte[] Hash(byte[] raw, string signatureAlgorithm)
        {
            if (signatureAlgorithm.Contains("HMAC-SHA256") || signatureAlgorithm.Contains("RSA-SHA256"))
            {
                using (SHA256 sha256 = SHA256.Create())
                {
                    return sha256.ComputeHash(raw);
                }
            }
            return null;
        }

        internal static byte[] HmacSHA256Sign(string stringToSign, string secret)
        {
            return HmacSHA256SignByBytes(stringToSign, Encoding.UTF8.GetBytes(secret));
        }

        internal static byte[] HmacSHA256SignByBytes(string stringToSign, byte[] secret)
        {
            byte[] signData;
            using (KeyedHashAlgorithm algorithm = CryptoConfig.CreateFromName("HMACSHA256") as KeyedHashAlgorithm)
            {
                algorithm.Key = secret;
                signData = algorithm.ComputeHash(Encoding.UTF8.GetBytes(stringToSign.ToCharArray()));
            }
            return signData;
        }

        internal static byte[] HmacSM3Sign(string stringToSign, string secret)
        {
            return HmacSM3SignByBytes(stringToSign, Encoding.Default.GetBytes(secret));
        }

        internal static byte[] HmacSM3SignByBytes(string stringToSign, byte[] secret)
        {
            byte[] msg = Encoding.Default.GetBytes(stringToSign);
            KeyParameter keyParameter = new KeyParameter(secret);
            SM3Digest sm3 = new SM3Digest();
            HMac mac = new HMac(sm3);
            mac.Init(keyParameter);
            mac.BlockUpdate(msg, 0, msg.Length);
            byte[] signData = new byte[mac.GetMacSize()];
            mac.DoFinal(signData, 0);
            return signData;
        }

        internal static List<string> Split(string raw, string sep, int? limit = null)
        {
            if (raw == null)
            {
                return null;
            }
            if (limit == null || limit == -1)
            {
                return new List<string>(raw.Split(new string[] { sep }, StringSplitOptions.None));
            }
            return new List<string>(raw.Split(new string[] { sep }, limit.Value, StringSplitOptions.None));
        }

        internal static string Replace(string raw, string oldStr, string newStr, int? count = null)
        {
            if (raw == null)
            {
                return null;
            }
            if (count == null || count == -1)
            {
                return raw.Replace(oldStr, newStr);
            }
            string[] rawSplit = raw.Split(new string[] { oldStr }, count.Value + 1, StringSplitOptions.None);
            return string.Join(newStr, rawSplit);
        }

        internal static bool Contains(string s, string substr)
        {
            if (s == null)
            {
                return false;
            }
            return s.Contains(substr);
        }

        internal static bool HasPrefix(string s, string prefix)
        {
            if (s == null)
            {
                return false;
            }
            return s.StartsWith(prefix);
        }

        internal static string ToLower(string s)
        {
            if (s == null)
            {
                return null;
            }
            return s.ToLower();
        }

        internal static string SubString(string s, int? start, int? end)
        {
            if (s == null)
            {
                return null;
            }
            if (end.Value >= s.Length)
            {
                return s.Substring(start.Value);
            }
            int length = end.Value - start.Value;
            if (length < 0)
            {
                return string.Empty;
            }
            return s.Substring(start.Value, length);
        }

        internal static bool Equals(string expect, string actual)
        {
            if (expect == null)
            {
                return false;
            }
            return expect.Equals(actual);
        }

        internal static string Trim(string raw)
        {
            return raw.Trim();
        }

        internal static List<string> KeySet(Dictionary<string, string> raw)
        {
            if (raw == null)
            {
                throw new ArgumentException("not a valid value for parameter");
            }
            return new List<string>(raw.Keys);
        }

        internal static string Join(List<string> raw, string sep)
        {
            return string.Join(sep, raw);
        }

        internal static int Size(List<string> raw)
        {
            return raw.Count;
        }

        internal static List<string> AscSort(List<string> raw)
        {
            return ListUtils.Sort(raw, "asc");
        }

        internal static bool ListContains(List<string> raw, string str)
        {
            return raw.Contains(str);
        }

        internal static void Append<T>(List<T> array, T item)
        {
            if (null != array)
            {
                array.Add(item);
            }
        }

        private static void TileDict(Dictionary<string, string> dicOut, object obj, string parentKey = "")
        {
            if (obj == null)
            {
                return;
            }
            if (typeof(IDictionary).IsAssignableFrom(obj.GetType()))
            {
                IDictionary source = (IDictionary)obj;
                foreach (object key in source.Keys)
                {
                    string keyName = parentKey + "." + (string)key;
                    if (source[key] == null)
                    {
                        continue;
                    }
                    TileDict(dicOut, source[key], keyName);
                }
            }
            else if (typeof(TeaModel).IsAssignableFrom(obj.GetType()))
            {
                Dictionary<string, object> dicIn = (Dictionary<string, object>)((TeaModel)obj).ToMap();
                foreach (KeyValuePair<string, object> keypair in dicIn)
                {
                    string keyName = parentKey + "." + keypair.Key;
                    if (keypair.Value == null)
                    {
                        continue;
                    }
                    TileDict(dicOut, keypair.Value, keyName);
                }
            }
            else if (typeof(IList).IsAssignableFrom(obj.GetType()) && !typeof(Array).IsAssignableFrom(obj.GetType()))
            {
                int index = 1;
                foreach (object temp in (IList)obj)
                {
                    TileDict(dicOut, temp, parentKey + "." + index.ToSafeString());
                    index++;
                }
            }
            else
            {
                if (obj.GetType() == typeof(byte[]))
                {
                    dicOut.Add(parentKey.TrimStart('.'), Encoding.UTF8.GetString((byte[])obj));
                }
                else
                {
                    dicOut.Add(parentKey.TrimStart('.'), obj.ToSafeString(""));
                }
            }
        }
    }
}
