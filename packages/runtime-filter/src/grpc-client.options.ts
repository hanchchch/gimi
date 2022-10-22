import { Transport, GrpcOptions } from '@nestjs/microservices';
import { join } from 'path';

export const grpcClientOptions: GrpcOptions = {
  transport: Transport.GRPC,
  options: {
    package: 'runtimefilter',
    url: '0.0.0.0:50051',
    protoPath: join(
      __dirname,
      '..',
      '..',
      '..',
      'packages/proto/messages/runtimefilter.proto'
    ),
  },
};
