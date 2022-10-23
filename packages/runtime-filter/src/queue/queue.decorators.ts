import { Inject } from "@nestjs/common";
import { QUEUE_SERVICE } from "./queue.symbol";

export const InjectQueue = () => Inject(QUEUE_SERVICE);
