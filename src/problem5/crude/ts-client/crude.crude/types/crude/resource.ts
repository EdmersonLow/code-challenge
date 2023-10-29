/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crude.crude";

export interface Resource {
  title: string;
  body: string;
  creatorr: string;
  idcreator: number;
}

const baseResource: object = {
  title: "",
  body: "",
  creatorr: "",
  idcreator: 0,
};

export const Resource = {
  encode(message: Resource, writer: Writer = Writer.create()): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.body !== "") {
      writer.uint32(18).string(message.body);
    }
    if (message.creatorr !== "") {
      writer.uint32(26).string(message.creatorr);
    }
    if (message.idcreator !== 0) {
      writer.uint32(32).uint64(message.idcreator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Resource {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseResource } as Resource;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.body = reader.string();
          break;
        case 3:
          message.creatorr = reader.string();
          break;
        case 4:
          message.idcreator = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Resource {
    const message = { ...baseResource } as Resource;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = String(object.body);
    } else {
      message.body = "";
    }
    if (object.creatorr !== undefined && object.creatorr !== null) {
      message.creatorr = String(object.creatorr);
    } else {
      message.creatorr = "";
    }
    if (object.idcreator !== undefined && object.idcreator !== null) {
      message.idcreator = Number(object.idcreator);
    } else {
      message.idcreator = 0;
    }
    return message;
  },

  toJSON(message: Resource): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.body !== undefined && (obj.body = message.body);
    message.creatorr !== undefined && (obj.creatorr = message.creatorr);
    message.idcreator !== undefined && (obj.idcreator = message.idcreator);
    return obj;
  },

  fromPartial(object: DeepPartial<Resource>): Resource {
    const message = { ...baseResource } as Resource;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = object.body;
    } else {
      message.body = "";
    }
    if (object.creatorr !== undefined && object.creatorr !== null) {
      message.creatorr = object.creatorr;
    } else {
      message.creatorr = "";
    }
    if (object.idcreator !== undefined && object.idcreator !== null) {
      message.idcreator = object.idcreator;
    } else {
      message.idcreator = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
