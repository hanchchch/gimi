import { Body, Controller, Post } from '@nestjs/common';
import { InspectDto } from './app.dto';

import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post(`inspect`)
  async inspect(@Body() body: InspectDto) {
    return this.appService.inspect(body);
  }
}
