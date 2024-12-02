import * as jspb from 'google-protobuf'



export class PingMessage extends jspb.Message {
  getPayload(): string;
  setPayload(value: string): PingMessage;

  getSeq(): number;
  setSeq(value: number): PingMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingMessage.AsObject;
  static toObject(includeInstance: boolean, msg: PingMessage): PingMessage.AsObject;
  static serializeBinaryToWriter(message: PingMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingMessage;
  static deserializeBinaryFromReader(message: PingMessage, reader: jspb.BinaryReader): PingMessage;
}

export namespace PingMessage {
  export type AsObject = {
    payload: string,
    seq: number,
  }
}

