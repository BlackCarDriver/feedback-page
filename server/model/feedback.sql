CREATE TABLE public.t_feedback
(
  id  SERIAL primary key ,   
  openid character(100),
  pblocation character(300), -- The location of the problem
  fbtype character(100), -- type of feedback 
  images character(100), -- the name of upload images, not including prefix path
  describe character(500), -- the detail of describe
  fbtime timestamp with time zone DEFAULT CURRENT_TIMESTAMP,-- feedback time
  fbstate integer default 0 -- Have read or not
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.t_feedback
  OWNER TO blackcardriver;
COMMENT ON COLUMN public.t_feedback.pblocation IS 'The location of the problem';
COMMENT ON COLUMN public.t_feedback.fbstate IS 'Have read or not';
COMMENT ON COLUMN public.t_feedback.fbtype IS 'type of feedback';
COMMENT ON COLUMN public.t_feedback.images IS 'the name of upload images, not including prefix path';
COMMENT ON COLUMN public.t_feedback.describe IS 'the detail of describe';
COMMENT ON COLUMN public.t_feedback.fbtime IS 'feedback time';