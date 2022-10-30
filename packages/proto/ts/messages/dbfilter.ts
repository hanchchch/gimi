/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "dbfilter";

export interface FindRequest {
  url: string;
}

export interface FindResponse {
  url: string;
  found: boolean;
}

function createBaseFindRequest(): FindRequest {
  return { url: "" };
}

export const FindRequest = {
  encode(
    message: FindRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FindRequest {
    return { url: isSet(object.url) ? String(object.url) : "" };
  },

  toJSON(message: FindRequest): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FindRequest>, I>>(
    object: I
  ): FindRequest {
    const message = createBaseFindRequest();
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseFindResponse(): FindResponse {
  return { url: "", found: false };
}

export const FindResponse = {
  encode(
    message: FindResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.found === true) {
      writer.uint32(16).bool(message.found);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.found = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FindResponse {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      found: isSet(object.found) ? Boolean(object.found) : false,
    };
  },

  toJSON(message: FindResponse): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.found !== undefined && (obj.found = message.found);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FindResponse>, I>>(
    object: I
  ): FindResponse {
    const message = createBaseFindResponse();
    message.url = object.url ?? "";
    message.found = object.found ?? false;
    return message;
  },
};

export interface DbFilterService {
  Find(request: FindRequest): Promise<FindResponse>;
}

export class DbFilterServiceClientImpl implements DbFilterService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "dbfilter.DbFilterService";
    this.rpc = rpc;
    this.Find = this.Find.bind(this);
  }
  Find(request: FindRequest): Promise<FindResponse> {
    const data = FindRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Find", data);
    return promise.then((data) => FindResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & {
      [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
    };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
