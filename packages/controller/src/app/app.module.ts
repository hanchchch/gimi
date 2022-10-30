import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { DbFilterModule } from "../db-filter/db-filter.module";
import { RuntimeFilterModule } from "../runtime-filter/runtime-filter.module";

import { AppController } from "./app.controller";
import { AppService } from "./app.service";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    RuntimeFilterModule,
    DbFilterModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
