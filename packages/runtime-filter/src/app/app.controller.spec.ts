import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';
import { Test, TestingModule } from '@nestjs/testing';

import { AppController } from './app.controller';
import { AppService } from './app.service';

describe('AppController', () => {
  let app: TestingModule;
  let controller: AppController;

  beforeAll(async () => {
    app = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService],
    }).compile();

    controller = app.get<AppController>(AppController);
  });

  describe('getData', () => {
    it('should return ""', () => {
      const result = controller.RuntimeFilter(
        { url: 'url' },
        new Metadata(),
        /* @ts-ignore */
        {}
      );

      expect(result).toBeDefined();
      expect(result.url).toEqual('url');
    });
  });
});
