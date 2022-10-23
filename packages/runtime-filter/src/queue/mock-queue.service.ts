import { QueueService } from "./queue.service";

export class MockQueueService implements QueueService {
  encode<T>(data: T): string {
    return JSON.stringify(data);
  }

  decode<T>(data: string): T {
    return JSON.parse(data);
  }

  push<T>(key: string[], value: T): Promise<any> {
    return Promise.resolve();
  }

  get(key: string[]): Promise<any> {
    return Promise.resolve({
      request_id: "",
      inspection_result: { url: "", stderr: "", stdout: "" },
    });
  }
}
