--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.12
-- Dumped by pg_dump version 9.5.12

-- Started on 2018-04-14 20:12:16 +03

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2175 (class 1262 OID 16387)
-- Name: go_homework_1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE go_homework_1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_GB.UTF-8' LC_CTYPE = 'en_GB.UTF-8';


ALTER DATABASE go_homework_1 OWNER TO postgres;

\connect go_homework_1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12395)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2178 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 184 (class 1259 OID 16401)
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    id integer NOT NULL,
    balance integer DEFAULT 0,
    client_id integer NOT NULL
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- TOC entry 183 (class 1259 OID 16399)
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO postgres;

--
-- TOC entry 2179 (class 0 OID 0)
-- Dependencies: 183
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- TOC entry 182 (class 1259 OID 16390)
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clients (
    id integer NOT NULL,
    name character varying(255),
    email character varying(255),
    phone character varying(255)
);


ALTER TABLE public.clients OWNER TO postgres;

--
-- TOC entry 181 (class 1259 OID 16388)
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clients_id_seq OWNER TO postgres;

--
-- TOC entry 2180 (class 0 OID 0)
-- Dependencies: 181
-- Name: clients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.clients_id_seq OWNED BY public.clients.id;


--
-- TOC entry 188 (class 1259 OID 18253)
-- Name: money_changes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.money_changes (
    id integer NOT NULL,
    operation character varying NOT NULL,
    account_id integer NOT NULL,
    amount integer NOT NULL,
    transaction_id integer NOT NULL
);


ALTER TABLE public.money_changes OWNER TO postgres;

--
-- TOC entry 187 (class 1259 OID 18251)
-- Name: money_changes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.money_changes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.money_changes_id_seq OWNER TO postgres;

--
-- TOC entry 2181 (class 0 OID 0)
-- Dependencies: 187
-- Name: money_changes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.money_changes_id_seq OWNED BY public.money_changes.id;


--
-- TOC entry 186 (class 1259 OID 16410)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    message character varying(255) NOT NULL,
    error character varying(255)
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 185 (class 1259 OID 16408)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 2182 (class 0 OID 0)
-- Dependencies: 185
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 2041 (class 2604 OID 16404)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- TOC entry 2040 (class 2604 OID 16393)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients ALTER COLUMN id SET DEFAULT nextval('public.clients_id_seq'::regclass);


--
-- TOC entry 2044 (class 2604 OID 18256)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.money_changes ALTER COLUMN id SET DEFAULT nextval('public.money_changes_id_seq'::regclass);


--
-- TOC entry 2043 (class 2604 OID 16413)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 2048 (class 2606 OID 16407)
-- Name: accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- TOC entry 2046 (class 2606 OID 16398)
-- Name: clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- TOC entry 2052 (class 2606 OID 18261)
-- Name: money_changes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.money_changes
    ADD CONSTRAINT money_changes_pk PRIMARY KEY (id);


--
-- TOC entry 2050 (class 2606 OID 16418)
-- Name: transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 2053 (class 2606 OID 16424)
-- Name: accounts_clients_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_clients_fk FOREIGN KEY (client_id) REFERENCES public.clients(id);


--
-- TOC entry 2054 (class 2606 OID 18262)
-- Name: money_changes_accounts_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.money_changes
    ADD CONSTRAINT money_changes_accounts_fk FOREIGN KEY (account_id) REFERENCES public.accounts(id);


--
-- TOC entry 2055 (class 2606 OID 18267)
-- Name: money_changes_transactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.money_changes
    ADD CONSTRAINT money_changes_transactions_fk FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- TOC entry 2177 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2018-04-14 20:12:16 +03

--
-- PostgreSQL database dump complete
--

