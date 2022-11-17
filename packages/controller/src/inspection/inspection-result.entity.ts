import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  PrimaryColumn,
} from "typeorm";
import { InspectionResult as IInspectionResult } from "@proto/ts/messages/inspection";

@Entity()
export class InspectionResult extends BaseEntity implements IInspectionResult {
  @PrimaryColumn()
  url: string;

  @Column({ unique: true })
  requestId: string;

  @Column()
  malicious: boolean;

  @Column()
  screenshot: string;

  @Column({ type: "text", array: true })
  locations: string[];

  @Column({ type: "text", array: true })
  hosts: string[];

  @Column({ type: "text", array: true })
  sendingTo: string[];

  @CreateDateColumn({
    transformer: {
      from: (value) => value.toISOString(),
      to: (value) => value,
    },
  })
  detectedAt: string;
}
