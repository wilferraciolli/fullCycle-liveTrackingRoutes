import { NestFactory } from '@nestjs/core';
import { KafkaServer } from '../kafka/kafka-server';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { AppModule } from '../app.module';

async function bootstrap() {
  const appConfigContext =
    await NestFactory.createApplicationContext(ConfigModule);
  const configService = appConfigContext.get(ConfigService);

  // create a new app and set its config
  const app = await NestFactory.createMicroservice(AppModule, {
    strategy: new KafkaServer({
      server: {
        'bootstrap.servers': configService.get('KAFKA_BROKER'),
      },
      consumer: {
        'group.id': 'nest-group',
        'client.id': `nest-group-${configService.get('HOSTNAME')}`,
        'max.poll.interval.ms': 10000,
        'session.timeout.ms': 10000,
      },
    }),
  });
  await app.listen();
}
bootstrap();
