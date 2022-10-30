/**
 * This is not a production server yet!
 * This is only a minimal backend to get started.
 */

import { Logger } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { NestFactory } from "@nestjs/core";

import { AppModule } from "./app/app.module";
import { EnvVars } from "./environments/environment.interface";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const config: ConfigService<EnvVars> = app.get(ConfigService);
  const port = config.get("PORT", 3000);

  await app.listen(port);
  Logger.log(`🚀 Controller is running on: ${await app.getUrl()}`);
}

bootstrap();
