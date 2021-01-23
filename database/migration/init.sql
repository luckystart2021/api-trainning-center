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
	created_at timestamptz NOT NULL DEFAULT now(),
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