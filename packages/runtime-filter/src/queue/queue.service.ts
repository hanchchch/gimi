export interface QueueService {
  push(key: string[], value: string): Promise<any>;

  get(key: string[]): Promise<string | null>;

  subOnce(key: string[], timeout: number): Promise<string | null>;
}
