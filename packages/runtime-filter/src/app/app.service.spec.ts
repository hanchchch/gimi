import { Test } from '@nestjs/testing';
import { testConfigProvider } from '../environments/environment.test-env';

import { AppService } from './app.service';

describe('AppService', () => {
  let service: AppService;

  beforeAll(async () => {
    const app = await Test.createTestingModule({
      providers: [AppService, testConfigProvider],
    }).compile();

    service = app.get<AppService>(AppService);
  });

  describe('start', () => {
    it('should return request id', async () => {
      const result = await service.start({ url: 'url' });
      expect(result).toBeDefined();
      expect(result.id).toBeDefined();
    });
  });
});
