-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS expenses_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
    "id" int4 NOT NULL DEFAULT nextval('expenses_id_seq'::regclass),
    "title" TEXT,
    "amount" DOUBLE PRECISION,
    "note" TEXT,
    "tags" TEXT[],
    PRIMARY KEY ("id")
);
