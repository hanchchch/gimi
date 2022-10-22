import { Logger } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';

import { createClientOptions } from './grpc-client.options';
import { AppModule } from './app/app.module';
import { ConfigService } from '@nestjs/config';
import { EnvVars } from './environments/environment.interface';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {});
  const config: ConfigService<EnvVars> = app.get(ConfigService);
  const options = createClientOptions({
    port: config.get('PORT'),
    protoPath: config.get('PROTO_PATH'),
  });

  app.connectMicroservice(options);
  await app.startAllMicroservices();
  Logger.log(`🚀 Application is running on ${options.options.url}`);
}

bootstrap();
