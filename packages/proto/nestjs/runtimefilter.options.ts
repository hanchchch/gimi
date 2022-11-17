import { Transport, GrpcOptions } from "@nestjs/microservices";
import { join } from "path";

export const createClientOptions = ({
  url = "0.0.0.0:50051",
}: {
  url?: string;
}): GrpcOptions => ({
  transport: Transport.GRPC,
  options: {
    package: "runtimefilter",
    url,
    protoPath: [join(
      __dirname,
      "..",
      "..",
      "..",
      "packages/proto/messages/runtimefilter.proto"
    ), join(
      __dirname,
      "..",
      "..",
      "..",
      "packages/proto/messages/inspection.proto"
    )],
  },
});
