import { HandlerResult } from "@proto/ts/messages/inspection";
import { QueueService } from "./queue.service";

export class MockQueueService implements QueueService {
  push(key: string[], value: string): Promise<any> {
    return Promise.resolve();
  }

  get(key: string[]): Promise<Buffer | null> {
    return Promise.resolve(
      Buffer.from(
        new TextDecoder().decode(
          HandlerResult.encode({
            requestId: "",
            result: {
              url: "",
              malicious: false,
              screenshot: "",
              reasons: [],
              locations: [],
              sendingTo: [],
              hosts: [],
              urls: [],
            },
          }).finish()
        )
      )
    );
  }
}
