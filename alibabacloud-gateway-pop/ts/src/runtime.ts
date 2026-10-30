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
  if (!body) {
    return {};
  }
  const reader = new XmlReader(body);
  reader.skipProlog();
  const root = reader.readElement();
  if (!root) {
    return {};
  }
  const result: { [key: string]: any } = {};
  result[root.name] = root.value;
  return result;
}

class XmlReader {
  private index: number;

  constructor(private readonly source: string) {
    this.index = 0;
  }

  skipProlog() {
    this.skipSpace();
    while (this.source.substr(this.index, 2) === '<?' || this.source.substr(this.index, 4) === '<!--') {
      if (this.source.substr(this.index, 4) === '<!--') {
        const end = this.source.indexOf('-->', this.index);
        this.index = end < 0 ? this.source.length : end + 3;
      } else {
        const end = this.source.indexOf('?>', this.index);
        this.index = end < 0 ? this.source.length : end + 2;
      }
      this.skipSpace();
    }
  }

  readElement(): { name: string, value: any } {
    this.skipSpace();
    if (this.source.charAt(this.index) !== '<') {
      return null;
    }
    this.index++;
    const name = this.readName();
    this.skipAttributes();
    if (this.source.substr(this.index, 2) === '/>') {
      this.index += 2;
      return { name: name, value: '' };
    }
    if (this.source.charAt(this.index) === '>') {
      this.index++;
    }
    const children: { [key: string]: any[] } = {};
    let text = '';
    let hasChild = false;
    while (this.index < this.source.length) {
      if (this.source.substr(this.index, 2) === '</') {
        const end = this.source.indexOf('>', this.index);
        this.index = end < 0 ? this.source.length : end + 1;
        break;
      }
      if (this.source.substr(this.index, 4) === '<!--') {
        const end = this.source.indexOf('-->', this.index);
        this.index = end < 0 ? this.source.length : end + 3;
        continue;
      }
      if (this.source.substr(this.index, 9) === '<![CDATA[') {
        const end = this.source.indexOf(']]>', this.index);
        text += end < 0 ? '' : this.source.substring(this.index + 9, end);
        this.index = end < 0 ? this.source.length : end + 3;
        continue;
      }
      if (this.source.charAt(this.index) === '<') {
        const child = this.readElement();
        if (!child) {
          break;
        }
        hasChild = true;
        if (!children[child.name]) {
          children[child.name] = [];
        }
        children[child.name].push(child.value);
        continue;
      }
      const next = this.source.indexOf('<', this.index);
      const chunk = this.source.substring(this.index, next < 0 ? this.source.length : next);
      text += decodeXml(chunk);
      this.index = next < 0 ? this.source.length : next;
    }
    if (!hasChild) {
      return { name: name, value: text.trim() };
    }
    const value: { [key: string]: any } = {};
    const names = Object.keys(children);
    for (let i = 0; i < names.length; i++) {
      const items = children[names[i]];
      value[names[i]] = items.length === 1 ? items[0] : items;
    }
    return { name: name, value: value };
  }

  private readName(): string {
    const start = this.index;
    while (this.index < this.source.length) {
      const ch = this.source.charAt(this.index);
      if (ch === ' ' || ch === '\n' || ch === '\r' || ch === '\t' || ch === '>' || ch === '/') {
        break;
      }
      this.index++;
    }
    return this.source.substring(start, this.index);
  }

  private skipAttributes() {
    while (this.index < this.source.length) {
      const ch = this.source.charAt(this.index);
      if (ch === '>' || ch === '/') {
        return;
      }
      this.index++;
    }
  }

  private skipSpace() {
    while (this.index < this.source.length) {
      const ch = this.source.charAt(this.index);
      if (ch !== ' ' && ch !== '\n' && ch !== '\r' && ch !== '\t') {
        return;
      }
      this.index++;
    }
  }
}

function decodeXml(value: string): string {
  return value
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&quot;/g, '"')
    .replace(/&apos;/g, '\'')
    .replace(/&amp;/g, '&');
}
