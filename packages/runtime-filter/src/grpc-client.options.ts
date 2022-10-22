import { Transport, GrpcOptions } from '@nestjs/microservices';
import { join } from 'path';

export const createClientOptions = ({
  port = 50051,
  protoPath = join(
    __dirname,
    '..',
    '..',
    '..',
    'packages/proto/messages/runtimefilter.proto'
  ),
}: {
  port?: number;
  protoPath?: string;
}): GrpcOptions => ({
  transport: Transport.GRPC,
  options: {
    package: 'runtimefilter',
    url: `0.0.0.0:${port}`,
    protoPath,
  },
});
