CREATE TABLE IF NOT EXISTS public.point (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    creator_id SERIAL REFERENCES public.user (id),

    password_hash VARCHAR(100);

    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
)