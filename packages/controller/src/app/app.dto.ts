export class InspectDto {
  url: string;
  os: string;
  timeout?: number;
}

export interface InspectParams {
  url: string;
  os: string;
  timeout?: number;
}
