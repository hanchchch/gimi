import { FindRequest, FindResponse } from "@proto/ts/messages/dbfilter";
import { Observable } from "rxjs";

export interface IDbFilterService {
  Find(data: FindRequest): Observable<FindResponse>;
}
