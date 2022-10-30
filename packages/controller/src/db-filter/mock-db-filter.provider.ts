import { MockDbFilterService } from "./mock-db-filter.service";
import { DbFilterService } from "./db-filter.service";

export const mockDbFilterServiceProvider = {
  provide: DbFilterService,
  useClass: MockDbFilterService,
};
