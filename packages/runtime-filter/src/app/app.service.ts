import { Injectable, OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { randomUUID } from 'crypto';
import { createClient } from 'redis';
import { EnvVars } from '../environments/environment.interface';
import {
  GetResultRequest,
  GetResultResponse,
  StartRequest,
  StartResponse,
} from './app.interface';

type RedisClient = ReturnType<typeof createClient>;

@Injectable()
export class AppService implements OnModuleInit {
  private redisClient: RedisClient;

  constructor(private readonly config: ConfigService<EnvVars>) {}

  async onModuleInit() {
    this.redisClient = createClient({
      url: this.config.get('REDIS_URL'),
    });
    await this.redisClient.connect();
  }

  async start(params: StartRequest): Promise<StartResponse> {
    const id = randomUUID();
    return { id };
  }

  async getResult(params: GetResultRequest): Promise<GetResultResponse> {
    return { id: params.id, url: '', stdout: '', stderr: '' };
  }
}
