import { Body, Controller, Get, Param, Post } from "@nestjs/common";
import { InspectDto } from "./inspection.dto";

import { InspectionService } from "./inspection.service";

@Controller(`inspection`)
export class InspectionController {
  constructor(private readonly inspectionService: InspectionService) {}

  @Post(``)
  async inspect(@Body() body: InspectDto) {
    return this.inspectionService.inspect(body);
  }

  @Get(`:id`)
  async fetchResult(@Param(`id`) id: string) {
    return this.inspectionService.fetchResult(id);
  }
}
