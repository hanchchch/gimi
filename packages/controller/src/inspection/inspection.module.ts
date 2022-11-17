import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { TypeOrmModule } from "@nestjs/typeorm";
import { RuntimeFilterModule } from "../runtime-filter/runtime-filter.module";

import { InspectionController } from "./inspection.controller";
import { Inspection } from "./inspection.entity";
import { InspectionService } from "./inspection.service";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    TypeOrmModule.forFeature([Inspection]),
    RuntimeFilterModule,
  ],
  controllers: [InspectionController],
  providers: [InspectionService],
})
export class InspectionModule {}
