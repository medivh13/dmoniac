CREATE SEQUENCE article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE article_id_seq OWNED BY article.id;

ALTER TABLE ONLY article ALTER COLUMN id SET DEFAULT nextval('article_id_seq'::regclass);

ALTER TABLE ONLY article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);