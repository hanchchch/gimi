import { Injectable, OnModuleInit } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { createClient } from "redis";
import { EnvVars } from "../environments/environment.interface";
import { QueueService } from "./queue.service";

type RedisClient = ReturnType<typeof createClient>;

@Injectable()
export class RedisService implements OnModuleInit, QueueService {
  private redisClient: RedisClient;
  private prefix = "gimi:inspection";

  constructor(private readonly config: ConfigService<EnvVars>) {}

  async onModuleInit() {
    this.redisClient = createClient({
      url: this.config.get("REDIS_URL"),
    });
    await this.redisClient.connect();
  }

  buildKey(...args: string[]) {
    return [this.prefix, ...args].join(":");
  }

  async push(key: string[], value: string) {
    return this.redisClient.lPush(this.buildKey(...key), value);
  }

  async get(key: string[]): Promise<any | null> {
    return this.redisClient.sendCommand(["GET", this.buildKey(...key)], {
      returnBuffers: true,
    });
  }

  async subOnce(key: string[], timeout: number): Promise<Buffer | null> {
    const builtKey = this.buildKey(...key);
    const message = await Promise.race([
      new Promise<Buffer>((resolve, reject) => {
        this.redisClient
          .subscribe(
            builtKey,
            (message) => {
              resolve(message);
            },
            true
          )
          .catch(reject);
      }),
      new Promise<null>((resolve) => setTimeout(() => resolve(null), timeout)),
    ]);
    await this.redisClient.unsubscribe(builtKey);
    return message;
  }
}
