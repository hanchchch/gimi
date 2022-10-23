import { Injectable } from '@nestjs/common';
import {
  GetResultRequest,
  StartRequest,
} from '@proto/nestjs/runtimefilter.interface';
import { InjectRuntimeFilter } from '../runtime-filter/runtime-filter.decorators';
import { RuntimeFilterService } from '../runtime-filter/runtime-filter.service';

@Injectable()
export class AppService {
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService
  ) {}

  start(data: StartRequest) {
    return this.runtimeFilterService.start(data);
  }

  getResult(data: GetResultRequest) {
    return this.runtimeFilterService.getResult(data);
  }
}
