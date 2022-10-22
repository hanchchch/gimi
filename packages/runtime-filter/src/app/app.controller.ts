import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';
import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import {
  GetResultRequest,
  GetResultResponse,
  StartRequest,
  StartResponse,
} from '@proto/nestjs/runtimefilter.interface';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @GrpcMethod('RuntimeFilterService', 'Start')
  async Start(
    data: StartRequest,
    metadata: Metadata,
    call: ServerUnaryCall<any, any>
  ): Promise<StartResponse> {
    return this.appService.start(data);
  }

  @GrpcMethod('RuntimeFilterService', 'GetResult')
  async GetResult(
    data: GetResultRequest,
    metadata: Metadata,
    call: ServerUnaryCall<any, any>
  ): Promise<GetResultResponse> {
    return this.appService.getResult(data);
  }

  @Post(`start`)
  async start(@Body() params: StartRequest) {
    return this.appService.start(params);
  }

  @Get(`result/:id`)
  async getResult(@Param('id') id: string) {
    return this.appService.getResult({ id });
  }
}
