using System;
using System.Globalization;
using System.Reflection;
using System.Text.RegularExpressions;
using AlibabaCloud.GatewayMns;
using Xunit;

namespace tests
{
    public class UnitTest
    {
        private static readonly Regex Rfc1123Pattern = new Regex(
            @"^[A-Z][a-z]{2}, \d{2} [A-Z][a-z]{2} \d{4} \d{2}:\d{2}:\d{2} GMT$",
            RegexOptions.Compiled);

        private static readonly Regex Iso8601Pattern = new Regex(
            @"^\d{8}T\d{6}Z$",
            RegexOptions.Compiled);

        [Fact]
        public void Test_GetDateISO8601FromUTC()
        {
            Assert.Equal("20200206T073254Z", InvokeGetDateISO8601FromUTC("Thu, 06 Feb 2020 07:32:54 GMT"));
            Assert.Equal("20250603T112348Z", InvokeGetDateISO8601FromUTC("Tue, 03 Jun 2025 11:23:48 GMT"));
            Assert.Equal("20251203T112348Z", InvokeGetDateISO8601FromUTC("Wed, 03 Dec 2025 11:23:48 GMT"));
            Assert.Equal("20260101T000000Z", InvokeGetDateISO8601FromUTC("Thu, 01 Jan 2026 00:00:00 GMT"));
            Assert.Equal("20261231T235959Z", InvokeGetDateISO8601FromUTC("Thu, 31 Dec 2026 23:59:59 GMT"));
        }

        [Fact]
        public void Test_GetDateISO8601FromUTCRoundTrip()
        {
            string rfc = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            Assert.True(Rfc1123Pattern.IsMatch(rfc), "unexpected Date header format: " + rfc);

            DateTimeOffset parsed = DateTimeOffset.ParseExact(rfc, "r", CultureInfo.InvariantCulture);
            string expected = parsed.UtcDateTime.ToString("yyyyMMdd'T'HHmmss'Z'", CultureInfo.InvariantCulture);

            string actual = InvokeGetDateISO8601FromUTC(rfc);
            Assert.Equal(expected, actual);
            Assert.True(Iso8601Pattern.IsMatch(actual), "unexpected iso8601 format: " + actual);
        }

        private static string InvokeGetDateISO8601FromUTC(string utc)
        {
            MethodInfo method = typeof(Client).GetMethod(
                "GetDateISO8601FromUTC",
                BindingFlags.Instance | BindingFlags.NonPublic);
            Assert.NotNull(method);
            return (string)method.Invoke(new Client(), new object[] { utc });
        }
    }
}
