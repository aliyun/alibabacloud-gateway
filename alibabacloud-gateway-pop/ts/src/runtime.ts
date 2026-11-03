/**
 * Gateway-private helpers. Not part of the supported public API.
 * @internal
 */
import crypto from 'crypto';
import os from 'os';
import querystring from 'querystring';
import { Readable } from 'stream';
import * as $dara from '@darabonba/typescript';

let nonceCounter = 0;
let nonceLast: number;

/** @internal */
export function isUnset(value: any): boolean {
  return $dara.isNull(value);
}

/** @internal */
export function empty(val: string): boolean {
  return !val;
}

/** @internal */
export function equalString(val1: string, val2: string): boolean {
  return val1 === val2;
}

/** @internal */
export function equalNumber(val1: number, val2: number): boolean {
  return val1 === val2;
}

/** @internal */
export function defaultString(real: string, defaultValue: string): string {
  return real || defaultValue;
}

/** @internal */
export function toBytes(val: string): Buffer {
  return Buffer.from(val);
}

/** @internal */
export function toJSONString(val: any): string {
  if (typeof val === 'string') {
    return val;
  }
  return JSON.stringify(val);
}

/** @internal */
export async function readAsBytes(stream: Readable): Promise<Buffer> {
  return $dara.Stream.readAsBytes(stream);
}

/** @internal */
export async function readAsString(stream: Readable): Promise<string> {
  return $dara.Stream.readAsString(stream);
}

/** @internal */
export async function readAsJSON(stream: Readable): Promise<any> {
  return $dara.Stream.readAsJSON(stream);
}

/** @internal */
export function assertAsMap(value: any): { [key: string]: any } {
  if (value && typeof value === 'object' && !Array.isArray(value)) {
    return value as { [key: string]: any };
  }
  throw new Error('The value is not a object');
}

/** @internal */
export function assertAsBytes(value: any): Buffer {
  if (Buffer.isBuffer(value)) {
    return value;
  }
  throw new Error('The value is not bytes');
}

/** @internal */
export function is4xx(code: number): boolean {
  return code >= 400 && code < 500;
}

/** @internal */
export function is5xx(code: number): boolean {
  return code >= 500 && code < 600;
}

/** @internal */
export function getNonce(): string {
  const machine = os.hostname();
  const pid = process.pid;
  const val = Math.floor(Math.random() * 1000000000000);
  if (val === nonceLast) {
    nonceCounter++;
  } else {
    nonceCounter = 0;
  }
  nonceLast = val;
  const uid = `${machine}${pid}${val}${nonceCounter}`;
  return crypto.createHash('md5').update(uid, 'utf8').digest('hex');
}

function pad2(num: number): string {
  return num < 10 ? '0' + num : '' + num;
}

/** @internal */
export function getTimestamp(): string {
  const date = new Date();
  const YYYY = date.getUTCFullYear();
  const MM = pad2(date.getUTCMonth() + 1);
  const DD = pad2(date.getUTCDate());
  const HH = pad2(date.getUTCHours());
  const mm = pad2(date.getUTCMinutes());
  const ss = pad2(date.getUTCSeconds());
  return `${YYYY}-${MM}-${DD}T${HH}:${mm}:${ss}Z`;
}

/** @internal */
export function getEndpointRules(product: string, regionId: string, endpointType: string, network: string, suffix: string): string {
  let networkPart = network;
  if (networkPart && networkPart.length && networkPart != 'public') {
    networkPart = '-' + networkPart;
  } else {
    networkPart = '';
  }
  let suffixPart = suffix || '';
  if (suffixPart.length) {
    suffixPart = '-' + suffixPart;
  }
  if (endpointType == 'regional') {
    if (!regionId || !regionId.length) {
      throw new Error('RegionId is empty, please set a valid RegionId');
    }
    return `${product}${suffixPart}${networkPart}.${regionId}.aliyuncs.com`;
  }
  return `${product}${suffixPart}${networkPart}.aliyuncs.com`;
}

function replaceRepeatList(target: { [key: string]: string }, repeat: any[], prefix: string) {
  if (prefix) {
    prefix = prefix + '.';
  }
  for (let i = 0; i < repeat.length; i++) {
    const item = repeat[i];
    const key = prefix + (i + 1);
    if (typeof item === 'undefined' || item == null) {
      continue;
    }
    if (Array.isArray(item)) {
      replaceRepeatList(target, item, key);
    } else if (item instanceof Object) {
      flatMap(target, item, key);
    } else {
      target[key] = item.toString();
    }
  }
}

function toMap(input: any): any {
  if (input == null || typeof input !== 'object') {
    return input;
  }
  if (typeof input.toMap === 'function') {
    return input.toMap();
  }
  if (Array.isArray(input)) {
    return input.map((value) => toMap(value));
  }
  const result = {};
  Object.keys(input).forEach((key) => {
    result[key] = toMap(input[key]);
  });
  return result;
}

function flatMap(target: { [key: string]: any }, params: { [key: string]: any }, prefix: string = '') {
  if (prefix) {
    prefix = prefix + '.';
  }
  params = toMap(params);
  const keys = Object.keys(params);
  for (let i = 0; i < keys.length; i++) {
    let key = keys[i];
    const value = params[key];
    key = prefix + key;
    if (typeof value === 'undefined' || value == null) {
      continue;
    }
    if (Array.isArray(value)) {
      replaceRepeatList(target, value, key);
    } else if (value instanceof Object) {
      flatMap(target, value, key);
    } else {
      target[key] = value.toString();
    }
  }
  return target;
}

/** @internal */
export function toForm(filter: { [key: string]: any }): string {
  if (!filter) {
    return '';
  }
  const target = {};
  flatMap(target, filter);
  return querystring.stringify(target);
}

/** @internal */
export function parseXml(body: string): any {
  return $dara.XML.parseXml(body, null);
}
