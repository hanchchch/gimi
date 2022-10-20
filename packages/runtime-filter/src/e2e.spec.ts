import {
  Controller,
  INestApplication,
  INestMicroservice,
  Inject,
  Module,
} from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { ClientGrpc, ClientsModule } from '@nestjs/microservices';
import { grpcClientOptions } from './grpc-client.options';
import { AppModule } from './app/app.module';
import {
  RuntimeFilterRequest,
  RuntimeFilterService,
} from './app/app.interface';
import { firstValueFrom } from 'rxjs';

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

  runtimeFilter(data: RuntimeFilterRequest) {
    return this.service.RuntimeFilter(data);
  }
}

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'RUNTIME_FILTER_PACKAGE',
        ...grpcClientOptions,
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
    app = await NestFactory.createMicroservice(AppModule, grpcClientOptions);
    await app.listen();

    clientApp = await NestFactory.create(TestClientModule);
    clientApp.connectMicroservice(grpcClientOptions);
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

    it('should be able to call rpc', async () => {
      const controller =
        clientApp.get<TestClientController>(TestClientController);

      const result = await firstValueFrom(
        controller.runtimeFilter({ url: 'url' })
      );

      expect(result).toBeDefined();
      expect(result.url).toEqual('url');
    });
  });
});
