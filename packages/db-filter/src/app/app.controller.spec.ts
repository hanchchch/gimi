import { Metadata } from "@grpc/grpc-js";
import { Test, TestingModule } from "@nestjs/testing";
import { testConfigProvider } from "../environments/environment.test-env";

import { AppController } from "./app.controller";
import { AppService } from "./app.service";

const service = { find: jest.fn() };

describe("AppController", () => {
  let app: TestingModule;
  let controller: AppController;

  beforeAll(async () => {
    app = await Test.createTestingModule({
      controllers: [AppController],
      providers: [
        { provide: AppService, useValue: service },
        testConfigProvider,
      ],
    }).compile();

    controller = app.get<AppController>(AppController);
  });

  describe("Find", () => {
    it("should return find", async () => {
      await controller.Find(
        { url: "url" },
        new Metadata(),
        /* @ts-ignore */
        {}
      );

      expect(service.find).toBeCalledWith({ url: "url" });
    });
  });
});
