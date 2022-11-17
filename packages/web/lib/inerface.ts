export interface Inspection {
  id: string;
  url: string;
  result: {
    url: string;
    malicious: boolean;
    screenshot: string;
    hosts: string[];
    locations: string[];
    sendingTo: string[];
  } | null;
  error: string | null;
  detectedAt: string;
}
