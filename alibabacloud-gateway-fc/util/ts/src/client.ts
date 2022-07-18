// This file is auto-generated, don't edit it
import Credential from '@alicloud/credentials';
import * as $tea from '@alicloud/tea-typescript';
import * as http from 'http';


export default class Client {
  _credential: Credential;

  constructor(cred: Credential) {
    this._credential = cred;
  }


  async InvokeHTTPTrigger(url: string, method: string, body: Buffer, headers: http.OutgoingHttpHeaders): Promise<http.IncomingMessage> {
    throw new Error('Un-implemented!');
  }

  async InvokeAnonymousHTTPTrigger(url: string, method: string, body: Buffer, headers: http.OutgoingHttpHeaders): Promise<http.IncomingMessage> {
    throw new Error('Un-implemented!');
  }

  async SendHTTPRequestWithAuthorization(req: http.ClientRequest): Promise<http.IncomingMessage> {
    throw new Error('Un-implemented!');
  }

  async SendHTTPRequest(req: http.ClientRequest): Promise<http.IncomingMessage> {
    throw new Error('Un-implemented!');
  }

  async SignRequest(req: http.ClientRequest): Promise<http.ClientRequest> {
    throw new Error('Un-implemented!');
  }

  async SignRequestWithContentMD5(req: http.ClientRequest, contentMD5: string): Promise<http.ClientRequest> {
    throw new Error('Un-implemented!');
  }

  async BuildHTTPRequest(url: string, method: string, body: Buffer, headers: http.OutgoingHttpHeaders): Promise<http.ClientRequest> {
    throw new Error('Un-implemented!');
  }

}
