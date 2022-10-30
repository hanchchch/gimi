import { Injectable } from "@nestjs/common";
import { FindRequest } from "@proto/ts/messages/dbfilter";

@Injectable()
export class MockDbFilterService {
  find(data: FindRequest) {
    return { url: data.url, found: true };
  }
}
