-- migrate:up
CREATE SEQUENCE IF NOT EXISTS loan_investor_sequence
            START WITH 1
            INCREMENT BY 1
            NO MINVALUE
            NO MAXVALUE
            CACHE 1;
CREATE TABLE loan_investor (
    id BIGINT DEFAULT nextval('loan_investor_sequence'::regclass) NOT NULL PRIMARY KEY,
    loan_id BIGINT NOT NULL,
    investor_id BIGINT NOT NULL,
    amount_invested BIGINT NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE loan_investor;
DROP SEQUENCE loan_investor_sequence;
