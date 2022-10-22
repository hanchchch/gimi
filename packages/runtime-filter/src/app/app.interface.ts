import { Observable } from 'rxjs';

export interface StartRequest {
  url: string;
}

export interface StartResponse {
  id: string;
}

export interface GetResultRequest {
  id: string;
}

export interface GetResultResponse {
  id: string;
  url: string;
  stdout: string;
  stderr: string;
}

export interface RuntimeFilterService {
  Start(data: StartRequest): Observable<StartResponse>;
  GetResult(data: GetResultRequest): Observable<GetResultResponse>;
}
