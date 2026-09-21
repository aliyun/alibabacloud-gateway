import Client from '../src/client';
import Util from '@alicloud/tea-util';
import assert from 'assert';
import 'mocha';

const MONTHS: { [key: string]: string } = {
  Jan: '01', Feb: '02', Mar: '03', Apr: '04', May: '05', Jun: '06',
  Jul: '07', Aug: '08', Sep: '09', Oct: '10', Nov: '11', Dec: '12',
};

const RFC1123_RE = /^([A-Z][a-z]{2}), (\d{2}) ([A-Z][a-z]{2}) (\d{4}) (\d{2}):(\d{2}):(\d{2}) GMT$/;

function parseRfc1123ToIso8601(rfc: string): string {
  const matched = rfc.match(RFC1123_RE);
  if (!matched) {
    throw new Error('invalid RFC1123: ' + rfc);
  }
  const month = MONTHS[matched[3]];
  if (!month) {
    throw new Error('invalid month: ' + matched[3]);
  }
  return `${matched[4]}${month}${matched[2]}T${matched[5]}${matched[6]}${matched[7]}Z`;
}

describe('Client date conversion', function () {
  it('getDateISO8601FromUTC should convert fixed RFC1123 strings', async function () {
    const client = new Client() as any;
    const cases: Array<[string, string]> = [
      ['Thu, 06 Feb 2020 07:32:54 GMT', '20200206T073254Z'],
      ['Tue, 03 Jun 2025 11:23:48 GMT', '20250603T112348Z'],
      ['Wed, 03 Dec 2025 11:23:48 GMT', '20251203T112348Z'],
      ['Thu, 01 Jan 2026 00:00:00 GMT', '20260101T000000Z'],
      ['Thu, 31 Dec 2026 23:59:59 GMT', '20261231T235959Z'],
    ];
    for (const [rfc, expected] of cases) {
      assert.strictEqual(await client.getDateISO8601FromUTC(rfc), expected);
    }
  });

  it('getDateISO8601FromUTC should convert live tea-util Date header', async function () {
    const client = new Client() as any;
    const rfc = Util.getDateUTCString();
    assert.ok(RFC1123_RE.test(rfc), 'unexpected Date header format: ' + rfc);
    const expected = parseRfc1123ToIso8601(rfc);
    const actual = await client.getDateISO8601FromUTC(rfc);
    assert.strictEqual(actual, expected);
    assert.ok(/^\d{8}T\d{6}Z$/.test(actual), 'unexpected iso8601 format: ' + actual);
  });
});
