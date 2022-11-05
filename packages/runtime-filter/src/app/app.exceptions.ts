import { RpcException } from "@nestjs/microservices";

export class ResultNotFoundException extends RpcException {
  constructor(id: string) {
    super(`ResultNotFoundException: Result with id ${id} not found`);
  }
}
