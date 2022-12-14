import { Injectable } from "@nestjs/common";
import {
  GetResultRequest,
  StartRequest,
} from "@proto/ts/messages/runtimefilter";
import { randomUUID } from "crypto";

@Injectable()
export class MockRuntimeFilterService {
  start(data: StartRequest) {
    return { id: randomUUID() };
  }

  getResult(data: GetResultRequest) {
    return {
      id: data.id,
      url: "example.com",
      stdout: "",
      stderr: "",
    };
  }
}
