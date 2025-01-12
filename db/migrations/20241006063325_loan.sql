-- migrate:up
CREATE SEQUENCE IF NOT EXISTS loan_sequence
            START WITH 1
            INCREMENT BY 1
            NO MINVALUE
            NO MAXVALUE
            CACHE 1;
CREATE TABLE loan (
    id BIGINT DEFAULT nextval('loan_sequence'::regclass) NOT NULL PRIMARY KEY,
    unique_id VARCHAR(200),
    borrower_id BIGINT NOT NULL,
    principal_amount BIGINT NOT NULL,
    rate BIGINT,
    roi BIGINT NOT NULL,
    current_status VARCHAR(10) NOT NULL,
    proof_image BYTEA,
    approved_by BIGINT,
    approved_date TIMESTAMP WITH TIME ZONE,
    invested_date TIMESTAMP WITH TIME ZONE,
    disbursed_by BIGINT,
    disbursed_date DATE,
    signed_agreement_letter VARCHAR(200),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE loan;
DROP SEQUENCE loan_sequence;
