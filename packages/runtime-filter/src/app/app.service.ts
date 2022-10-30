import { Injectable, NotFoundException } from "@nestjs/common";
import { randomUUID } from "crypto";
import {
  GetResultRequest,
  GetResultResponse,
  StartRequest,
  StartResponse,
} from "@proto/ts/messages/runtimefilter";
import { HandlerArgs, HandlerResult } from "@proto/ts/messages/inspection";
import { InjectQueue } from "../queue/queue.decorators";
import { QueueService } from "../queue/queue.service";

@Injectable()
export class AppService {
  constructor(@InjectQueue() private readonly queue: QueueService) {}

  async put(os: string, data: HandlerArgs) {
    return this.queue.push(
      ["request", os],
      new TextDecoder().decode(HandlerArgs.encode(data).finish())
    );
  }

  async get(id: string): Promise<HandlerResult | null> {
    const str = await this.queue.get(["result", id]);
    if (!str) {
      return null;
    }
    return HandlerResult.decode(new TextEncoder().encode(str));
  }

  async start(params: StartRequest): Promise<StartResponse> {
    const { url, os } = params;
    const id = randomUUID();

    await this.put(os, { requestId: id, args: { url } });

    return { id };
  }

  async getResult(params: GetResultRequest): Promise<GetResultResponse> {
    const { id } = params;
    const result = await this.get(id);

    if (!result) {
      throw new NotFoundException();
    }

    const {
      result: { url, stderr, stdout },
    } = result;

    return { id, url, stderr, stdout };
  }
}
