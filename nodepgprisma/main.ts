import { PrismaClient } from "@prisma/client";

export async function main() {
  const prisma = new PrismaClient({
    errorFormat: "minimal",
  });

  const now = new Date("2023-04-04T10:59:57Z");

  const all = await prisma.datetimestamptz.create({
    data: {
      date: now,
      time: now,
      timestamp: now,
      timestamptz: now,
      timetz: now,
    },
  });

  console.log(all.timestamp?.toLocaleString());
}

main();
