import { Inject } from '@nestjs/common';
import { RuntimeFilterService } from './runtime-filter.service';

export const InjectRuntimeFilter = () => Inject(RuntimeFilterService);
