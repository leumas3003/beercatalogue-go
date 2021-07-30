# beercatalogue-go
Api for a beer catalogue. 
Two method, addbeer, getbeer.
Access to localhost:3001/swagger/index.html

## Installation and SetUp
 1. Download and install Golang. 
 2. Donwload and install PostgreSQL. 
 3. On PostgreSQL up and running use this script to create and add a new table into a database

```sql
CREATE SEQUENCE public.beer_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.beer_id_seq
    OWNER TO samuelmartel;

CREATE TABLE IF NOT EXISTS public.beer
(
    beer_id integer NOT NULL DEFAULT nextval('beer_id_seq'::regclass),
    beer_name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    graduation real,
    CONSTRAINT beer_pkey PRIMARY KEY (beer_name)
)

TABLESPACE pg_default;

ALTER TABLE public.beer
    OWNER to samuelmartel;
```

 4. Replace in main.go ln 27 with your PostgreSQL data
```go
dbClient, err = sql.Open("postgres", fmt.Sprintf(DbConn, "localhost", "5432", "USERNAME", "", "DATABASENAME"))
```

## Package Used

| Function| Package|
|-------|--|
| HTTP  | https://echo.labstack.com/ |
| POSTGRES | https://github.com/jackc/pgx/v4 |
| Swagger | https://github.com/swaggo/echo-swagger |

## License

MIT License

Copyright (c) [2021] [Samuel Martel]

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
