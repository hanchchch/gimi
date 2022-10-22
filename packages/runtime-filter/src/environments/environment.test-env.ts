import { ConfigService } from '@nestjs/config';
import { join } from 'path';
import { EnvVars } from './environment.interface';

export const testEnv: EnvVars = {
  REDIS_URL: 'redis://localhost:6379',
  PORT: 50051,
  PROTO_PATH: join(
    __dirname,
    '..',
    '..',
    '..',
    'packages/proto/messages/runtimefilter.proto'
  ),
};

export const testConfigProvider = {
  provide: ConfigService,
  useValue: {
    get: (key: keyof EnvVars) => testEnv[key],
  },
};
