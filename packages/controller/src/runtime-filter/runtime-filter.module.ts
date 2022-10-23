import { Module } from "@nestjs/common";
import { ConfigModule, ConfigService } from "@nestjs/config";
import { ClientsModule } from "@nestjs/microservices";
import { createClientOptions } from "@proto/nestjs/runtimefilter.options";
import { EnvVars } from "../environments/environment.interface";
import { RuntimeFilterService } from "./runtime-filter.service";

import { RUNTIME_FILTER_PACKAGE } from "./runtime-filter.symbols";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    ClientsModule.registerAsync([
      {
        name: RUNTIME_FILTER_PACKAGE,
        useFactory: (config: ConfigService<EnvVars>) =>
          createClientOptions({
            url: config.get("RUNTIME_FILTER_URL"),
          }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [],
  providers: [RuntimeFilterService],
  exports: [RuntimeFilterService],
})
export class RuntimeFilterModule {}
