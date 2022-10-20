import { Transport, ClientOptions } from '@nestjs/microservices';
import { join } from 'path';

export const grpcClientOptions: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    package: 'chopstick',
    url: 'localhost:50051',
    protoPath: join(
      __dirname,
      '..',
      '..',
      '..',
      'packages/proto/messages/runtime-filter.proto'
    ),
  },
};
