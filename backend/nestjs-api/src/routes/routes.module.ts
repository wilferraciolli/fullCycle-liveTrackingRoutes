import { Module } from '@nestjs/common';
import { KafkaModule } from '../kafka/kafka.module';
import { MapsModule } from '../maps/maps.module';
import { RoutesController } from './routes.controller';
import { RoutesService } from './routes.service';
import { RoutesDriverService } from './routes-driver/routes-driver.service';
import { RoutesDriverGateway } from './routes-driver/routes-driver.gateway';

@Module({
  imports: [MapsModule, KafkaModule],
  controllers: [RoutesController],
  providers: [RoutesService, RoutesDriverService, RoutesDriverGateway]
})
export class RoutesModule {
}
