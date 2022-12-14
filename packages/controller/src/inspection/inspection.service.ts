import { Injectable, NotFoundException } from "@nestjs/common";
import { InjectRepository } from "@nestjs/typeorm";
import { randomUUID } from "crypto";
import { firstValueFrom } from "rxjs";
import { Repository } from "typeorm";
import { Inspection } from "../inspection/inspection.entity";
import { InjectRuntimeFilter } from "../runtime-filter/runtime-filter.decorators";
import { RuntimeFilterService } from "../runtime-filter/runtime-filter.service";
import { InspectParams } from "./inspection.dto";

@Injectable()
export class InspectionService {
  constructor(
    @InjectRuntimeFilter()
    private readonly runtimeFilterService: RuntimeFilterService,
    @InjectRepository(Inspection)
    private readonly inspectionRepository: Repository<Inspection>
  ) {}

  async inspect(params: InspectParams): Promise<Inspection> {
    const { url } = params;

    const inspection = await this.inspectionRepository.findOne({
      where: { url },
    });

    if (inspection) {
      return inspection;
    }

    const id = randomUUID();
    const newInspection = this.inspectionRepository.create({ url, id });

    await firstValueFrom(this.runtimeFilterService.start({ url, id }));

    await this.inspectionRepository.insert(newInspection);

    return newInspection;
  }

  async fetchInspection(id: string) {
    const inspection = await this.inspectionRepository.findOne({
      where: { id },
    });

    if (!inspection) {
      throw new NotFoundException();
    }

    if (!inspection.result) {
      try {
        const { result, error } = await firstValueFrom(
          this.runtimeFilterService.getResult({ id })
        );
        await this.inspectionRepository.update(id, { result, error });
      } catch (e) {
        if (!e.message.includes(`ResultNotFoundException`)) {
          throw e;
        }
      }
    }
    return inspection;
  }
}
