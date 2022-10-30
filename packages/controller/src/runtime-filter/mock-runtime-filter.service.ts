import { Injectable } from "@nestjs/common";
import {
  GetResultRequest,
  StartRequest,
  SubResultRequest,
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

  subResult(data: SubResultRequest) {
    return this.getResult({ id: data.id });
  }
}
