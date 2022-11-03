import { Injectable } from "@nestjs/common";
import { InjectRepository } from "@nestjs/typeorm";
import { FindRequest } from "@proto/ts/messages/dbfilter";
import { Repository } from "typeorm";
import { Blacklist } from "./blacklist.entity";

@Injectable()
export class AppService {
  constructor(
    @InjectRepository(Blacklist)
    private readonly blacklistRepository: Repository<Blacklist>
  ) {}

  async find(params: FindRequest) {
    const { url } = params;
    const blacklist = await this.blacklistRepository.findOne({
      where: { url },
    });

    if (!blacklist) {
      return { found: false };
    }

    return { blacklist, found: true };
  }
}
