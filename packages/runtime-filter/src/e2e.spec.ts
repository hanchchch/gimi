import {
  Controller,
  INestApplication,
  INestMicroservice,
  Inject,
  Module,
} from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { ClientGrpc, ClientsModule } from '@nestjs/microservices';
import { createClientOptions } from './grpc-client.options';
import { AppModule } from './app/app.module';
import {
  GetResultRequest,
  StartRequest,
  RuntimeFilterService,
} from './app/app.interface';
import { firstValueFrom } from 'rxjs';
import { Test, TestingModule } from '@nestjs/testing';
import { ConfigService } from '@nestjs/config';
import { testConfigProvider } from './environments/environment.test-env';
import { QUEUE_SERVICE } from './queue/queue.symbol';
import { MockQueueService } from './queue/mock-queue.service';

const clientOptions = createClientOptions({});

@Controller()
class TestClientController {
  private service: RuntimeFilterService;

  constructor(
    @Inject('RUNTIME_FILTER_PACKAGE') private readonly client: ClientGrpc
  ) {
    this.service = this.client.getService<RuntimeFilterService>(
      'RuntimeFilterService'
    );
  }

  getService() {
    return this.service;
  }

  getClient() {
    return this.client;
  }

  start(data: StartRequest) {
    return this.service.Start(data);
  }

  getResult(data: GetResultRequest) {
    return this.service.GetResult(data);
  }
}

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'RUNTIME_FILTER_PACKAGE',
        ...clientOptions,
      },
    ]),
  ],
  controllers: [TestClientController],
})
class TestClientModule {}

describe('App', () => {
  let app: INestMicroservice;
  let clientApp: INestApplication;

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    })
      .overrideProvider(ConfigService)
      .useValue(testConfigProvider.useValue)
      .overrideProvider(QUEUE_SERVICE)
      .useClass(MockQueueService)
      .compile();

    app = moduleFixture.createNestMicroservice(clientOptions);
    await app.listen();

    clientApp = await NestFactory.create(TestClientModule);
    clientApp.connectMicroservice(clientOptions);
    await clientApp.init();
  });

  afterAll(async () => {
    await app.close();
    await clientApp.close();
  });

  describe('client', () => {
    it('should be able to define service', async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      expect(controller).toBeDefined();
      expect(controller.getClient()).toBeDefined();
      expect(controller.getService()).toBeDefined();
    });

    it('should be able to call rpc Start', async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      const result = await firstValueFrom(
        controller.start({ os: 'linux', url: 'url' })
      );

      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
    });

    it('should be able to call rpc GetResult', async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      const result = await firstValueFrom(controller.getResult({ id: 'id' }));

      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
      expect(result.url).toBeDefined();
      expect(result.stdout).toBeDefined();
      expect(result.stderr).toBeDefined();
    });
  });
});
