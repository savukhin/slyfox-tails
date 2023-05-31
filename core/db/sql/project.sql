CREATE TABLE IF NOT EXISTS public.project (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    creator_id SERIAL REFERENCES public.user (id),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT Email_UC UNIQUE(email),
    CONSTRAINT Username_UC UNIQUE(username)
);