-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial PRIMARY KEY,
	username text NOT NULL,
	"password" text NOT NULL,
	email text NULL,
	"role" text NOT NULL,
	sex text NOT NULL,
	dateofbirth text NOT NULL,
	phone text NOT NULL,
	fullname text NOT NULL,
	address text NOT NULL,
	is_delete bool NOT NULL DEFAULT false,
	available bool NOT NULL DEFAULT true,
	created_at timestamptz NOT NULL DEFAULT now(),
);
CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username, email, phone);


INSERT INTO public."users"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address)
VALUES('phong', '$2a$10$JBHml.bnZYSSIVN7ZjaLpOjOGBzv7YXauYBQ6CVaJ/prdsU/0soNO', 'thanhphong@gmail.com', 'ADMIN', 'Nam', '29/04/1997', '0832210125', 'Nguyễn Thanh Phong', '2021-01-12 15:48:15.000','Long An');
INSERT INTO public."users"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address)
VALUES('teacher', '$2a$10$buAgmI6iKeV6QP823HHqE.91GqVAUQoXb01IvJIYEw.sr/NfXWm/S', 'thanhphong1@gmail.com', 'TEACHER', 'Nam', '29/04/1999', '0832210124', 'Nguyễn Thanh Phong', '2021-01-13 13:36:21.000','Long An');


-- public.course definition

-- Drop table

-- DROP TABLE public.course;

CREATE TABLE public.course (
	id serial NOT NULL,
	code text NOT NULL,
	"name" text NOT NULL,
	start_date timestamptz(0) NOT NULL,
	end_date timestamptz(0) NOT NULL,
	graduation_date timestamptz(0) NULL,
	test_date timestamptz(0) NOT NULL,
	training_system text NOT NULL,
	status bool NOT NULL DEFAULT false,
	created_by text NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_by text NOT NULL,
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT course_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX course_code_idx ON public.course USING btree (code);