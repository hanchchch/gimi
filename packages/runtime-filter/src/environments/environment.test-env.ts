import { ConfigService } from '@nestjs/config';
import { EnvVars } from './environment.interface';

export const testEnv: EnvVars = {
  REDIS_URL: 'redis://localhost:6379',
};

export const testConfigProvider = {
  provide: ConfigService,
  useValue: {
    get: (key: keyof EnvVars) => testEnv[key],
  },
};
