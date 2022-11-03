import { BaseEntity, CreateDateColumn, Entity, PrimaryColumn } from "typeorm";

@Entity()
export class Blacklist extends BaseEntity {
  @PrimaryColumn()
  url: string;

  @CreateDateColumn()
  detectedAt: Date;
}
