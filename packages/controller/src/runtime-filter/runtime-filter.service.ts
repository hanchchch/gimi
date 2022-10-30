import { Inject, Injectable } from "@nestjs/common";
import {
  StartRequest,
  GetResultRequest,
  SubResultRequest,
} from "@proto/ts/messages/runtimefilter";
import { IRuntimeFilterService } from "@proto/nestjs/runtimefilter.interface";
import { ClientGrpc } from "@nestjs/microservices";
import { RUNTIME_FILTER_PACKAGE } from "./runtime-filter.symbols";

@Injectable()
export class RuntimeFilterService {
  private service: IRuntimeFilterService;

  constructor(
    @Inject(RUNTIME_FILTER_PACKAGE) private readonly client: ClientGrpc
  ) {
    this.service = this.client.getService<IRuntimeFilterService>(
      "RuntimeFilterService"
    );
  }

  start(data: StartRequest) {
    return this.service.Start(data);
  }

  getResult(data: GetResultRequest) {
    return this.service.GetResult(data);
  }

  subResult(data: SubResultRequest) {
    return this.service.SubResult(data);
  }
}
