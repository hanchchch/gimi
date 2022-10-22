import { Logger } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions } from '@nestjs/microservices';

import { grpcClientOptions } from './grpc-client.options';
import { AppModule } from './app/app.module';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      ...grpcClientOptions,
    }
  );

  await app.listen();
  Logger.log(`🚀 Application is running on ${grpcClientOptions.options.url}`);
}

bootstrap();
