/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "inspection";

export interface InspectionArgs {
  url: string;
}

export interface InspectionResult {
  url: string;
  stdout: string;
  stderr: string;
}

export interface HandlerArgs {
  requestId: string;
  args: InspectionArgs | undefined;
}

export interface HandlerResult {
  requestId: string;
  result: InspectionResult | undefined;
}

function createBaseInspectionArgs(): InspectionArgs {
  return { url: "" };
}

export const InspectionArgs = {
  encode(
    message: InspectionArgs,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InspectionArgs {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInspectionArgs();
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

  fromJSON(object: any): InspectionArgs {
    return { url: isSet(object.url) ? String(object.url) : "" };
  },

  toJSON(message: InspectionArgs): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InspectionArgs>, I>>(
    object: I
  ): InspectionArgs {
    const message = createBaseInspectionArgs();
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseInspectionResult(): InspectionResult {
  return { url: "", stdout: "", stderr: "" };
}

export const InspectionResult = {
  encode(
    message: InspectionResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.stdout !== "") {
      writer.uint32(18).string(message.stdout);
    }
    if (message.stderr !== "") {
      writer.uint32(26).string(message.stderr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InspectionResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInspectionResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.stdout = reader.string();
          break;
        case 3:
          message.stderr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InspectionResult {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      stdout: isSet(object.stdout) ? String(object.stdout) : "",
      stderr: isSet(object.stderr) ? String(object.stderr) : "",
    };
  },

  toJSON(message: InspectionResult): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.stdout !== undefined && (obj.stdout = message.stdout);
    message.stderr !== undefined && (obj.stderr = message.stderr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InspectionResult>, I>>(
    object: I
  ): InspectionResult {
    const message = createBaseInspectionResult();
    message.url = object.url ?? "";
    message.stdout = object.stdout ?? "";
    message.stderr = object.stderr ?? "";
    return message;
  },
};

function createBaseHandlerArgs(): HandlerArgs {
  return { requestId: "", args: undefined };
}

export const HandlerArgs = {
  encode(
    message: HandlerArgs,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.requestId !== "") {
      writer.uint32(10).string(message.requestId);
    }
    if (message.args !== undefined) {
      InspectionArgs.encode(message.args, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HandlerArgs {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHandlerArgs();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requestId = reader.string();
          break;
        case 2:
          message.args = InspectionArgs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HandlerArgs {
    return {
      requestId: isSet(object.requestId) ? String(object.requestId) : "",
      args: isSet(object.args)
        ? InspectionArgs.fromJSON(object.args)
        : undefined,
    };
  },

  toJSON(message: HandlerArgs): unknown {
    const obj: any = {};
    message.requestId !== undefined && (obj.requestId = message.requestId);
    message.args !== undefined &&
      (obj.args = message.args
        ? InspectionArgs.toJSON(message.args)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HandlerArgs>, I>>(
    object: I
  ): HandlerArgs {
    const message = createBaseHandlerArgs();
    message.requestId = object.requestId ?? "";
    message.args =
      object.args !== undefined && object.args !== null
        ? InspectionArgs.fromPartial(object.args)
        : undefined;
    return message;
  },
};

function createBaseHandlerResult(): HandlerResult {
  return { requestId: "", result: undefined };
}

export const HandlerResult = {
  encode(
    message: HandlerResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.requestId !== "") {
      writer.uint32(10).string(message.requestId);
    }
    if (message.result !== undefined) {
      InspectionResult.encode(
        message.result,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HandlerResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHandlerResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requestId = reader.string();
          break;
        case 2:
          message.result = InspectionResult.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HandlerResult {
    return {
      requestId: isSet(object.requestId) ? String(object.requestId) : "",
      result: isSet(object.result)
        ? InspectionResult.fromJSON(object.result)
        : undefined,
    };
  },

  toJSON(message: HandlerResult): unknown {
    const obj: any = {};
    message.requestId !== undefined && (obj.requestId = message.requestId);
    message.result !== undefined &&
      (obj.result = message.result
        ? InspectionResult.toJSON(message.result)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HandlerResult>, I>>(
    object: I
  ): HandlerResult {
    const message = createBaseHandlerResult();
    message.requestId = object.requestId ?? "";
    message.result =
      object.result !== undefined && object.result !== null
        ? InspectionResult.fromPartial(object.result)
        : undefined;
    return message;
  },
};

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
