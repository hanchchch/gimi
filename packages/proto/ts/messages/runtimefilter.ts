/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "runtimefilter";

export interface StartRequest {
  url: string;
  os: string;
}

export interface StartResponse {
  id: string;
}

export interface GetResultRequest {
  id: string;
}

export interface GetResultResponse {
  id: string;
  url: string;
  malicious: boolean;
}

export interface SubResultRequest {
  id: string;
  timeout?: number | undefined;
}

function createBaseStartRequest(): StartRequest {
  return { url: "", os: "" };
}

export const StartRequest = {
  encode(
    message: StartRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.os !== "") {
      writer.uint32(18).string(message.os);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.os = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StartRequest {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      os: isSet(object.os) ? String(object.os) : "",
    };
  },

  toJSON(message: StartRequest): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.os !== undefined && (obj.os = message.os);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StartRequest>, I>>(
    object: I
  ): StartRequest {
    const message = createBaseStartRequest();
    message.url = object.url ?? "";
    message.os = object.os ?? "";
    return message;
  },
};

function createBaseStartResponse(): StartResponse {
  return { id: "" };
}

export const StartResponse = {
  encode(
    message: StartResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StartResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: StartResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StartResponse>, I>>(
    object: I
  ): StartResponse {
    const message = createBaseStartResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetResultRequest(): GetResultRequest {
  return { id: "" };
}

export const GetResultRequest = {
  encode(
    message: GetResultRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetResultRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetResultRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: GetResultRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetResultRequest>, I>>(
    object: I
  ): GetResultRequest {
    const message = createBaseGetResultRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetResultResponse(): GetResultResponse {
  return { id: "", url: "", malicious: false };
}

export const GetResultResponse = {
  encode(
    message: GetResultResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.url !== "") {
      writer.uint32(18).string(message.url);
    }
    if (message.malicious === true) {
      writer.uint32(24).bool(message.malicious);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetResultResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetResultResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.url = reader.string();
          break;
        case 3:
          message.malicious = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetResultResponse {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      url: isSet(object.url) ? String(object.url) : "",
      malicious: isSet(object.malicious) ? Boolean(object.malicious) : false,
    };
  },

  toJSON(message: GetResultResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.url !== undefined && (obj.url = message.url);
    message.malicious !== undefined && (obj.malicious = message.malicious);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetResultResponse>, I>>(
    object: I
  ): GetResultResponse {
    const message = createBaseGetResultResponse();
    message.id = object.id ?? "";
    message.url = object.url ?? "";
    message.malicious = object.malicious ?? false;
    return message;
  },
};

function createBaseSubResultRequest(): SubResultRequest {
  return { id: "", timeout: undefined };
}

export const SubResultRequest = {
  encode(
    message: SubResultRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.timeout !== undefined) {
      writer.uint32(16).int32(message.timeout);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubResultRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubResultRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.timeout = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubResultRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      timeout: isSet(object.timeout) ? Number(object.timeout) : undefined,
    };
  },

  toJSON(message: SubResultRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.timeout !== undefined &&
      (obj.timeout = Math.round(message.timeout));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SubResultRequest>, I>>(
    object: I
  ): SubResultRequest {
    const message = createBaseSubResultRequest();
    message.id = object.id ?? "";
    message.timeout = object.timeout ?? undefined;
    return message;
  },
};

export interface RuntimeFilterService {
  Start(request: StartRequest): Promise<StartResponse>;
  GetResult(request: GetResultRequest): Promise<GetResultResponse>;
  SubResult(request: SubResultRequest): Promise<GetResultResponse>;
}

export class RuntimeFilterServiceClientImpl implements RuntimeFilterService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "runtimefilter.RuntimeFilterService";
    this.rpc = rpc;
    this.Start = this.Start.bind(this);
    this.GetResult = this.GetResult.bind(this);
    this.SubResult = this.SubResult.bind(this);
  }
  Start(request: StartRequest): Promise<StartResponse> {
    const data = StartRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Start", data);
    return promise.then((data) => StartResponse.decode(new _m0.Reader(data)));
  }

  GetResult(request: GetResultRequest): Promise<GetResultResponse> {
    const data = GetResultRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetResult", data);
    return promise.then((data) =>
      GetResultResponse.decode(new _m0.Reader(data))
    );
  }

  SubResult(request: SubResultRequest): Promise<GetResultResponse> {
    const data = SubResultRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SubResult", data);
    return promise.then((data) =>
      GetResultResponse.decode(new _m0.Reader(data))
    );
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
