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
