import { Inject, Injectable } from '@nestjs/common';
import { RuntimeFilterService } from '@proto/nestjs/runtimefilter.interface';
import { ClientGrpc } from '@nestjs/microservices';

@Injectable()
export class AppService {
  private service: RuntimeFilterService;

  constructor(
    @Inject('RUNTIME_FILTER_PACKAGE') private readonly client: ClientGrpc
  ) {
    this.service = this.client.getService<RuntimeFilterService>(
      'RuntimeFilterService'
    );
  }
}
