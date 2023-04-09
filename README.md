# Datetimestamptz

## Go PostgresQL `jackc/pgx`

| Database Type | Availability | Program Type       |
| ------------- | ------------ | ------------------ |
| Time          | ✅           | pgtype.Time        |
| Timetz        | ❌           |                    |
| Timestamp     | ✅           | pgtype.Timestamp   |
| Timestamptz   | ✅           | pgtype.Timestamptz |
| Date          | ✅           | pgtype.Date        |
| Interval      | ✅           | pgtype.Interval    |

## Node PostgresQL `prisma`

| Database Type | Availability | Program Type |
| ------------- | ------------ | ------------ |
| Time          | ✅           | Date         |
| Timetz        | ✅           |              |
| Timestamp     | ✅           | Date         |
| Timestamptz   | ✅           | Date         |
| Date          | ✅           | Date         |
| Interval      | ✅           | String       |

Prisma, at the moment did not provide an API for interval data type. So it must be casted into string (since Prisma did not provide deserialized from interval) if we need to use the value in your code.

Remember to store date always in UTC format.
Example:

```js
new Date("2023-04-04T10:59:57Z"); // Remember that this is UTC-0
new Date(); // Stored in UTC-0
```
