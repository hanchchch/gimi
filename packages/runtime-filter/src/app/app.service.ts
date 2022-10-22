import { Injectable, NotFoundException, OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { randomUUID } from 'crypto';
import { createClient } from 'redis';
import { EnvVars } from '../environments/environment.interface';
import {
  GetResultRequest,
  GetResultResponse,
  HandlerArgs,
  HandlerResult,
  StartRequest,
  StartResponse,
} from './app.interface';

type RedisClient = ReturnType<typeof createClient>;

@Injectable()
export class AppService implements OnModuleInit {
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

  async put({ id, url, os }: { id: string; url: string; os: string }) {
    return this.redisClient.lPush(
      this.buildKey('request', os),
      this.encode<HandlerArgs>({ request_id: id, inspection_args: { url } })
    );
  }

  async get(id: string): Promise<HandlerResult | null> {
    const record = await this.redisClient.get(this.buildKey('results', id));
    if (!record) {
      return null;
    }

    return this.decode<HandlerResult>(record);
  }

  async start(params: StartRequest): Promise<StartResponse> {
    const { url, os } = params;
    const id = randomUUID();

    await this.put({ id, url, os });

    return { id };
  }

  async getResult(params: GetResultRequest): Promise<GetResultResponse> {
    const { id } = params;
    const result = await this.get(id);

    if (!result) {
      throw new NotFoundException();
    }

    const {
      inspection_result: { url, stderr, stdout },
    } = result;

    return { id, url, stderr, stdout };
  }
}
