


CREATE SEQUENCE public.beer_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;


CREATE TABLE public.beer
(
    beer_id integer NOT NULL DEFAULT nextval('beer_id_seq'::regclass),
    beer_name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    graduation real,
    CONSTRAINT beer_pkey PRIMARY KEY (beer_name)
);
