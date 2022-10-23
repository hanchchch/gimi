import { Injectable } from '@nestjs/common';
import { firstValueFrom } from 'rxjs';
import { InjectRuntimeFilter } from '../runtime-filter/runtime-filter.decorators';
import { RuntimeFilterService } from '../runtime-filter/runtime-filter.service';
import { InspectParams } from './app.dto';

@Injectable()
export class AppService {
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService
  ) {}

  async inspect(params: InspectParams) {
    const { url, os } = params;
    const { id } = await firstValueFrom(
      this.runtimeFilterService.start({ url, os })
    );
    // TODO: add blocking get - maybe with pub/sub
    const { stdout, stderr } = await firstValueFrom(
      this.runtimeFilterService.getResult({ id })
    );
    return { stdout, stderr };
  }
}
