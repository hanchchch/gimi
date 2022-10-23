import { Test, TestingModule } from '@nestjs/testing';
import { mockRuntimeFilterServiceProvider } from '../runtime-filter/mock-runtime-filter.provider';

import { AppController } from './app.controller';
import { AppService } from './app.service';

describe('AppController', () => {
  let app: TestingModule;

  beforeAll(async () => {
    app = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService, mockRuntimeFilterServiceProvider],
    }).compile();
  });

  describe('getData', () => {
    it('should be defined', () => {
      const appController = app.get<AppController>(AppController);
      expect(appController).toBeDefined();
    });
  });
});
