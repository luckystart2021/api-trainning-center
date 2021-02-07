-- users definition

-- Drop table

-- DROP TABLE users;

CREATE TABLE users (
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
	created_at timestamptz NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX users_username_idx ON users USING btree (username, email, phone);


INSERT INTO "users"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address)
VALUES('phong', '$2a$10$JBHml.bnZYSSIVN7ZjaLpOjOGBzv7YXauYBQ6CVaJ/prdsU/0soNO', 'thanhphong@gmail.com', 'ADMIN', 'Nam', '29/04/1997', '0832210125', 'Nguyễn Thanh Phong', '2021-01-12 15:48:15.000','Long An');
INSERT INTO "users"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address)
VALUES('teacher', '$2a$10$buAgmI6iKeV6QP823HHqE.91GqVAUQoXb01IvJIYEw.sr/NfXWm/S', 'thanhphong1@gmail.com', 'TEACHER', 'Nam', '29/04/1999', '0832210124', 'Nguyễn Thanh Phong', '2021-01-13 13:36:21.000','Long An');

-- course definition

-- Drop table

-- DROP TABLE course;

CREATE TABLE course (
	id serial NOT NULL,
	code text NOT NULL,
	name text NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	graduation_date date NULL,
	test_date date NOT NULL,
	training_system text NOT NULL,
	status bool NOT NULL DEFAULT false,
	created_by text NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_by text NOT NULL,
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT course_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX course_code_idx ON course USING btree (code);

-- contact definition

-- Drop table

-- DROP TABLE contact;

CREATE TABLE contact (
	id serial NOT NULL,
	fullname text NOT NULL,
	phone text NOT NULL,
	email text NULL,
	message text NOT NULL,
	subject text NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT contact_pk PRIMARY KEY (id)
);

-- public.information definition

-- Drop table

-- DROP TABLE public.information;

CREATE TABLE information (
	id serial NOT NULL,
	address text NOT NULL,
	email text NOT NULL,
	phone text NOT NULL,
	maps text NOT NULL,
	title text NOT NULL,
	description text NOT NULL,
	img text NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT information_pk PRIMARY KEY (id)
);

-- public.testsuite definition

-- Drop table

-- DROP TABLE public.testsuite;

CREATE TABLE testsuite (
	id serial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT testsuite_pk PRIMARY KEY (id)
);

-- public.question definition

-- Drop table

-- DROP TABLE public.question;

CREATE TABLE question (
	id serial NOT NULL,
	"name" text NOT NULL,
	"result" text NOT NULL,
	paralysis bool NOT NULL DEFAULT false,
	id_code_test int4 NOT NULL,
	answera text NULL,
	answerb text NULL,
	answerc text NULL,
	answerd text NULL,
	img text NULL,
	CONSTRAINT question_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX question_name_idx ON question USING btree (name);


-- public.question foreign keys

ALTER TABLE question ADD CONSTRAINT question_fk FOREIGN KEY (id_code_test) REFERENCES testsuite(id);


-- public.notification definition

-- Drop table

-- DROP TABLE public.notification;

CREATE TABLE notification (
	id serial NOT NULL,
	title text NOT NULL,
	description text NOT NULL,
	subtitle text NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	img text NOT NULL,
	CONSTRAINT notification_pk PRIMARY KEY (id)
);
INSERT INTO "notification"
(id, title, description, subtitle, created_at, img)
VALUES(1, 'PHONG', 'PHONG', 'OK', '2021-02-06 18:38:23.000', 'banner.jpg');
