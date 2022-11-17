import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { TypeOrmModule } from "@nestjs/typeorm";
import { InspectionResult } from "../inspection/inspection-result.entity";
import { RuntimeFilterModule } from "../runtime-filter/runtime-filter.module";

import { InspectionController } from "./inspection.controller";
import { InspectionService } from "./inspection.service";

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    TypeOrmModule.forFeature([InspectionResult]),
    RuntimeFilterModule,
  ],
  controllers: [InspectionController],
  providers: [InspectionService],
})
export class InspectionModule {}
