import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MapsModule } from './maps/maps.module';
import { RoutesModule } from './routes/routes.module';
import { PrismaModule } from './prisma/prisma.module';

@Module({
  imports: [
    // add config module from nestjs to be able to inject environment variables
    ConfigModule.forRoot({
      isGlobal: true
    }),
    PrismaModule,
    MapsModule,
    RoutesModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
