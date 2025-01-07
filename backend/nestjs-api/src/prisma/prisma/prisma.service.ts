import { Injectable, OnModuleInit } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';


// service used to run on start up, this is to connect to the database as the module is initialized
@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit {
   async onModuleInit() {
    await this.$connect();
  }
}
