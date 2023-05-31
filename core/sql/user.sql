CREATE TABLE IF NOT EXISTS public.user (
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT Email_UC UNIQUE(email),
    CONSTRAINT Username_UC UNIQUE(username)
);