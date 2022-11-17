import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  PrimaryColumn,
} from "typeorm";
import { InspectionResult } from "@proto/ts/messages/inspection";

@Entity()
export class Inspection extends BaseEntity {
  @PrimaryColumn()
  id: string;

  @Column()
  url: string;

  @Column({ type: "json", nullable: true })
  result: InspectionResult | null;

  @Column({ nullable: true })
  error: string | null;

  @CreateDateColumn()
  detectedAt: Date;
}
