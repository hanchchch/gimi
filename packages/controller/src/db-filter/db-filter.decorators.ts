import { Inject } from "@nestjs/common";
import { DbFilterService } from "./db-filter.service";

export const InjectDbFilter = () => Inject(DbFilterService);
