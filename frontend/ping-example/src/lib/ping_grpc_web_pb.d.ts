import * as grpcWeb from 'grpc-web';

import * as ping_pb from './ping_pb'; // proto import: "ping.proto"


export class RootClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: ping_pb.PingMessage,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: ping_pb.PingMessage) => void
  ): grpcWeb.ClientReadableStream<ping_pb.PingMessage>;

}

export class RootPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: ping_pb.PingMessage,
    metadata?: grpcWeb.Metadata
  ): Promise<ping_pb.PingMessage>;

}

