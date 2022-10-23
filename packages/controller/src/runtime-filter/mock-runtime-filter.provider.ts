import { MockRuntimeFilterService } from './mock-runtime-filter.service';
import { RuntimeFilterService } from './runtime-filter.service';

export const mockRuntimeFilterServiceProvider = {
  provide: RuntimeFilterService,
  useClass: MockRuntimeFilterService,
};
