import { Metadata, ServerUnaryCall } from "@grpc/grpc-js";
import { Controller } from "@nestjs/common";
import { GrpcMethod } from "@nestjs/microservices";
import { FindRequest, FindResponse } from "@proto/ts/messages/dbfilter";
import { AppService } from "./app.service";

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @GrpcMethod("DbFilterService", "Find")
  async Find(
    data: FindRequest,
    metadata: Metadata,
    call: ServerUnaryCall<any, any>
  ): Promise<FindResponse> {
    return this.appService.find(data);
  }
}
