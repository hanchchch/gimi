import { Injectable } from "@nestjs/common";
import { firstValueFrom } from "rxjs";
import { InjectRuntimeFilter } from "../runtime-filter/runtime-filter.decorators";
import { RuntimeFilterService } from "../runtime-filter/runtime-filter.service";
import { InspectParams } from "./app.dto";

@Injectable()
export class AppService {
  private defaultTimeout = 1000;
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService
  ) {}

  async inspect(params: InspectParams) {
    const { url, os, timeout = this.defaultTimeout } = params;
    const { id } = await firstValueFrom(
      this.runtimeFilterService.start({ url, os })
    );
    const { stdout, stderr } = await firstValueFrom(
      this.runtimeFilterService.subResult({ id, timeout })
    );
    return { stdout, stderr };
  }
}
