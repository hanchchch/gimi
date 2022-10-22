import { Observable } from 'rxjs';

export interface StartRequest {
  os: string;
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

export interface HandlerArgs {
  request_id: string;
  inspection_args: {
    url: string;
  };
}

export interface HandlerResult {
  request_id: string;
  inspection_result: {
    url: string;
    stdout: string;
    stderr: string;
  };
}
