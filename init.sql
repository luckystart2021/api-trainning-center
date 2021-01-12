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
	created_at timestamp(0) NOT NULL DEFAULT now(),

	CONSTRAINT user_pkey PRIMARY KEY (username)
);

CREATE UNIQUE INDEX user_unique_idx ON public."user" USING btree (username);
CREATE UNIQUE INDEX email_unique_idx ON public."user" USING btree (email);