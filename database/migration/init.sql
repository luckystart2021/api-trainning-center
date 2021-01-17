-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	username text NOT NULL,
	"password" text NOT NULL,
	email text NULL,
	"role" text NOT NULL,
	sex text NOT NULL,
	dateofbirth text NOT NULL,
	phone text NOT NULL,
	fullname text NOT NULL,
	is_delete bool NOT NULL DEFAULT false,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

	CONSTRAINT user_pkey PRIMARY KEY (username)
);

CREATE UNIQUE INDEX user_unique_idx ON public."user" USING btree (username);
CREATE UNIQUE INDEX email_unique_idx ON public."user" USING btree (email);
CREATE UNIQUE INDEX phone_unique_idx ON public."user" USING btree (phone);

INSERT INTO public."user"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at)
VALUES('phong', '$2a$10$JBHml.bnZYSSIVN7ZjaLpOjOGBzv7YXauYBQ6CVaJ/prdsU/0soNO', 'thanhphong@gmail.com', 'ADMIN', 'Nam', '29/04/1997', '0832210125', 'Nguyễn Thanh Phong', '2021-01-12 15:48:15.000');
INSERT INTO public."user"
(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at)
VALUES('teacher', '$2a$10$buAgmI6iKeV6QP823HHqE.91GqVAUQoXb01IvJIYEw.sr/NfXWm/S', 'thanhphong1@gmail.com', 'TEACHER', 'Nam', '29/04/1999', '0832210124', 'Nguyễn Thanh Phong', '2021-01-13 13:36:21.000');