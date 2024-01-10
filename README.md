# This is a Work in progress basic Apollo GraphQL backend for handling Organic Plant tracing.

Use `endpoint.MakeGraphQLHandler(pool *pgxpool.Pool)` to get a server running.
You can have a look at the `server.go` file to see how to run it localy.

## Expected SQL Schema inside the database

Most of the names in here are transalted into better names entities in the graphql interface.

```SQL

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-1.pgdg100+1)
-- Dumped by pg_dump version 14.9 (Homebrew)


-- CREATE DATABASE ab_tracing WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


--
-- TOC entry 628 (class 1247 OID 17528)
-- Name: aquire_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.aquire_type AS ENUM (
    'grown',
    'gathered',
    'bought',
    'home_made'
);


--
-- TOC entry 638 (class 1247 OID 17561)
-- Name: reproductive_material_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.reproductive_material_type AS ENUM (
    'seed',
    'cutting',
    'graft'
);


--
-- TOC entry 631 (class 1247 OID 17536)
-- Name: visibility_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.visibility_type AS ENUM (
    'public_facing',
    'internal'
);

--
-- TOC entry 204 (class 1259 OID 17617)
-- Name: gather_places; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.gather_places (
    id character varying DEFAULT concat('place:', gen_random_uuid()) NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL,
    name character varying,
    is_organic_compatible boolean DEFAULT true NOT NULL,
    country character varying DEFAULT ''::character varying NOT NULL,
    address character varying DEFAULT ''::character varying NOT NULL
);


--
-- TOC entry 203 (class 1259 OID 17603)
-- Name: growing_material; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.growing_material (
    id character varying DEFAULT concat('mat:', gen_random_uuid()) NOT NULL,
    aquisition_type public.aquire_type DEFAULT 'home_made'::public.aquire_type NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    production_steps character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    creation_date date,
    quantity real DEFAULT 0 NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL,
    aquisition_places character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    aquisition_bought character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    name character varying,
    is_organic_compliant boolean DEFAULT true NOT NULL,
    unit character varying DEFAULT 'Kg'::character varying NOT NULL
);


--
-- TOC entry 200 (class 1259 OID 17544)
-- Name: plant; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.plant (
    id character varying DEFAULT concat('plant:', gen_random_uuid()) NOT NULL,
    aquisition_type public.aquire_type DEFAULT 'grown'::public.aquire_type NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    source character varying,
    grafting_sources character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    maturation_sources character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    treatment_sources character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    planting_date date,
    name_latin character varying DEFAULT 'name'::character varying NOT NULL,
    quantity integer DEFAULT 0 NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL,
    aquisition_places character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    aquisition_bought character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    is_stock_plant boolean DEFAULT false NOT NULL,
    name character varying,
    is_organic boolean DEFAULT true NOT NULL,
    unit character varying DEFAULT 'plants'::character varying NOT NULL
);


--
-- TOC entry 201 (class 1259 OID 17570)
-- Name: plant_reproduction_material; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.plant_reproduction_material (
    id character varying DEFAULT concat('prm:', gen_random_uuid()) NOT NULL,
    aquisition_type public.aquire_type DEFAULT 'grown'::public.aquire_type NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    germination_source character varying,
    treatment_steps character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    harvest_source character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    production_date date,
    name_latin character varying DEFAULT 'name'::character varying NOT NULL,
    quantity real DEFAULT 0 NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL,
    aquisition_places character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    aquisition_bought character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    name character varying,
    is_organic boolean DEFAULT true NOT NULL,
    unit character varying DEFAULT 'plants'::character varying NOT NULL,
    type public.reproductive_material_type DEFAULT 'seed'::public.reproductive_material_type NOT NULL
);


--
-- TOC entry 202 (class 1259 OID 17589)
-- Name: plant_treatments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.plant_treatments (
    id character varying DEFAULT concat('treat:', gen_random_uuid()) NOT NULL,
    aquisition_type public.aquire_type DEFAULT 'home_made'::public.aquire_type NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    production_ingredients character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    creation_date date,
    quantity integer DEFAULT 0 NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL,
    aquisition_places character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    aquisition_bought character varying[] DEFAULT ARRAY[]::character varying[] NOT NULL,
    name character varying DEFAULT ''::character varying NOT NULL,
    is_organic_compatible boolean DEFAULT true NOT NULL,
    unit character varying DEFAULT 'Kg'::character varying NOT NULL
);


--
-- TOC entry 205 (class 1259 OID 17631)
-- Name: supply_info; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.supply_info (
    id character varying DEFAULT concat('extrnl:', gen_random_uuid()) NOT NULL,
    visibility public.visibility_type DEFAULT 'internal'::public.visibility_type NOT NULL,
    name character varying,
    country character varying,
    supplier character varying DEFAULT ''::character varying NOT NULL,
    bill_link character varying DEFAULT ''::character varying NOT NULL,
    notes character varying DEFAULT ''::character varying NOT NULL
);


```
