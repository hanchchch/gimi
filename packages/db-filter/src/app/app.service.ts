import { Injectable } from "@nestjs/common";
import { FindRequest } from "@proto/ts/messages/dbfilter";

@Injectable()
export class AppService {
  async find(params: FindRequest) {
    const { url } = params;
    return { url, found: true };
  }
}
