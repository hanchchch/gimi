import { Inject, Injectable } from "@nestjs/common";
import { ClientGrpc } from "@nestjs/microservices";
import { IDbFilterService } from "@proto/nestjs/dbfilter.interface";
import { FindRequest } from "@proto/ts/messages/dbfilter";
import { DB_FILTER_PACKAGE } from "./db-filter.symbols";

@Injectable()
export class DbFilterService {
  private service: IDbFilterService;

  constructor(@Inject(DB_FILTER_PACKAGE) private readonly client: ClientGrpc) {
    this.service = this.client.getService<IDbFilterService>("DbFilterService");
  }

  find(data: FindRequest) {
    return this.service.Find(data);
  }
}
