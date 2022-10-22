import { Injectable, NotFoundException } from '@nestjs/common';
import { randomUUID } from 'crypto';
import { InjectQueue } from '../queue/queue.decorators';
import { QueueService } from '../queue/queue.service';
import {
  GetResultRequest,
  GetResultResponse,
  HandlerArgs,
  HandlerResult,
  StartRequest,
  StartResponse,
} from './app.interface';

@Injectable()
export class AppService {
  constructor(@InjectQueue() private readonly queue: QueueService) {}

  async put(os: string, data: HandlerArgs) {
    return this.queue.push(['request', os], data);
  }

  async get(id: string): Promise<HandlerResult | null> {
    return this.queue.get(['results', id]);
  }

  async start(params: StartRequest): Promise<StartResponse> {
    const { url, os } = params;
    const id = randomUUID();

    await this.put(os, { request_id: id, inspection_args: { url } });

    return { id };
  }

  async getResult(params: GetResultRequest): Promise<GetResultResponse> {
    const { id } = params;
    const result = await this.get(id);

    if (!result) {
      throw new NotFoundException();
    }

    const {
      inspection_result: { url, stderr, stdout },
    } = result;

    return { id, url, stderr, stdout };
  }
}
