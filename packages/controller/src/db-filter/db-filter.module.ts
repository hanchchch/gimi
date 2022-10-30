import { Module } from "@nestjs/common";
import { ConfigModule, ConfigService } from "@nestjs/config";
import { ClientsModule } from "@nestjs/microservices";
import { createClientOptions } from "@proto/nestjs/dbfilter.options";
import { EnvVars } from "../environments/environment.interface";
import { DbFilterService } from "./db-filter.service";

import { DB_FILTER_PACKAGE } from "./db-filter.symbols";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    ClientsModule.registerAsync([
      {
        name: DB_FILTER_PACKAGE,
        useFactory: (config: ConfigService<EnvVars>) =>
          createClientOptions({
            url: config.get("DB_FILTER_URL"),
          }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [],
  providers: [DbFilterService],
  exports: [DbFilterService],
})
export class DbFilterModule {}
