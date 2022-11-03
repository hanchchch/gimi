import { BaseEntity, CreateDateColumn, Entity, PrimaryColumn } from "typeorm";
import { Blacklist as IBlacklist } from "@proto/ts/messages/dbfilter";

@Entity()
export class Blacklist extends BaseEntity implements IBlacklist {
  @PrimaryColumn()
  url: string;

  @CreateDateColumn({
    transformer: {
      from: (value) => value.toISOString(),
      to: (value) => value,
    },
  })
  detectedAt: string;
}
