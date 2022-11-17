import { Module } from "@nestjs/common";
import { ConfigModule, ConfigService } from "@nestjs/config";
import { TypeOrmModule } from "@nestjs/typeorm";
import { EnvVars } from "../environments/environment.interface";

import { InspectionModule } from "../inspection/inspection.module";

import { AppController } from "./app.controller";
import { AppService } from "./app.service";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    TypeOrmModule.forRootAsync({
      useFactory: async (config: ConfigService<EnvVars>) => {
        return {
          type: "postgres",
          url: config.get("DB_URL"),
          synchronize: config.get("DB_SYNC") || false,
          autoLoadEntities: true,
        };
      },
      inject: [ConfigService],
    }),
    InspectionModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
