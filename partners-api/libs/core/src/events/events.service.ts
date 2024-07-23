import { Injectable } from '@nestjs/common';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { ReserveSpotDto } from './dto/reserve-spot.dto';
import { Prisma, SpotStatus, TicketStatus } from '@prisma/client';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class EventsService {
  constructor(private prismaService: PrismaService) {}

  create(createEventDto: CreateEventDto) {
    return this.prismaService.event.create({
      data: {
        ...createEventDto,
        date: new Date(createEventDto.date),
      },
    });
  }

  findAll() {
    return this.prismaService.event.findMany();
  }

  findOne(id: string) {
    return this.prismaService.event.findUnique({
      where: {
        id,
      },
    });
  }

  update(id: string, updateEventDto: UpdateEventDto) {
    return this.prismaService.event.update({
      data: {
        ...updateEventDto,
        date: new Date(updateEventDto.date),
      },
      where: {
        id,
      },
    });
  }

  remove(id: string) {
    return this.prismaService.event.delete({
      where: {
        id,
      },
    });
  }

  async reserverSpot(dto: ReserveSpotDto) {
    console.log(dto)
    const spots = await this.prismaService.spot.findMany({
      where: {
        name: {
          in: dto.spots
        },
        eventId: dto.eventId
      }
    })

    if(spots.length !== dto.spots.length){
      const foundSpotsName = spots.map((spot)=> spot.name)
      const notFoundSpotsName = dto.spots.filter( (spotName) => !foundSpotsName.includes(spotName))
      throw new Error(`Spots ${notFoundSpotsName.join(', ')} not found`)
    }

    try{
      const tickets = await this.prismaService.$transaction(async (prisma)=> {
        await prisma.reservationHistory.createMany({
          data: spots.map((spot) =>({
            spotId: spot.id,
            tickeKind: dto.ticket_kind,
            email: dto.email,
            status: TicketStatus.reserved
          }))
        })
    
        await prisma.spot.updateMany({
          where: {
            id: {
              in: spots.map((spot)=>spot.id)
            }
          },
          data:{
            status: SpotStatus.reserved
          }
        })
    
        const tickets = await Promise.all(
          spots.map((spot)=> prisma.ticket.create({
            data:{
              spotId: spot.id,
              tickeKind: dto.ticket_kind,
              email: dto.email
            }
          }))
        )
        return tickets
      }, { isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted })
      return tickets
    }
    catch(err){
      if(err instanceof Prisma.PrismaClientKnownRequestError){
        switch(err.code){
          case 'P2002' : 
          case 'P2034' : 
            throw new Error('Some spots are already reserved');
        }
      }
      throw err
    }
    
   
  }
}
