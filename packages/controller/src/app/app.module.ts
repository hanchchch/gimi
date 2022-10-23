import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { RuntimeFilterModule } from "../runtime-filter/runtime-filter.module";

import { AppController } from "./app.controller";
import { AppService } from "./app.service";

@Module({
  imports: [ConfigModule.forRoot({ isGlobal: true }), RuntimeFilterModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
