import {
  Controller,
  INestApplication,
  INestMicroservice,
  Inject,
  Module,
} from "@nestjs/common";
import { NestFactory } from "@nestjs/core";
import { ClientGrpc, ClientsModule } from "@nestjs/microservices";
import { Test, TestingModule } from "@nestjs/testing";
import { ConfigService } from "@nestjs/config";
import { firstValueFrom } from "rxjs";
import { createClientOptions } from "@proto/nestjs/dbfilter.options";
import { IDbFilterService } from "@proto/nestjs/dbfilter.interface";
import { FindRequest } from "@proto/ts/messages/dbfilter";
import { AppModule } from "./app/app.module";
import { testConfigProvider } from "./environments/environment.test-env";
import { getRepositoryToken } from "@nestjs/typeorm";
import { Blacklist } from "./app/blacklist.entity";

const clientOptions = createClientOptions({});

@Controller()
class TestClientController {
  private service: IDbFilterService;

  constructor(
    @Inject("DB_FILTER_PACKAGE") private readonly client: ClientGrpc
  ) {
    this.service = this.client.getService<IDbFilterService>("DbFilterService");
  }

  getService() {
    return this.service;
  }

  getClient() {
    return this.client;
  }

  find(data: FindRequest) {
    return this.service.Find(data);
  }
}

@Module({
  imports: [
    ClientsModule.register([
      {
        name: "DB_FILTER_PACKAGE",
        ...clientOptions,
      },
    ]),
  ],
  controllers: [TestClientController],
})
class TestClientModule {}

describe("App", () => {
  let app: INestMicroservice;
  let clientApp: INestApplication;

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      providers: [
        {
          provide: getRepositoryToken(Blacklist),
          useValue: { findOne: jest.fn(() => ({ found: true })) },
        },
      ],
      imports: [AppModule],
    })
      .overrideProvider(ConfigService)
      .useValue(testConfigProvider.useValue)
      .compile();

    app = moduleFixture.createNestMicroservice(clientOptions);
    await app.listen();

    clientApp = await NestFactory.create(TestClientModule);
    await clientApp.init();
  });

  afterAll(async () => {
    await app.close();
    await clientApp.close();
  });

  describe("client", () => {
    it("should be able to define service", async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      expect(controller).toBeDefined();
      expect(controller.getClient()).toBeDefined();
      expect(controller.getService()).toBeDefined();
    });

    it("should be able to call rpc Find", async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      const result = await firstValueFrom(controller.find({ url: "url" }));

      expect(result).toBeDefined();
      expect(result.found).toBeDefined();
    });
  });
});
