# Datetimestamptz

## Go PostgresQL `jackc/pg`

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
| Interval      | ❌           | Date         |

Remember to store date always in UTC format.
Example:

```js
new Date("2023-04-04T10:59:57Z"); // Remember that this is UTC-0
new Date(); // Stored in UTC-0
```
