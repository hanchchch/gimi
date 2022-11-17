import { Injectable, NotFoundException } from "@nestjs/common";
import { InjectRepository } from "@nestjs/typeorm";
import { firstValueFrom } from "rxjs";
import { Repository } from "typeorm";
import { InspectionResult } from "../inspection/inspection-result.entity";
import { InjectRuntimeFilter } from "../runtime-filter/runtime-filter.decorators";
import { RuntimeFilterService } from "../runtime-filter/runtime-filter.service";
import { InspectParams } from "./inspection.dto";

@Injectable()
export class InspectionService {
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService,
    @InjectRepository(InspectionResult)
    private readonly inspectionRepository: Repository<InspectionResult>
  ) {}

  async inspect(params: InspectParams) {
    const { url, os } = params;

    const blacklist = await this.inspectionRepository.findOne({
      where: { url },
    });

    if (blacklist) {
      return { blacklist };
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
