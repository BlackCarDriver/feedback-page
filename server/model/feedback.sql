-- Table: public.t_feedback2

-- DROP TABLE public.t_feedback2;

CREATE TABLE public.t_feedback2
(
    id integer NOT NULL DEFAULT nextval('t_feedback2_id_seq'::regclass),
    user_id character varying COLLATE pg_catalog."default" DEFAULT 'anonymity'::character varying,
    email character(100) COLLATE pg_catalog."default" DEFAULT 'unset'::bpchar,
    fb_type character(100) COLLATE pg_catalog."default" NOT NULL,
    fb_location character(300) COLLATE pg_catalog."default" NOT NULL,
    images_name character(200) COLLATE pg_catalog."default",
    describe character(500) COLLATE pg_catalog."default" NOT NULL,
    fb_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fb_status integer NOT NULL DEFAULT 0,
    CONSTRAINT t_feedback2_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.t_feedback2
    OWNER to blackcardriver;

COMMENT ON COLUMN public.t_feedback2.user_id
    IS 'user''s id';

COMMENT ON COLUMN public.t_feedback2.email
    IS 'user''s email';

COMMENT ON COLUMN public.t_feedback2.fb_type
    IS 'feedback''s type';

COMMENT ON COLUMN public.t_feedback2.fb_location
    IS 'problem location';

COMMENT ON COLUMN public.t_feedback2.images_name
    IS 'name of  iamges';

COMMENT ON COLUMN public.t_feedback2.describe
    IS 'describe of the problem';

COMMENT ON COLUMN public.t_feedback2.fb_time
    IS 'when the problem happen';

COMMENT ON COLUMN public.t_feedback2.fb_status
    IS 'feedback status, 0 mean not read, 1 mean have been read';