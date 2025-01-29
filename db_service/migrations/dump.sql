--
-- PostgreSQL database dump
--

-- Dumped from database version 15.10 (Homebrew)
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: krillkovalev
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    task_title character varying(255),
    task_body text,
    done boolean DEFAULT false
);


ALTER TABLE public.tasks OWNER TO krillkovalev;

--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: krillkovalev
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tasks_id_seq OWNER TO krillkovalev;

--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: krillkovalev
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: krillkovalev
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: krillkovalev
--

COPY public.tasks (id, task_title, task_body, done) FROM stdin;
12	kdkgs	dfk,dsk,fdsf	f
15	Сделать что-то	я что-то делаю 	f
6	Записаться к врачу	Прием у терапевта	t
9	Удалить из кеша	Redis тест	t
8	Тест	Тест	t
16	Тест кафки	Тест кафки	f
17	ладвыльбамв	двбаыджамбсывджм	f
18	Логи сделать	тест логов	f
19	авыдащвбп	ыбадвыба	f
\.


--
-- Name: tasks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: krillkovalev
--

SELECT pg_catalog.setval('public.tasks_id_seq', 19, true);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: krillkovalev
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--
