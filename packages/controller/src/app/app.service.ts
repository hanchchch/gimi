import { Injectable } from "@nestjs/common";

@Injectable()
export class AppService {
  async healthCheck() {
    return { message: "OK" };
  }
}
