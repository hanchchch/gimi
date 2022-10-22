import { Injectable, OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { createClient } from 'redis';
import { EnvVars } from '../environments/environment.interface';
import { QueueService } from './queue.service';

type RedisClient = ReturnType<typeof createClient>;

@Injectable()
export class RedisService implements OnModuleInit, QueueService {
  private redisClient: RedisClient;
  private prefix = 'gimi:inspection';

  constructor(private readonly config: ConfigService<EnvVars>) {}

  async onModuleInit() {
    this.redisClient = createClient({
      url: this.config.get('REDIS_URL'),
    });
    await this.redisClient.connect();
  }

  buildKey(...args: string[]) {
    return [this.prefix, ...args].join(':');
  }

  encode<T>(data: T) {
    return JSON.stringify(data);
  }

  decode<T>(data: string): T {
    return JSON.parse(data);
  }

  async push<T>(key: string[], value: T) {
    return this.redisClient.lPush(this.buildKey(...key), this.encode<T>(value));
  }

  async get<T>(key: string[]): Promise<T | null> {
    const record = await this.redisClient.get(this.buildKey(...key));
    if (!record) {
      return null;
    }

    return this.decode<T>(record);
  }
}
