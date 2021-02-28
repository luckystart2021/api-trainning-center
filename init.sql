-- -- users definition

-- -- Drop table

-- -- DROP TABLE users;

-- CREATE TABLE users (
-- 	id serial PRIMARY KEY,
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
-- 	created_at timestamptz NOT NULL DEFAULT now()
-- );
-- CREATE UNIQUE INDEX users_username_idx ON users USING btree (username, email, phone);


-- INSERT INTO "users"(username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address) VALUES('phong', '$2a$10$JBHml.bnZYSSIVN7ZjaLpOjOGBzv7YXauYBQ6CVaJ/prdsU/0soNO', 'thanhphong@gmail.com', 'ADMIN', 'Nam', '29/04/1997', '0832210125', 'Nguyễn Thanh Phong', '2021-01-12 15:48:15.000','Long An');
-- INSERT INTO "users"
-- (username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, address)
-- VALUES('teacher', '$2a$10$buAgmI6iKeV6QP823HHqE.91GqVAUQoXb01IvJIYEw.sr/NfXWm/S', 'thanhphong1@gmail.com', 'TEACHER', 'Nam', '29/04/1999', '0832210124', 'Nguyễn Thanh Phong', '2021-01-13 13:36:21.000','Long An');

-- -- course definition

-- -- Drop table

-- -- DROP TABLE course;

-- CREATE TABLE course (
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
-- CREATE UNIQUE INDEX course_code_idx ON course USING btree (code);

-- -- contact definition

-- -- Drop table

-- -- DROP TABLE contact;

-- CREATE TABLE contact (
-- 	id serial NOT NULL,
-- 	fullname text NOT NULL,
-- 	phone text NOT NULL,
-- 	email text NULL,
-- 	message text NOT NULL,
-- 	subject text NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	CONSTRAINT contact_pk PRIMARY KEY (id)
-- );

-- -- information definition

-- -- Drop table

-- -- DROP TABLE information;

-- CREATE TABLE information (
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

-- INSERT INTO "information"
-- (id, address, email, phone, maps, title, description, img, created_at)
-- VALUES(1, '38 Tây Hòa', '0832210125', 'thanhphong9718@gmail.com', '<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3918.807549595758!2d106.76057895063911!3d10.826034992250055!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x317527bd92bda2c1%3A0x16607d0fd6c0392f!2zMzggVMOieSBIw7JhLCBQaMaw4bubYyBMb25nIEEsIFF14bqtbiA5LCBUaMOgbmggcGjhu5EgSOG7kyBDaMOtIE1pbmgsIFZp4buHdCBOYW0!5e0!3m2!1svi!2s!4v1612799368639!5m2!1svi!2s" width="600" height="450" frameborder="0" style="border:0;" allowfullscreen="" aria-hidden="false" tabindex="0"></iframe>', 'Phong', 'Phong', 'tt.jpg', '2021-02-08 22:49:34.000');


-- -- testsuite definition

-- -- Drop table

-- -- DROP TABLE testsuite;

-- CREATE TABLE testsuite (
-- 	id serial NOT NULL,
-- 	"name" text NOT NULL,
-- 	CONSTRAINT testsuite_pk PRIMARY KEY (id)
-- );

-- -- question definition

-- -- Drop table

-- -- DROP TABLE question;

-- CREATE TABLE question (
-- 	id serial NOT NULL,
-- 	"name" text NOT NULL,
-- 	"result" text NOT NULL,
-- 	paralysis bool NOT NULL DEFAULT false,
-- 	id_code_test int4 NOT NULL,
-- 	answera text NULL,
-- 	answerb text NULL,
-- 	answerc text NULL,
-- 	answerd text NULL,
-- 	img text NULL,
-- 	CONSTRAINT question_pk PRIMARY KEY (id)
-- );
-- CREATE UNIQUE INDEX question_name_idx ON question USING btree (name);
-- -- question foreign keys

-- ALTER TABLE question ADD CONSTRAINT question_fk FOREIGN KEY (id_code_test) REFERENCES testsuite(id);

-- -- notification definition

-- -- Drop table

-- -- DROP TABLE notification;

-- CREATE TABLE notification (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	description text NOT NULL,
-- 	subtitle text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	img text NOT NULL,
-- 	CONSTRAINT notification_pk PRIMARY KEY (id)
-- );
-- INSERT INTO "notification"
-- (id, title, description, subtitle, created_at, img)
-- VALUES(1, 'PHONG', 'PHONG', 'OK', '2021-02-06 18:38:23.000', 'banner.jpg');




-- -- public.category definition

-- -- Drop table

-- -- DROP TABLE public.category;

-- CREATE TABLE public.category (id serial NOT NULL,
-- 	title text NOT NULL,
-- 	created_by text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	meta text NOT NULL,
-- 	CONSTRAINT category_pk PRIMARY KEY (id)
-- );

-- INSERT INTO public.category
-- (id, title, created_by, created_at, updated_by, updated_at, meta)
-- VALUES(1, 'Thông tin', 'admin', '2021-02-14 12:55:48.000', 'admin', '2021-02-14 12:55:48.000', 'thong-tin');
-- INSERT INTO public.category
-- (id, title, created_by, created_at, updated_by, updated_at, meta)
-- VALUES(2, 'Tin tức', 'admin', '2021-02-14 12:55:48.000', 'admin', '2021-02-14 12:55:48.000', 'tin-tuc');

-- -- public.child_category definition

-- -- Drop table

-- -- DROP TABLE public.child_category;

-- CREATE TABLE public.child_category (
-- 	id serial NOT NULL,
-- 	title text NOT NULL,
-- 	id_category int4 NOT NULL,
-- 	meta text NOT NULL,
-- 	created_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	created_by text NOT NULL,
-- 	updated_at timestamptz(0) NOT NULL DEFAULT now(),
-- 	updated_by text NOT NULL,
-- 	CONSTRAINT child_category_pk PRIMARY KEY (id)
-- );


-- -- public.child_category foreign keys

-- ALTER TABLE public.child_category ADD CONSTRAINT child_category_fk FOREIGN KEY (id_category) REFERENCES category(id);

-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(1, 'Thông báo', 1, 'thong-bao', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(2, 'Thông báo - Chiêu Sinh', 1, 'thong-bao-chieu-sinh', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(3, 'Thông tin liên quan đến GPLX', 1, 'thong-tin-lien-quan-den-GPLX', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(4, 'Hoạt động từ TT Hoàng Gia', 2, 'hoat-dong', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(5, 'Những lưu ý khi thi GPLX', 2, 'nhung-luu-y-khi-thi-GPLX', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(6, 'An toàn giao thông', 2, 'an-toan-giao-thong', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');
-- INSERT INTO public.child_category
-- (id, title, id_category, meta, created_at, created_by, updated_at, updated_by)
-- VALUES(7, 'Kinh Nghiệm', 2, 'kinh-nghiem', '2021-02-14 12:14:02.000', 'phong', '2021-02-14 12:14:02.000', 'phong');

-- -- public.articles definition

-- -- Drop table

-- -- DROP TABLE public.articles;

-- CREATE TABLE public.articles (
-- 	id serial NOT NULL,
-- 	id_user int4 NOT NULL,
-- 	id_child_category int4 NOT NULL,
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
-- 	CONSTRAINT articles_pk PRIMARY KEY (id)
-- );
-- CREATE UNIQUE INDEX articles_title_idx ON public.articles USING btree (title);


-- -- public.articles foreign keys

-- ALTER TABLE public.articles ADD CONSTRAINT articles_fk FOREIGN KEY (id_child_category) REFERENCES child_category(id);
-- ALTER TABLE public.articles ADD CONSTRAINT articles_user_fk FOREIGN KEY (id_user) REFERENCES users(id);