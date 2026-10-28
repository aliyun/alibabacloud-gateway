/**
 * Gateway-private signing helpers. Not part of the supported public API.
 * @internal
 */
import crypto from 'crypto';
import * as $dara from '@darabonba/typescript';

/** @internal */
export function hexEncode(raw: Buffer): string {
  if (raw === null) {
    return null;
  }
  return raw.toString('hex');
}

/** @internal */
export function hash(raw: Buffer, signatureAlgorithm: string): Buffer {
  if (signatureAlgorithm == null) {
    return null;
  }
  if (signatureAlgorithm.includes('HMAC-SHA256') || signatureAlgorithm.includes('RSA-SHA256')) {
    const obj = crypto.createHash('sha256');
    obj.update(raw);
    return obj.digest();
  } else if (signatureAlgorithm.includes('HMAC-SM3')) {
    const obj = crypto.createHash('sm3');
    obj.update(raw);
    return obj.digest();
  }
  return null;
}

/** @internal */
export function percentEncode(raw: string): string {
  return $dara.URL.percentEncode(raw);
}

/** @internal */
export function hmacSha256(stringToSign: string, secret: string): Buffer {
  const obj = crypto.createHmac('sha256', secret);
  obj.update(stringToSign);
  return obj.digest();
}

/** @internal */
export function hmacSha256ByBytes(stringToSign: string, secret: Buffer): Buffer {
  const obj = crypto.createHmac('sha256', secret);
  obj.update(stringToSign);
  return obj.digest();
}

/** @internal */
export function hmacSm3(stringToSign: string, secret: string): Buffer {
  const obj = crypto.createHmac('sm3', secret);
  obj.update(stringToSign);
  return obj.digest();
}

/** @internal */
export function hmacSm3ByBytes(stringToSign: string, secret: Buffer): Buffer {
  const obj = crypto.createHmac('sm3', secret);
  obj.update(stringToSign);
  return obj.digest();
}

/** @internal */
export function split(raw: string, sep: string, limit: number): string[] {
  if (null === limit || typeof (limit) === 'undefined' || limit < 1) {
    return raw.split(sep);
  }
  const result = raw.split(sep, limit);
  if (result.length >= limit) {
    result[limit - 1] = [result[limit - 1], ...raw.split(sep).splice(limit)].join(sep);
  }
  return result;
}

/** @internal */
export function replace(raw: string, oldStr: string, newStr: string, count: number = null): string {
  if (null === count || typeof (count) === 'undefined' || count < 0) {
    return raw.split(oldStr).join(newStr);
  }
  const tmp = raw.split(oldStr);
  if (count >= tmp.length - 1) {
    return raw.split(oldStr).join(newStr);
  }
  const left = tmp.slice(0, count + 1).join(newStr);
  const right = tmp.slice(count + 1).join(oldStr);
  return left + oldStr + right;
}

/** @internal */
export function contains(s: string, substr: string): boolean {
  return s.indexOf(substr) !== -1;
}

/** @internal */
export function hasPrefix(s: string, prefix: string): boolean {
  return s.startsWith(prefix);
}

/** @internal */
export function toLower(s: string): string {
  return s.toLowerCase();
}

/** @internal */
export function subString(s: string, start: number, end: number): string {
  return s.slice(start, end);
}

/** @internal */
export function equals(expect: string, actual: string): boolean {
  return expect === actual;
}

/** @internal */
export function trim(str: string): string {
  return str.trim();
}

/** @internal */
export function keySet(raw: { [key: string]: any }): string[] {
  if (null === raw || typeof (raw) === 'undefined') {
    throw new Error('not a valid value for parameter');
  }
  return Object.keys(raw);
}

/** @internal */
export function join(raw: string[], sep: string): string {
  if (null === raw || typeof (raw) === 'undefined') {
    throw new Error('not a valid value for parameter');
  }
  return raw.join(sep);
}

/** @internal */
export function size(raw: string[]): number {
  if (null === raw || typeof (raw) === 'undefined') {
    throw new Error('not a valid value for parameter');
  }
  return raw.length;
}

/** @internal */
export function ascSort(raw: string[]): string[] {
  if (null === raw || typeof (raw) === 'undefined') {
    throw new Error('not a valid value for parameter');
  }
  return raw.sort();
}

/** @internal */
export function append(array: any, item: any): void {
  if (null !== array && Array.isArray(array)) {
    array.push(item);
  }
}

/** @internal */
export function listContains(raw: string[], str: string): boolean {
  if (null === raw || typeof (raw) === 'undefined') {
    throw new Error('not a valid value for parameter');
  }
  return raw.includes(str);
}
