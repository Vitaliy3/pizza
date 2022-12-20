--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0
-- Dumped by pg_dump version 15.0

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
-- Name: banned; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banned (
    id integer NOT NULL,
    telnumber character varying(20)
);


ALTER TABLE public.banned OWNER TO postgres;

--
-- Name: banned_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banned_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.banned_id_seq OWNER TO postgres;

--
-- Name: banned_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banned_id_seq OWNED BY public.banned.id;


--
-- Name: buy; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.buy (
    id integer NOT NULL,
    amount integer,
    inbound boolean DEFAULT true NOT NULL,
    fk_product integer,
    c_location text,
    telnumber text,
    c_text text,
    fk_user integer,
    called boolean DEFAULT false NOT NULL
);


ALTER TABLE public.buy OWNER TO postgres;

--
-- Name: buy_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.buy_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buy_id_seq OWNER TO postgres;

--
-- Name: buy_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.buy_id_seq OWNED BY public.buy.id;


--
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category (
    id integer NOT NULL,
    c_name text
);


ALTER TABLE public.category OWNER TO postgres;

--
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.category_id_seq OWNER TO postgres;

--
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.category_id_seq OWNED BY public.category.id;


--
-- Name: product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    price double precision NOT NULL,
    image text NOT NULL,
    fk_user integer,
    fk_category integer
);


ALTER TABLE public.product OWNER TO postgres;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO postgres;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    rname character varying(50)
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    telnumber character varying(20),
    pass text,
    fk_role integer,
    blocked boolean NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: banned id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banned ALTER COLUMN id SET DEFAULT nextval('public.banned_id_seq'::regclass);


--
-- Name: buy id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buy ALTER COLUMN id SET DEFAULT nextval('public.buy_id_seq'::regclass);


--
-- Name: category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.category_id_seq'::regclass);


--
-- Name: product id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: banned; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banned (id, telnumber) FROM stdin;
5	79873841215
\.


--
-- Data for Name: buy; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.buy (id, amount, inbound, fk_product, c_location, telnumber, c_text, fk_user, called) FROM stdin;
26	1	t	15	46.016895:51.538736	79873848536	79873848536, в помещении 1	\N	f
28	1	f	11	46.016895:51.538736	79873848536	79873848536, в помещении 1	\N	f
29	1	t	14	46.019502:51.537480	79873841215	79873841215 , в помещении 2	\N	f
30	1	t	16	46.020531:51.536342	79873841215	79873841215 , в помещении не указано	\N	f
31	1	t	16	46.022102:51.541489	79873841215	79873841215 , в помещении не указано	\N	f
32	1	t	15	46.021669:51.537697	79873841215	79873841215 , в помещении не указано	\N	f
33	1	t	15	46.019664:51.543711	79873841215	79873841215 , в помещении не указано	\N	f
34	1	t	16	46.021290:51.536776	79873841215	79873841215 , в помещении не указано	\N	f
35	1	t	15	46.022265:51.535204	79873841215	79873841215 , в помещении не указано	\N	f
36	1	t	16	46.022563:51.538130	79873841215	79873841215 , в помещении не указано	\N	f
27	1	f	15	46.016895:51.538736	79873848536	79873848536, в помещении 1	\N	f
25	1	f	15	46.016895:51.538736	79873848536	79873848536, в помещении 1	\N	t
\.


--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category (id, c_name) FROM stdin;
1	Пицца
2	Соки/Напитки
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (id, title, description, price, image, fk_user, fk_category) FROM stdin;
11	Деревенская	Ветчина, куриное филе, сыр пармезан, сыр моцарелла, томаты, орегано, соус горчичный.	340	cheese.png	18	1
12	Сырная гурмэ	Куриное филе, ветчина, огурец маринованый, помидоры, сыр моцарелла, соус сливочный, соус сырный	350	dereven.png	18	1
13	Фантола бабл гам 0,5		255	fantola.png	18	2
14	Морс клюква-виноград 0,3л		265	dobriy.jpg	18	2
15	Чай Rich чер. 0,5л		260	tea.jpg	18	2
16	Сок Клюква		250	dobriy.jpg	18	2
5	4 сезона	Пепперони, курица, ветчина, шампиньоны, моцарелла, томат, томатный соус.	400	4season.png	18	1
3	Бургерная	Говядина копченая, томаты, лук красный, огурцы маринованные, сыр пармезан, сыр моцарелла, соус бургерный.	370	burger.jpg	18	1
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, rname) FROM stdin;
3	ADMIN
1	CUSTOMER
2	MANAGER
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, telnumber, pass, fk_role, blocked) FROM stdin;
34	79873848535	123	3	f
18	79873848536	123	1	f
35	79873848537	123	1	f
36	79873848538	123	1	f
37	79873848539	123	1	f
38	79873848540	123	1	f
33	79873841215	123	3	f
\.


--
-- Name: banned_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banned_id_seq', 5, true);


--
-- Name: buy_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.buy_id_seq', 36, true);


--
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.category_id_seq', 12, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_id_seq', 21, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 38, true);


--
-- Name: banned banned_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banned
    ADD CONSTRAINT banned_pkey PRIMARY KEY (id);


--
-- Name: buy buy_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buy
    ADD CONSTRAINT buy_pkey PRIMARY KEY (id);


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: buy buy_fk_product_product_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buy
    ADD CONSTRAINT buy_fk_product_product_id FOREIGN KEY (fk_product) REFERENCES public.product(id);


--
-- Name: buy buy_fk_user_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buy
    ADD CONSTRAINT buy_fk_user_user_id FOREIGN KEY (fk_user) REFERENCES public.users(id);


--
-- Name: product product_fk_caterogy_category_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_fk_caterogy_category_id FOREIGN KEY (fk_category) REFERENCES public.category(id);


--
-- Name: users users_fk_role_roles_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk_role_roles_id FOREIGN KEY (fk_role) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

