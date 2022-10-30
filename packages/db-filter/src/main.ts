import { Logger } from "@nestjs/common";
import { NestFactory } from "@nestjs/core";
import { ConfigService } from "@nestjs/config";
import { createClientOptions } from "@proto/nestjs/dbfilter.options";

import { AppModule } from "./app/app.module";
import { EnvVars } from "./environments/environment.interface";

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {});
  const config: ConfigService<EnvVars> = app.get(ConfigService);
  const options = createClientOptions({
    url: `0.0.0.0:${config.get("PORT", 50052)}`,
  });

  app.connectMicroservice(options);
  await app.startAllMicroservices();
  await app.init();
  Logger.log(`🚀 Db-filter (gRPC) is running on ${options.options.url}`);
}

bootstrap();
