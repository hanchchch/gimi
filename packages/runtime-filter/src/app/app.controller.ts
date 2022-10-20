import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';
import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { RuntimeFilterRequest, RuntimeFilterResponse } from './app.interface';

@Controller()
export class AppController {
  @GrpcMethod('RuntimeFilterService', 'RuntimeFilter')
  RuntimeFilter(
    data: RuntimeFilterRequest,
    metadata: Metadata,
    call: ServerUnaryCall<any, any>
  ): RuntimeFilterResponse {
    return {
      url: data.url,
      stdout: 'stdout',
      stderr: 'stderr',
    };
  }
}
