-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS expenses_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
    "id" int4 NOT NULL DEFAULT nextval('expenses_id_seq'::regclass),
    "title" TEXT,
    "amount" FLOAT,
    "note" TEXT,
    "tags" TEXT[]
    PRIMARY KEY ("id")
);

-- INSERT INTO "expenses" ("id", "title", "amount", "note","tags") VALUES (1, 'test-title', 'test-content', 'test-author');