CREATE TABLE IF NOT EXISTS public.point_stage (
    point_id SERIAL REFERENCES public.point (id),
    stage_id SERIAL REFERENCES public.stage (id),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
)