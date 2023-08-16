--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg110+1)
-- Dumped by pg_dump version 15.4 (Debian 15.4-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: categories; Type: TABLE; Schema: public; Owner: recipe
--

CREATE TABLE public.categories (
    id bytea NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO recipe;

--
-- Name: materials; Type: TABLE; Schema: public; Owner: recipe
--

CREATE TABLE public.materials (
    id bytea NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.materials OWNER TO recipe;

--
-- Name: recipe_materials; Type: TABLE; Schema: public; Owner: recipe
--

CREATE TABLE public.recipe_materials (
    recipe_id bytea NOT NULL,
    material_id bytea NOT NULL,
    quantity bigint,
    unit_id bytea NOT NULL,
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.recipe_materials OWNER TO recipe;

--
-- Name: recipe_materials_id_seq; Type: SEQUENCE; Schema: public; Owner: recipe
--

CREATE SEQUENCE public.recipe_materials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_materials_id_seq OWNER TO recipe;

--
-- Name: recipe_materials_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: recipe
--

ALTER SEQUENCE public.recipe_materials_id_seq OWNED BY public.recipe_materials.id;


--
-- Name: recipes; Type: TABLE; Schema: public; Owner: recipe
--

CREATE TABLE public.recipes (
    id bytea NOT NULL,
    name text,
    category_id bytea,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.recipes OWNER TO recipe;

--
-- Name: units; Type: TABLE; Schema: public; Owner: recipe
--

CREATE TABLE public.units (
    id bytea NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.units OWNER TO recipe;

--
-- Name: recipe_materials id; Type: DEFAULT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipe_materials ALTER COLUMN id SET DEFAULT nextval('public.recipe_materials_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: recipe
--

COPY public.categories (id, name, created_at, updated_at, deleted_at) FROM stdin;
\\x0189fb14508a7ca7e5475ac42dc9e61f	Main course	2023-08-15 21:22:39.882163+00	2023-08-15 21:22:39.882631+00	\N
\.


--
-- Data for Name: materials; Type: TABLE DATA; Schema: public; Owner: recipe
--

COPY public.materials (id, name, created_at, updated_at, deleted_at) FROM stdin;
\\x0189fb1455c0be924aa39c2b62ad96a6	Nasi	2023-08-15 21:22:41.216441+00	2023-08-15 21:22:41.216751+00	\N
\\x0189fb5ba05076706741a7acfe2c7073	Garam	2023-08-15 22:40:33.360114+00	2023-08-15 22:40:33.361523+00	\N
\.


--
-- Data for Name: recipe_materials; Type: TABLE DATA; Schema: public; Owner: recipe
--

COPY public.recipe_materials (recipe_id, material_id, quantity, unit_id, id, created_at, updated_at, deleted_at) FROM stdin;
\\x0189fc8d7c71d36dced02931568b3ee0	\\x0189fb1455c0be924aa39c2b62ad96a6	1	\\x0189fb144a13161ffb759f3928d8ef6e	5	\N	\N	\N
\\x0189fc8d7c71d36dced02931568b3ee0	\\x0189fb5ba05076706741a7acfe2c7073	2	\\x0189fb5a8b6ac52cf9cb198c9939314f	6	\N	\N	\N
\\x0189fcb8c2a32736015a386767e86a06	\\x0189fb1455c0be924aa39c2b62ad96a6	2	\\x0189fb144a13161ffb759f3928d8ef6e	7	\N	\N	\N
\.


--
-- Data for Name: recipes; Type: TABLE DATA; Schema: public; Owner: recipe
--

COPY public.recipes (id, name, category_id, created_at, updated_at, deleted_at) FROM stdin;
\\x0189fc8d7c71d36dced02931568b3ee0	Nasi goreng	\\x0189fb14508a7ca7e5475ac42dc9e61f	2023-08-16 04:14:38.193525+00	2023-08-16 04:15:03.189442+00	\N
\\x0189fcb8c2a32736015a386767e86a06	Nasi timbel	\\x0189fb14508a7ca7e5475ac42dc9e61f	2023-08-16 05:01:54.211921+00	2023-08-16 05:01:54.214731+00	\N
\.


--
-- Data for Name: units; Type: TABLE DATA; Schema: public; Owner: recipe
--

COPY public.units (id, name, created_at, updated_at, deleted_at) FROM stdin;
\\x0189fb144a13161ffb759f3928d8ef6e	Plate	2023-08-15 21:22:38.227866+00	2023-08-15 21:22:38.228528+00	\N
\\x0189fb5a8b6ac52cf9cb198c9939314f	Tablespoon	2023-08-15 22:39:22.474573+00	2023-08-15 22:39:22.476707+00	\N
\.


--
-- Name: recipe_materials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: recipe
--

SELECT pg_catalog.setval('public.recipe_materials_id_seq', 7, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: materials materials_pkey; Type: CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.materials
    ADD CONSTRAINT materials_pkey PRIMARY KEY (id);


--
-- Name: recipe_materials recipe_materials_pkey; Type: CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipe_materials
    ADD CONSTRAINT recipe_materials_pkey PRIMARY KEY (recipe_id, material_id, unit_id);


--
-- Name: recipes recipes_pkey; Type: CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT recipes_pkey PRIMARY KEY (id);


--
-- Name: units units_pkey; Type: CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: recipe
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_materials_deleted_at; Type: INDEX; Schema: public; Owner: recipe
--

CREATE INDEX idx_materials_deleted_at ON public.materials USING btree (deleted_at);


--
-- Name: idx_recipe_materials_deleted_at; Type: INDEX; Schema: public; Owner: recipe
--

CREATE INDEX idx_recipe_materials_deleted_at ON public.recipe_materials USING btree (deleted_at);


--
-- Name: idx_recipes_deleted_at; Type: INDEX; Schema: public; Owner: recipe
--

CREATE INDEX idx_recipes_deleted_at ON public.recipes USING btree (deleted_at);


--
-- Name: idx_units_deleted_at; Type: INDEX; Schema: public; Owner: recipe
--

CREATE INDEX idx_units_deleted_at ON public.units USING btree (deleted_at);


--
-- Name: recipe_materials fk_materials_recipes; Type: FK CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipe_materials
    ADD CONSTRAINT fk_materials_recipes FOREIGN KEY (material_id) REFERENCES public.materials(id);


--
-- Name: recipe_materials fk_recipe_materials_unit; Type: FK CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipe_materials
    ADD CONSTRAINT fk_recipe_materials_unit FOREIGN KEY (unit_id) REFERENCES public.units(id);


--
-- Name: recipes fk_recipes_category; Type: FK CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT fk_recipes_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: recipe_materials fk_recipes_materials; Type: FK CONSTRAINT; Schema: public; Owner: recipe
--

ALTER TABLE ONLY public.recipe_materials
    ADD CONSTRAINT fk_recipes_materials FOREIGN KEY (recipe_id) REFERENCES public.recipes(id);


--
-- PostgreSQL database dump complete
--

