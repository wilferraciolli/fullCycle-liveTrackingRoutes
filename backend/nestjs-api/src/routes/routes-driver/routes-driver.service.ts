import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma/prisma.service';

@Injectable()
export class RoutesDriverService {
  constructor(
    private prismaService: PrismaService
  ) {
  }

  public processRoute(dto: { route_id: string; lat: number; lng: number }) {
    // add points or create route for driver
    return this.prismaService.routeDriver.upsert({
      include: {
        route: true //eager loading
      },
      where: { route_id: dto.route_id },
      create: {
        route_id: dto.route_id,
        points: {
          set: {
            location: {
              lat: dto.lat,
              lng: dto.lng
            }
          }
        }
      },
      update: {
        points: {
          push: {
            location: {
              lat: dto.lat,
              lng: dto.lng
            }
          }
        }
      }
    });
  }
}
