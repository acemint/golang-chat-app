CREATE TABLE ca_member (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    age INT,
    gender VARCHAR(10),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE ca_transaction (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    member_id_sender UUID NOT NULL,
    member_id_receiver UUID NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE NOT NULL,
    amount INT NOT NULL
);
