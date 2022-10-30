import { Logger } from "@nestjs/common";
import { NestFactory } from "@nestjs/core";
import { ConfigService } from "@nestjs/config";
import { createClientOptions } from "@proto/nestjs/runtimefilter.options";

import { AppModule } from "./app/app.module";
import { EnvVars } from "./environments/environment.interface";

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {});
  const config: ConfigService<EnvVars> = app.get(ConfigService);
  const options = createClientOptions({
    url: `0.0.0.0:${config.get("PORT", 50051)}`,
  });

  app.connectMicroservice(options);
  await app.startAllMicroservices();
  await app.init();
  Logger.log(`🚀 Runtime-filter (gRPC) is running on ${options.options.url}`);

  const httpPort = config.get("HTTP_PORT", "");
  if (httpPort) {
    await app.listen(httpPort);
    Logger.log(
      `🚀  Runtime-filter (HTTP) is running on: ${await app.getUrl()}`
    );
  }
}

bootstrap();
