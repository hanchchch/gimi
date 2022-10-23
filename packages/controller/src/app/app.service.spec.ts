import { Test } from "@nestjs/testing";
import { mockRuntimeFilterServiceProvider } from "../runtime-filter/mock-runtime-filter.provider";
import { AppService } from "./app.service";

describe("AppService", () => {
  let service: AppService;

  beforeAll(async () => {
    const app = await Test.createTestingModule({
      providers: [AppService, mockRuntimeFilterServiceProvider],
    }).compile();

    service = app.get<AppService>(AppService);
  });

  describe("getData", () => {
    it("should be defined", () => {
      expect(service).toBeDefined();
    });
  });
});
