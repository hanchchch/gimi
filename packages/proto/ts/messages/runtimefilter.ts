/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { InspectionResult } from "./inspection";

export const protobufPackage = "runtimefilter";

export interface StartRequest {
  id: string;
  url: string;
}

export interface StartResponse {
  id: string;
}

export interface GetResultRequest {
  id: string;
}

export interface GetResultResponse {
  result: InspectionResult | undefined;
  error?: string | undefined;
}

function createBaseStartRequest(): StartRequest {
  return { id: "", url: "" };
}

export const StartRequest = {
  encode(message: StartRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.url !== "") {
      writer.uint32(18).string(message.url);
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
          message.id = reader.string();
          break;
        case 2:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StartRequest {
    return { id: isSet(object.id) ? String(object.id) : "", url: isSet(object.url) ? String(object.url) : "" };
  },

  toJSON(message: StartRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StartRequest>, I>>(object: I): StartRequest {
    const message = createBaseStartRequest();
    message.id = object.id ?? "";
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseStartResponse(): StartResponse {
  return { id: "" };
}

export const StartResponse = {
  encode(message: StartResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  fromPartial<I extends Exact<DeepPartial<StartResponse>, I>>(object: I): StartResponse {
    const message = createBaseStartResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetResultRequest(): GetResultRequest {
  return { id: "" };
}

export const GetResultRequest = {
  encode(message: GetResultRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  fromPartial<I extends Exact<DeepPartial<GetResultRequest>, I>>(object: I): GetResultRequest {
    const message = createBaseGetResultRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseGetResultResponse(): GetResultResponse {
  return { result: undefined, error: undefined };
}

export const GetResultResponse = {
  encode(message: GetResultResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== undefined) {
      InspectionResult.encode(message.result, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      writer.uint32(18).string(message.error);
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
          message.result = InspectionResult.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = reader.string();
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
      result: isSet(object.result) ? InspectionResult.fromJSON(object.result) : undefined,
      error: isSet(object.error) ? String(object.error) : undefined,
    };
  },

  toJSON(message: GetResultResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = message.result ? InspectionResult.toJSON(message.result) : undefined);
    message.error !== undefined && (obj.error = message.error);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetResultResponse>, I>>(object: I): GetResultResponse {
    const message = createBaseGetResultResponse();
    message.result = (object.result !== undefined && object.result !== null)
      ? InspectionResult.fromPartial(object.result)
      : undefined;
    message.error = object.error ?? undefined;
    return message;
  },
};

export interface RuntimeFilterService {
  Start(request: StartRequest): Promise<StartResponse>;
  GetResult(request: GetResultRequest): Promise<GetResultResponse>;
}

export class RuntimeFilterServiceClientImpl implements RuntimeFilterService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "runtimefilter.RuntimeFilterService";
    this.rpc = rpc;
    this.Start = this.Start.bind(this);
    this.GetResult = this.GetResult.bind(this);
  }
  Start(request: StartRequest): Promise<StartResponse> {
    const data = StartRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Start", data);
    return promise.then((data) => StartResponse.decode(new _m0.Reader(data)));
  }

  GetResult(request: GetResultRequest): Promise<GetResultResponse> {
    const data = GetResultRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetResult", data);
    return promise.then((data) => GetResultResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
