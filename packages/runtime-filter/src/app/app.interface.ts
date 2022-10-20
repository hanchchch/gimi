import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';
import { Observable } from 'rxjs';

export interface RuntimeFilterRequest {
  url: string;
}

export interface RuntimeFilterResponse {
  url: string;
  stdout: string;
  stderr: string;
}

export interface RuntimeFilterService {
  RuntimeFilter(data: RuntimeFilterRequest): Observable<RuntimeFilterResponse>;
}
