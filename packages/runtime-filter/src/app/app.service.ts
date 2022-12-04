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
import { ResultNotFoundException } from "./app.exceptions";

@Injectable()
export class AppService {
  constructor(@InjectQueue() private readonly queue: QueueService) {}

  async put(data: HandlerArgs) {
    return this.queue.push(
      ["request"],
      new TextDecoder().decode(HandlerArgs.encode(data).finish())
    );
  }

  async get(id: string): Promise<HandlerResult | null> {
    const buf = await this.queue.get(["results", id]);
    if (!buf) {
      return null;
    }
    return HandlerResult.decode(buf);
  }

  async start(params: StartRequest): Promise<StartResponse> {
    const { url } = params;
    const id = randomUUID();

    await this.put({ requestId: id, args: { url } });

    return { id };
  }

  async getResult(params: GetResultRequest): Promise<GetResultResponse> {
    const { id } = params;
    const result = await this.get(id);
    if (!result) {
      throw new ResultNotFoundException(id);
    }
    return { ...result };
  }
}
