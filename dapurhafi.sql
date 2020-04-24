--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)

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

ALTER TABLE IF EXISTS ONLY public.menuprices DROP CONSTRAINT IF EXISTS menuprices_menu_id_menus_id_foreign;
ALTER TABLE IF EXISTS ONLY public.menuprices DROP CONSTRAINT IF EXISTS menuprices_campaign_id_campaigns_id_foreign;
ALTER TABLE IF EXISTS ONLY public.menupics DROP CONSTRAINT IF EXISTS menupics_menu_id_menus_id_foreign;
DROP INDEX IF EXISTS public.idx_users_deleted_at;
DROP INDEX IF EXISTS public.idx_menus_deleted_at;
DROP INDEX IF EXISTS public.idx_menuprices_deleted_at;
DROP INDEX IF EXISTS public.idx_menupics_deleted_at;
DROP INDEX IF EXISTS public.idx_campaigns_deleted_at;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_pkey;
ALTER TABLE IF EXISTS ONLY public.menus DROP CONSTRAINT IF EXISTS menus_pkey;
ALTER TABLE IF EXISTS ONLY public.menuprices DROP CONSTRAINT IF EXISTS menuprices_pkey;
ALTER TABLE IF EXISTS ONLY public.menupics DROP CONSTRAINT IF EXISTS menupics_pkey;
ALTER TABLE IF EXISTS ONLY public.campaigns DROP CONSTRAINT IF EXISTS campaigns_pkey;
ALTER TABLE IF EXISTS public.users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.menus ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.menuprices ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.menupics ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.campaigns ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.users_id_seq;
DROP TABLE IF EXISTS public.users;
DROP SEQUENCE IF EXISTS public.menus_id_seq;
DROP TABLE IF EXISTS public.menus;
DROP SEQUENCE IF EXISTS public.menuprices_id_seq;
DROP TABLE IF EXISTS public.menuprices;
DROP SEQUENCE IF EXISTS public.menupics_id_seq;
DROP TABLE IF EXISTS public.menupics;
DROP SEQUENCE IF EXISTS public.campaigns_id_seq;
DROP TABLE IF EXISTS public.campaigns;
DROP EXTENSION IF EXISTS plpgsql;
DROP SCHEMA IF EXISTS public;
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: campaigns; Type: TABLE; Schema: public; Owner: dapurhafi
--

CREATE TABLE public.campaigns (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(128) NOT NULL,
    description character varying(512) NOT NULL
);


ALTER TABLE public.campaigns OWNER TO dapurhafi;

--
-- Name: campaigns_id_seq; Type: SEQUENCE; Schema: public; Owner: dapurhafi
--

CREATE SEQUENCE public.campaigns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.campaigns_id_seq OWNER TO dapurhafi;

--
-- Name: campaigns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dapurhafi
--

ALTER SEQUENCE public.campaigns_id_seq OWNED BY public.campaigns.id;


--
-- Name: menupics; Type: TABLE; Schema: public; Owner: dapurhafi
--

CREATE TABLE public.menupics (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    menu_id integer NOT NULL,
    filename character varying(255)
);


ALTER TABLE public.menupics OWNER TO dapurhafi;

--
-- Name: menupics_id_seq; Type: SEQUENCE; Schema: public; Owner: dapurhafi
--

CREATE SEQUENCE public.menupics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menupics_id_seq OWNER TO dapurhafi;

--
-- Name: menupics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dapurhafi
--

ALTER SEQUENCE public.menupics_id_seq OWNED BY public.menupics.id;


--
-- Name: menuprices; Type: TABLE; Schema: public; Owner: dapurhafi
--

CREATE TABLE public.menuprices (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    campaign_id integer NOT NULL,
    menu_id integer NOT NULL,
    price bigint NOT NULL
);


ALTER TABLE public.menuprices OWNER TO dapurhafi;

--
-- Name: menuprices_id_seq; Type: SEQUENCE; Schema: public; Owner: dapurhafi
--

CREATE SEQUENCE public.menuprices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menuprices_id_seq OWNER TO dapurhafi;

--
-- Name: menuprices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dapurhafi
--

ALTER SEQUENCE public.menuprices_id_seq OWNED BY public.menuprices.id;


--
-- Name: menus; Type: TABLE; Schema: public; Owner: dapurhafi
--

CREATE TABLE public.menus (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(128) NOT NULL,
    description character varying(512) NOT NULL,
    tags character varying(255),
    order_count integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.menus OWNER TO dapurhafi;

--
-- Name: menus_id_seq; Type: SEQUENCE; Schema: public; Owner: dapurhafi
--

CREATE SEQUENCE public.menus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menus_id_seq OWNER TO dapurhafi;

--
-- Name: menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dapurhafi
--

ALTER SEQUENCE public.menus_id_seq OWNED BY public.menus.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: dapurhafi
--

CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    fullname character varying(128) NOT NULL,
    email character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(64) NOT NULL,
    verification_token character varying(255),
    verification_expired timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    forgot_token character varying(255),
    forgot_expired timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    access_token character varying(255),
    access_expired timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    profile_image character varying(128),
    facebook_id character varying(16),
    google_id character varying(16)
);


ALTER TABLE public.users OWNER TO dapurhafi;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: dapurhafi
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO dapurhafi;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dapurhafi
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: campaigns id; Type: DEFAULT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.campaigns ALTER COLUMN id SET DEFAULT nextval('public.campaigns_id_seq'::regclass);


--
-- Name: menupics id; Type: DEFAULT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menupics ALTER COLUMN id SET DEFAULT nextval('public.menupics_id_seq'::regclass);


--
-- Name: menuprices id; Type: DEFAULT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menuprices ALTER COLUMN id SET DEFAULT nextval('public.menuprices_id_seq'::regclass);


--
-- Name: menus id; Type: DEFAULT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menus ALTER COLUMN id SET DEFAULT nextval('public.menus_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: campaigns; Type: TABLE DATA; Schema: public; Owner: dapurhafi
--

INSERT INTO public.campaigns VALUES (1, '2020-04-19 11:22:10.893208+07', '2020-04-19 11:22:10.893208+07', NULL, 'Regular', 'Regular campaign');


--
-- Data for Name: menupics; Type: TABLE DATA; Schema: public; Owner: dapurhafi
--

INSERT INTO public.menupics VALUES (1, '2020-04-19 11:22:10.925324+07', '2020-04-19 11:22:10.925324+07', NULL, 1, 'c37ab8d0072937659975b874dafe04cb.jpg');
INSERT INTO public.menupics VALUES (2, '2020-04-19 11:22:10.958262+07', '2020-04-19 11:22:10.958262+07', NULL, 2, '1b70abfe370b98393d3e08fcf97891cc.jpg');
INSERT INTO public.menupics VALUES (3, '2020-04-19 11:22:10.991317+07', '2020-04-19 11:22:10.991317+07', NULL, 3, '6533e11e314b63821e1f569899d8f331.jpg');


--
-- Data for Name: menuprices; Type: TABLE DATA; Schema: public; Owner: dapurhafi
--

INSERT INTO public.menuprices VALUES (1, '2020-04-19 11:22:10.936297+07', '2020-04-19 11:22:10.936297+07', NULL, 1, 1, 10000);
INSERT INTO public.menuprices VALUES (2, '2020-04-19 11:22:10.969279+07', '2020-04-19 11:22:10.969279+07', NULL, 1, 2, 25000);
INSERT INTO public.menuprices VALUES (3, '2020-04-19 11:22:11.002352+07', '2020-04-19 11:22:11.002352+07', NULL, 1, 3, 20000);


--
-- Data for Name: menus; Type: TABLE DATA; Schema: public; Owner: dapurhafi
--

INSERT INTO public.menus VALUES (1, '2020-04-19 11:22:10.905084+07', '2020-04-19 11:22:10.905084+07', NULL, 'Nasi Goreng', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. 

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.', 'nasi, nasi goreng', 0);
INSERT INTO public.menus VALUES (2, '2020-04-19 11:22:10.947236+07', '2020-04-19 11:22:10.947236+07', NULL, 'Rendang Daging Sapi', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. 

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.', 'daging, rendang, sapi', 0);
INSERT INTO public.menus VALUES (3, '2020-04-19 11:22:10.98043+07', '2020-04-19 11:22:10.98043+07', NULL, 'Ayam Balado', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit. 

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vestibulum convallis ligula ac cursus. Curabitur leo augue, sagittis vitae ante non, ullamcorper iaculis dui. Proin ac hendrerit velit.', 'daging, balado, ayam', 0);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: dapurhafi
--

INSERT INTO public.users VALUES (1, '2020-04-19 11:22:10.870636+07', '2020-04-19 11:22:10.870636+07', NULL, 'Albert Panesse', 'albert.panesse@gmail.com', 'albert.panesse@gmail.com', '4297f44b13955235245b2497399d7a93', '', '2020-04-19 11:22:10.869862+07', '', '2020-04-19 11:22:10.869862+07', '', '2020-04-19 11:22:10.869862+07', '', '', '');


--
-- Name: campaigns_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dapurhafi
--

SELECT pg_catalog.setval('public.campaigns_id_seq', 1, true);


--
-- Name: menupics_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dapurhafi
--

SELECT pg_catalog.setval('public.menupics_id_seq', 3, true);


--
-- Name: menuprices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dapurhafi
--

SELECT pg_catalog.setval('public.menuprices_id_seq', 3, true);


--
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dapurhafi
--

SELECT pg_catalog.setval('public.menus_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dapurhafi
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: campaigns campaigns_pkey; Type: CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);


--
-- Name: menupics menupics_pkey; Type: CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menupics
    ADD CONSTRAINT menupics_pkey PRIMARY KEY (id);


--
-- Name: menuprices menuprices_pkey; Type: CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menuprices
    ADD CONSTRAINT menuprices_pkey PRIMARY KEY (id);


--
-- Name: menus menus_pkey; Type: CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_campaigns_deleted_at; Type: INDEX; Schema: public; Owner: dapurhafi
--

CREATE INDEX idx_campaigns_deleted_at ON public.campaigns USING btree (deleted_at);


--
-- Name: idx_menupics_deleted_at; Type: INDEX; Schema: public; Owner: dapurhafi
--

CREATE INDEX idx_menupics_deleted_at ON public.menupics USING btree (deleted_at);


--
-- Name: idx_menuprices_deleted_at; Type: INDEX; Schema: public; Owner: dapurhafi
--

CREATE INDEX idx_menuprices_deleted_at ON public.menuprices USING btree (deleted_at);


--
-- Name: idx_menus_deleted_at; Type: INDEX; Schema: public; Owner: dapurhafi
--

CREATE INDEX idx_menus_deleted_at ON public.menus USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: dapurhafi
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: menupics menupics_menu_id_menus_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menupics
    ADD CONSTRAINT menupics_menu_id_menus_id_foreign FOREIGN KEY (menu_id) REFERENCES public.menus(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: menuprices menuprices_campaign_id_campaigns_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menuprices
    ADD CONSTRAINT menuprices_campaign_id_campaigns_id_foreign FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: menuprices menuprices_menu_id_menus_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: dapurhafi
--

ALTER TABLE ONLY public.menuprices
    ADD CONSTRAINT menuprices_menu_id_menus_id_foreign FOREIGN KEY (menu_id) REFERENCES public.menus(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

