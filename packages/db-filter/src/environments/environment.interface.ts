export interface EnvVars {
  PROTO_PATH: string;
  PORT: number;
  DB_TYPE: "postgres" | "sqlite";
  DB_SYNC: boolean;
  DB_URL: string;
}
