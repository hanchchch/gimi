import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ClientsModule } from '@nestjs/microservices';
import { createClientOptions } from '@proto/nestjs/runtimefilter.options';
import { EnvVars } from '../environments/environment.interface';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { join } from 'path';

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    ClientsModule.registerAsync([
      {
        name: 'RUNTIME_FILTER_PACKAGE',
        useFactory: (config: ConfigService<EnvVars>) =>
          createClientOptions({
            url: config.get('RUNTIME_FILTER_URL'),
          }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
