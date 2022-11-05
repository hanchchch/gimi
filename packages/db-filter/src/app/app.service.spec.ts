import { Test } from "@nestjs/testing";
import { getRepositoryToken } from "@nestjs/typeorm";
import { testConfigProvider } from "../environments/environment.test-env";

import { AppService } from "./app.service";
import { Blacklist } from "./blacklist.entity";

const repo = {
  findOne: jest.fn(),
};

describe("AppService", () => {
  let service: AppService;

  beforeAll(async () => {
    const app = await Test.createTestingModule({
      providers: [
        AppService,
        testConfigProvider,
        { provide: getRepositoryToken(Blacklist), useValue: repo },
      ],
    }).compile();

    service = app.get<AppService>(AppService);
  });

  describe("find", () => {
    it("should return found", async () => {
      await service.find({ url: "url" });
      expect(repo.findOne).toBeCalledWith({ where: { url: "url" } });
    });
  });
});
