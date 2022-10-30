import { Transport, GrpcOptions } from "@nestjs/microservices";
import { join } from "path";

export const createClientOptions = ({
  url = "0.0.0.0:50052",
  protoPath = join(
    __dirname,
    "..",
    "..",
    "..",
    "packages/proto/messages/dbfilter.proto"
  ),
}: {
  url?: string;
  protoPath?: string;
}): GrpcOptions => ({
  transport: Transport.GRPC,
  options: {
    package: "dbfilter",
    url,
    protoPath,
  },
});
