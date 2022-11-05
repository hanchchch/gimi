import { HandlerResult } from "@proto/ts/messages/inspection";
import { QueueService } from "./queue.service";

export class MockQueueService implements QueueService {
  push(key: string[], value: string): Promise<any> {
    return Promise.resolve();
  }

  get(key: string[]): Promise<string | null> {
    return Promise.resolve(
      new TextDecoder().decode(
        HandlerResult.encode({
          requestId: "",
          result: { url: "", malicious: false, locations: [] },
        }).finish()
      )
    );
  }

  subOnce(key: string[], timeout: number): Promise<string | null> {
    return this.get(key);
  }
}
