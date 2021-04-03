-- -- public.album definition

-- -- Drop table

-- -- DROP TABLE public.album;

-- CREATE TABLE public.album (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	meta text NOT NULL,
-- 	CONSTRAINT album_pk PRIMARY KEY (id)
-- );


-- -- public.article_tag definition

-- -- Drop table

-- -- DROP TABLE public.article_tag;

-- CREATE TABLE public.article_tag (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	CONSTRAINT article_tag_pk PRIMARY KEY (id)
-- );


-- -- public.category definition

-- -- Drop table

-- -- DROP TABLE public.category;

-- CREATE TABLE public.category (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	meta text NOT NULL,
-- 	CONSTRAINT category_pk PRIMARY KEY (id)
-- );


-- -- public.contact definition

-- -- Drop table

-- -- DROP TABLE public.contact;

-- CREATE TABLE public.contact (
-- 	id serial NOT NULL,
-- 	fullname text NOT NULL,
-- 	phone text NOT NULL,
-- 	email text NULL,
-- 	message text NOT NULL,
-- 	subject text NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT contact_pk PRIMARY KEY (id)
-- );


-- -- public.course definition

-- -- Drop table

-- -- DROP TABLE public.course;

-- CREATE TABLE public.course (
-- 	id serial NOT NULL,
-- 	code text NOT NULL,
-- 	name text NOT NULL,
-- 	start_date date NOT NULL,
-- 	end_date date NOT NULL,
-- 	graduation_date date NULL,
-- 	test_date date NOT NULL,
-- 	training_system text NOT NULL,
-- 	status bool NOT NULL DEFAULT false,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT course_pk PRIMARY KEY (id)
-- );
-- CREATE UNIQUE INDEX course_code_idx ON public.course USING btree (code);


-- -- public.information definition

-- -- Drop table

-- -- DROP TABLE public.information;

-- CREATE TABLE public.information (
-- 	id serial NOT NULL,
-- 	address text NOT NULL,
-- 	email text NOT NULL,
-- 	phone text NOT NULL,
-- 	maps text NOT NULL,
-- 	title text NOT NULL,
-- 	description text NOT NULL,
-- 	img text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT information_pk PRIMARY KEY (id)
-- );


-- -- public.notification definition

-- -- Drop table

-- -- DROP TABLE public.notification;

-- CREATE TABLE public.notification (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	description text NOT NULL,
-- 	subtitle text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	img text NOT NULL,
-- 	CONSTRAINT notification_pk PRIMARY KEY (id)
-- );


-- -- public.question definition

-- -- Drop table

-- -- DROP TABLE public.question;

-- CREATE TABLE public.question (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	anwser_correct text NOT NULL,
-- 	paralysis bool NOT NULL DEFAULT false,
-- 	answera text NOT NULL,
-- 	answerb text NOT NULL,
-- 	answerc text NULL,
-- 	answerd text NULL,
-- 	img text NULL,
-- 	question_type text NOT NULL,
-- 	CONSTRAINT question_pk PRIMARY KEY (id)
-- );


-- -- public.rank_vehicle definition

-- -- Drop table

-- -- DROP TABLE public.rank_vehicle;

-- CREATE TABLE public.rank_vehicle (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	"time" int4 NOT NULL,
-- 	number_question int4 NOT NULL,
-- 	point_pass int4 NOT NULL,
-- 	CONSTRAINT rank_vehicle_pk PRIMARY KEY (id)
-- );


-- -- public."role" definition

-- -- Drop table

-- -- DROP TABLE public."role";

-- CREATE TABLE public."role" (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	CONSTRAINT role_pk PRIMARY KEY (id)
-- );


-- -- public.seo definition

-- -- Drop table

-- -- DROP TABLE public.seo;

-- CREATE TABLE public.seo (
-- 	id serial NOT NULL,
-- 	description text NOT NULL,
-- 	keywords text NOT NULL,
-- 	fb_app_id text NOT NULL,
-- 	og_title text NOT NULL,
-- 	og_url text NOT NULL,
-- 	og_image text NOT NULL,
-- 	og_description text NOT NULL,
-- 	og_site_name text NOT NULL,
-- 	og_see_also text NOT NULL,
-- 	og_locale text NOT NULL,
-- 	article_author text NOT NULL,
-- 	twitter_card text NOT NULL,
-- 	twitter_url text NOT NULL,
-- 	twitter_title text NOT NULL,
-- 	twitter_description text NOT NULL,
-- 	twitter_image text NOT NULL,
-- 	author text NOT NULL,
-- 	generator text NOT NULL,
-- 	copyright text NOT NULL,
-- 	CONSTRAINT seo_pk PRIMARY KEY (id)
-- );


-- -- public.slide definition

-- -- Drop table

-- -- DROP TABLE public.slide;

-- CREATE TABLE public.slide (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	img text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	created_by text NOT NULL,
-- 	hide bool NOT NULL DEFAULT false,
-- 	CONSTRAINT slide_pk PRIMARY KEY (id)
-- );


-- -- public.subject definition

-- -- Drop table

-- -- DROP TABLE public.subject;

-- CREATE TABLE public.subject (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	"time" int4 NOT NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT subject_pk PRIMARY KEY (id)
-- );


-- -- public.teacher definition

-- -- Drop table

-- -- DROP TABLE public.teacher;

-- CREATE TABLE public.teacher (
-- 	id serial NOT NULL,
-- 	fullname text NOT NULL,
-- 	sex text NOT NULL,
-- 	dateofbirth text NOT NULL,
-- 	phone text NOT NULL,
-- 	address text NOT NULL,
-- 	cmnd text NOT NULL,
-- 	cnsk bool NOT NULL DEFAULT false,
-- 	gplx varchar NULL,
-- 	experience_driver int4 NOT NULL,
-- 	km_safe int4 NOT NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT teacher_pk PRIMARY KEY (id)
-- );


-- -- public.users definition

-- -- Drop table

-- -- DROP TABLE public.users;

-- CREATE TABLE public.users (
-- 	id serial NOT NULL,
-- 	username text NOT NULL,
-- 	"password" text NOT NULL,
-- 	email text NULL,
-- 	"role" text NOT NULL,
-- 	sex text NOT NULL,
-- 	dateofbirth text NOT NULL,
-- 	phone text NOT NULL,
-- 	fullname text NOT NULL,
-- 	address text NOT NULL,
-- 	is_delete bool NOT NULL DEFAULT false,
-- 	available bool NOT NULL DEFAULT true,
-- 	created_at timestamptz NOT NULL DEFAULT now(),
-- 	CONSTRAINT users_pkey PRIMARY KEY (id)
-- );
-- CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username, email, phone);


-- -- public.vehicle definition

-- -- Drop table

-- -- DROP TABLE public.vehicle;

-- CREATE TABLE public.vehicle (
-- 	id serial NOT NULL,
-- 	biensoxe text NOT NULL,
-- 	loaixe text NOT NULL,
-- 	status bool NOT NULL DEFAULT false,
-- 	is_deleted bool NOT NULL DEFAULT false,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	CONSTRAINT vehicle_pk PRIMARY KEY (id)
-- );
-- CREATE UNIQUE INDEX vehicle_biensoxe_idx ON public.vehicle USING btree (biensoxe);


-- -- public.child_category definition

-- -- Drop table

-- -- DROP TABLE public.child_category;

-- CREATE TABLE public.child_category (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	category_id int4 NOT NULL,
-- 	meta text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	created_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	is_deleted bool NOT NULL DEFAULT false,
-- 	CONSTRAINT child_category_pk PRIMARY KEY (id),
-- 	CONSTRAINT child_category_fk FOREIGN KEY (category_id) REFERENCES category(id)
-- );
-- CREATE UNIQUE INDEX child_category_title_idx ON public.child_category USING btree (title);


-- -- public."class" definition

-- -- Drop table

-- -- DROP TABLE public."class";

-- CREATE TABLE public."class" (
-- 	id serial NOT NULL,
-- 	code text NOT NULL,
-- 	name text NOT NULL,
-- 	course_id int4 NOT NULL,
-- 	quantity int4 NOT NULL,
-- 	id_teacher int4 NOT NULL,
-- 	created_by text NOT NULL,
-- 	updated_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	is_deleted bool NOT NULL DEFAULT false,
-- 	CONSTRAINT class_pk PRIMARY KEY (id),
-- 	CONSTRAINT class_fk FOREIGN KEY (course_id) REFERENCES course(id)
-- );
-- CREATE UNIQUE INDEX class_code_idx ON public.class USING btree (code);


-- -- public.photos definition

-- -- Drop table

-- -- DROP TABLE public.photos;

-- CREATE TABLE public.photos (
-- 	id serial NOT NULL,
-- 	img text NOT NULL,
-- 	title text NULL,
-- 	meta text NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	album_id int4 NOT NULL,
-- 	CONSTRAINT photos_pk PRIMARY KEY (id),
-- 	CONSTRAINT photos_fk FOREIGN KEY (album_id) REFERENCES album(id)
-- );


-- -- public.student definition

-- -- Drop table

-- -- DROP TABLE public.student;

-- CREATE TABLE public.student (
-- 	id serial NOT NULL,
-- 	code text NOT NULL,
-- 	sex text NOT NULL,
-- 	dateofbirth text NOT NULL,
-- 	phone text NOT NULL,
-- 	address text NOT NULL,
-- 	fullname text NOT NULL,
-- 	class_id int4 NOT NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	cmnd text NOT NULL,
-- 	cnsk bool NOT NULL DEFAULT false,
-- 	gplx varchar NULL,
-- 	experience_driver int4 NOT NULL,
-- 	km_safe int4 NOT NULL,
-- 	CONSTRAINT student_pk PRIMARY KEY (id),
-- 	CONSTRAINT student_fk FOREIGN KEY (class_id) REFERENCES class(id)
-- );
-- CREATE UNIQUE INDEX student_code_idx ON public.student USING btree (code, phone);


-- -- public.testsuite definition

-- -- Drop table

-- -- DROP TABLE public.testsuite;

-- CREATE TABLE public.testsuite (
-- 	id serial NOT NULL,
-- 	name text NOT NULL,
-- 	rank_id int4 NOT NULL,
-- 	CONSTRAINT testsuite_pk PRIMARY KEY (id),
-- 	CONSTRAINT testsuite_fk FOREIGN KEY (rank_id) REFERENCES rank_vehicle(id)
-- );


-- -- public.testsuite_question definition

-- -- Drop table

-- -- DROP TABLE public.testsuite_question;

-- CREATE TABLE public.testsuite_question (
-- 	id bigserial NOT NULL,
-- 	testsuite_id int4 NOT NULL,
-- 	question_id int4 NOT NULL,
-- 	CONSTRAINT testsuite_question_pk PRIMARY KEY (id),
-- 	CONSTRAINT testsuite_question_fk_id_question FOREIGN KEY (question_id) REFERENCES question(id),
-- 	CONSTRAINT testsuite_question_fk_id_testsuite FOREIGN KEY (testsuite_id) REFERENCES testsuite(id)
-- );
-- CREATE UNIQUE INDEX testsuite_question_id_testsuite_idx ON public.testsuite_question USING btree (testsuite_id, question_id);


-- -- public.articles definition

-- -- Drop table

-- -- DROP TABLE public.articles;

-- CREATE TABLE public.articles (
-- 	id serial NOT NULL,
-- 	user_id int4 NOT NULL,
-- 	child_category_id int4 NOT NULL,
-- 	title text NOT NULL,
-- 	description text NOT NULL,
-- 	details text NOT NULL,
-- 	image text NOT NULL,
-- 	meta text NOT NULL,
-- 	keywordseo text NOT NULL,
-- 	"view" int8 NOT NULL DEFAULT 0,
-- 	status bool NOT NULL DEFAULT false,
-- 	is_deleted bool NOT NULL DEFAULT false,
-- 	created_by text NOT NULL,
-- 	updated_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT articles_pk PRIMARY KEY (id),
-- 	CONSTRAINT articles_fk FOREIGN KEY (child_category_id) REFERENCES child_category(id),
-- 	CONSTRAINT articles_user_fk FOREIGN KEY (user_id) REFERENCES users(id)
-- );