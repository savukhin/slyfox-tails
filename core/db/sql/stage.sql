CREATE TABLE IF NOT EXISTS public.stage (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    creator_id SERIAL REFERENCES public.user (id),
    job_id SERIAL REFERENCES public.project (id),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
)