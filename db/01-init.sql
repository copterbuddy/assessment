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

-- id 1 use for service get-by-id_it_test
INSERT INTO "expenses" ("title","amount","note","tags") VALUES ('strawberry smoothie', 79, 'night market promotion discount 10 bath', '{"food", "beverage"}');
-- id 2 use for service update-by-id_it_test
INSERT INTO "expenses" ("title","amount","note","tags") VALUES ('strawberry smoothie', 19, 'night market promotion discount 10 bath', '{"food", "beverage"}');
