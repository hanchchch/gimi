export interface QueueService {
  push(key: string[], value: string): Promise<any>;

  get(key: string[]): Promise<Buffer | null>;
}
