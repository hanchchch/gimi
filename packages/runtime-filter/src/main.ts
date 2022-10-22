import { Logger } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { ConfigService } from '@nestjs/config';
import { createClientOptions } from '@proto/nestjs/runtimefilter.options';

import { AppModule } from './app/app.module';
import { EnvVars } from './environments/environment.interface';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {});
  const config: ConfigService<EnvVars> = app.get(ConfigService);
  const options = createClientOptions({
    url: `0.0.0.0:${config.get('PORT', 50051)}`,
    protoPath: join(
      __dirname,
      '..',
      '..',
      '..',
      'packages/proto/messages/runtimefilter.proto'
    ),
  });

  app.connectMicroservice(options);
  await app.startAllMicroservices();
  await app.listen(config.get('HTTP_PORT', 3000));
  Logger.log(`🚀 Application is running on ${options.options.url}`);
}

bootstrap();
