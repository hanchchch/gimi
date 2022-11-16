import {
  GetResultRequest,
  GetResultResponse,
  StartRequest,
  StartResponse,
} from "@proto/ts/messages/runtimefilter";
import { Observable } from "rxjs";

export interface IRuntimeFilterService {
  Start(data: StartRequest): Observable<StartResponse>;
  GetResult(data: GetResultRequest): Observable<GetResultResponse>;
}
