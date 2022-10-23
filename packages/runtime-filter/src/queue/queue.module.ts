import { Module } from "@nestjs/common";
import { QUEUE_SERVICE } from "./queue.symbol";

import { RedisService } from "./redis.service";

@Module({
  imports: [],
  controllers: [],
  providers: [{ provide: QUEUE_SERVICE, useClass: RedisService }],
  exports: [QUEUE_SERVICE],
})
export class QueueModule {}
