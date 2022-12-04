import { Test } from "@nestjs/testing";
import { testConfigProvider } from "../environments/environment.test-env";
import { MockQueueService } from "../queue/mock-queue.service";
import { QUEUE_SERVICE } from "../queue/queue.symbol";

import { AppService } from "./app.service";

describe("AppService", () => {
  let service: AppService;

  beforeAll(async () => {
    const app = await Test.createTestingModule({
      providers: [
        AppService,
        testConfigProvider,
        { provide: QUEUE_SERVICE, useClass: MockQueueService },
      ],
    }).compile();

    service = app.get<AppService>(AppService);
  });

  describe("start", () => {
    it("should return request id", async () => {
      const result = await service.start({ url: "url" });
      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
    });
  });

  describe("getResult", () => {
    it("should return result", async () => {
      const result = await service.getResult({ id: "id" });

      expect(result).toBeDefined();
      expect(result.result).toBeDefined();
    });
  });
});
