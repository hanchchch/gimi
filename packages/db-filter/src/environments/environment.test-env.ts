import { ConfigService } from "@nestjs/config";
import { join } from "path";
import { EnvVars } from "./environment.interface";

export const testEnv: EnvVars = {
  DB_URL: "postgres://localhost:6379",
  DB_SYNC: true,
  PORT: 50051,
  PROTO_PATH: join(
    __dirname,
    "..",
    "..",
    "..",
    "packages/proto/messages/runtimefilter.proto"
  ),
};

export const testConfigProvider = {
  provide: ConfigService,
  useValue: {
    get: (key: keyof EnvVars) => testEnv[key],
  },
};
