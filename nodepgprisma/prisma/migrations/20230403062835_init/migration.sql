-- CreateTable
CREATE TABLE "datetimestamptz" (
    "id" SERIAL NOT NULL,
    "time" TIME(6),
    "timetz" TIMETZ(6),
    "timestamp" TIMESTAMP(6),
    "timestamptz" TIMESTAMPTZ(6),
    "date" DATE,
    "interval" interval,

    CONSTRAINT "datetimestamptz_pkey" PRIMARY KEY ("id")
);
