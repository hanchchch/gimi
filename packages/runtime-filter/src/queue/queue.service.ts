export interface QueueService {
  encode<T>(data: T): string;

  decode<T>(data: string): T;

  push<T>(key: string[], value: T): Promise<any>;

  get<T>(key: string[]): Promise<T>;
}
