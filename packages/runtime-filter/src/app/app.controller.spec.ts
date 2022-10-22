import { Metadata } from '@grpc/grpc-js';
import { Test, TestingModule } from '@nestjs/testing';
import { testConfigProvider } from '../environments/environment.test-env';

import { AppController } from './app.controller';
import { AppService } from './app.service';

describe('AppController', () => {
  let app: TestingModule;
  let controller: AppController;

  beforeAll(async () => {
    app = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService, testConfigProvider],
    }).compile();

    controller = app.get<AppController>(AppController);
  });

  describe('Start', () => {
    it('should return id', async () => {
      const result = await controller.Start(
        { url: 'url' },
        new Metadata(),
        /* @ts-ignore */
        {}
      );

      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
    });
  });

  describe('GetResult', () => {
    it('should return ...', async () => {
      const result = await controller.GetResult(
        { id: 'id' },
        new Metadata(),
        /* @ts-ignore */
        {}
      );

      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
      expect(result.url).toBeDefined();
      expect(result.stdout).toBeDefined();
      expect(result.stderr).toBeDefined();
    });
  });
});
