export interface QueueService {
  push(key: string[], value: string): Promise<any>;

  get(key: string[]): Promise<Buffer | null>;

  subOnce(key: string[], timeout: number): Promise<Buffer | null>;
}
