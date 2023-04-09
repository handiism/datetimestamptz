import { PrismaClient, datetimestamptz } from "@prisma/client";

export async function main() {
  const prisma = new PrismaClient({
    errorFormat: "minimal",
  });

  const now = new Date("2023-04-04T10:59:57Z");

  const all = await prisma.datetimestamptz.create({
    data: {
      date: now,
      time: now,
      timetz: now,
      timestamp: now,
      timestamptz: now,
    },
  });

  console.log(all.timestamp?.toLocaleString());

  type dts = datetimestamptz & {
    interval: string | null;
  };

  const allRaw =
    await prisma.$queryRaw<dts>`INSERT INTO datetimestamptz(date,time,timetz,timestamp,timestamptz,interval) VALUES(${now},${now},${now},${now},${now},${"1 days 1 hours 1 minutes"}::text::interval) RETURNING id,date,time,timetz,timestamp,timestamptz,interval::text`;
  console.log(allRaw);
}

main();
