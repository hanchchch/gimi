import { Injectable, NotFoundException } from "@nestjs/common";
import { firstValueFrom } from "rxjs";
import { InjectDbFilter } from "../db-filter/db-filter.decorators";
import { DbFilterService } from "../db-filter/db-filter.service";
import { InjectRuntimeFilter } from "../runtime-filter/runtime-filter.decorators";
import { RuntimeFilterService } from "../runtime-filter/runtime-filter.service";
import { InspectParams } from "./app.dto";

@Injectable()
export class AppService {
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService,
    @InjectDbFilter()
    private readonly dbFilterService: DbFilterService
  ) {}

  async inspect(params: InspectParams) {
    const { url, os } = params;

    const { found, blacklist } = await firstValueFrom(
      this.dbFilterService.find({ url })
    );

    if (found) {
      return { malicious: true };
    }

    const { id } = await firstValueFrom(
      this.runtimeFilterService.start({ url, os })
    );

    return { id };
  }

  async fetchResult(id: string) {
    try {
      const result = await firstValueFrom(
        this.runtimeFilterService.getResult({ id })
      );
      return { id, ...result };
    } catch (e) {
      if (e.message.includes(`ResultNotFoundException`)) {
        throw new NotFoundException();
      }
      throw e;
    }
  }
}
