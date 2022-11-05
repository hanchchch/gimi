import { Body, Controller, Get, Param, Post } from "@nestjs/common";
import { InspectDto } from "./app.dto";

import { AppService } from "./app.service";

@Controller(``)
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post(`inspect`)
  async inspect(@Body() body: InspectDto) {
    return this.appService.inspect(body);
  }

  @Get(`results/:id`)
  async fetchResult(@Param(`id`) id: string) {
    return this.appService.fetchResult(id);
  }
}
