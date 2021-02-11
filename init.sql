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

-- information definition

-- Drop table

-- DROP TABLE information;

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

INSERT INTO "information"
(id, address, email, phone, maps, title, description, img, created_at)
VALUES(1, '38 Tây Hòa', '0832210125', 'thanhphong9718@gmail.com', '<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3918.807549595758!2d106.76057895063911!3d10.826034992250055!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x317527bd92bda2c1%3A0x16607d0fd6c0392f!2zMzggVMOieSBIw7JhLCBQaMaw4bubYyBMb25nIEEsIFF14bqtbiA5LCBUaMOgbmggcGjhu5EgSOG7kyBDaMOtIE1pbmgsIFZp4buHdCBOYW0!5e0!3m2!1svi!2s!4v1612799368639!5m2!1svi!2s" width="600" height="450" frameborder="0" style="border:0;" allowfullscreen="" aria-hidden="false" tabindex="0"></iframe>', 'Phong', 'Phong', 'tt.jpg', '2021-02-08 22:49:34.000');


-- testsuite definition

-- Drop table

-- DROP TABLE testsuite;

CREATE TABLE testsuite (
	id serial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT testsuite_pk PRIMARY KEY (id)
);

-- question definition

-- Drop table

-- DROP TABLE question;

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


-- question foreign keys

ALTER TABLE question ADD CONSTRAINT question_fk FOREIGN KEY (id_code_test) REFERENCES testsuite(id);


-- notification definition

-- Drop table

-- DROP TABLE notification;

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

INSERT INTO testsuite
(id, "name")
VALUES(1, 'Đề 1');
INSERT INTO testsuite
(id, "name")
VALUES(2, 'Đề 2');
INSERT INTO testsuite
(id, "name")
VALUES(3, 'Đề 3');
INSERT INTO testsuite
(id, "name")
VALUES(4, 'Đề 4');
INSERT INTO testsuite
(id, "name")
VALUES(5, 'Đề 5');
INSERT INTO testsuite
(id, "name")
VALUES(6, 'Đề 6');
INSERT INTO testsuite
(id, "name")
VALUES(7, 'Đề 7');
INSERT INTO testsuite
(id, "name")
VALUES(8, 'Đề 8');
INSERT INTO testsuite
(id, "name")
VALUES(9, 'Đề 9');
INSERT INTO testsuite
(id, "name")
VALUES(10, 'Đề 10');
INSERT INTO testsuite
(id, "name")
VALUES(11, 'Đề 11');
INSERT INTO testsuite
(id, "name")
VALUES(12, 'Đề 12');
INSERT INTO testsuite
(id, "name")
VALUES(13, 'Đề 13');
INSERT INTO testsuite
(id, "name")
VALUES(14, 'Đề 14');
INSERT INTO testsuite
(id, "name")
VALUES(15, 'Đề 15');
INSERT INTO testsuite
(id, "name")
VALUES(16, 'Đề 16');
INSERT INTO testsuite
(id, "name")
VALUES(17, 'Đề 17');
INSERT INTO testsuite
(id, "name")
VALUES(18, 'Đề 18');

INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(76, 'aaaaaaa', 'A', true, 1, 'A', 'B', '', '', '20200527094752-351f_wm_20210211123843.jpg');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(77, '
“Phương tiện tham gia giao thông đường bộ” gồm những loại nào? ', 'B', false, 1, 'Phương tiện giao thông cơ giới đường bộ.', 'Phương tiện giao thông thô sơ đường bộ và xe máy chuyên dùng.', 'Cả ý 1 và ý 2.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(78, '
Người lái xe không được lùi xe ở những khu vực nào dưới đây? ', 'C', false, 1, 'Ở khu vực cho phép đỗ xe.', 'Ở khu vực cấm dừng và trên phần đường dành cho người đi bộ qua đường.', '', 'Cả ý 2 và ý 3.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(79, '
Biển báo hiệu có dạng tam giác đều, viền đỏ, nền màu vàng, trên có hình vẽ màu đen là loại biển gì dưới đây? 

', 'A', false, 1, 'Biển báo nguy hiểm.', 'Biển báo cấm.', '', 'Biển báo chỉ dẫn.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(80, '
Việc sát hạch cấp giấy phép lái xe ô tô phải thực hiện ở đâu và như thế nào? ', 'A', false, 1, 'Tại các cơ sở đào tạo lái xe có đủ điều kiện và phải bảo đảm công khai, minh bạch.', 'Tại sân tập lái của cơ sở đào tạo lái xe và phải đảm bảo công khai, minh bạch.', 'Tại các trung tâm sát hạch lái xe có đủ điều kiện hoạt động và phải bảo đảm công khai, minh bạch.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(81, '
Trên đường bộ (trừ đường cao tốc) trong khu vực đông dân cư, đường hai chiều không có dải phân cách giữa, xe mô tô hai bánh, ô tô chở người đến 30 chỗ tham gia giao thông với tốc độ tối đa cho phép là bao nhiêu? ', 'A', false, 1, '60 km/h.', '50 km/h.', '40 km/h.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(82, '
Trên đường bộ ngoài khu vực đông dân cư, đường hai chiều không có dải phân cách giữa; đường một chiều có một làn xe cơ giới (trừ đường cao tốc), loại xe nào tham gia giao thông với tốc độ tối đa cho phép là 80 km/h? ', 'A', false, 1, 'Ô tô kéo rơ moóc, ô tô kéo xe khác, xe gắn máy.', 'Ô tô chở người trên 30 chỗ (trừ ô tô buýt), ô tô tải có trọng tải trên 3.500 kg.', 'Xe ô tô con, xe ô tô chở người đến 30 chỗ (trừ xe buýt), ô tô tải có trọng tải nhỏ hơn hoặc bằng 3.500 kg.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(83, '
Khi tham gia giao thông trên đoạn đường không có biển báo “Cự ly tối thiểu giữa hai xe”, với điều kiện mặt đường khô ráo, xe cơ giới đang chạy với tốc độ từ trên 60 km/h đến 80 km/h, người lái xe phải duy trì khoảng cách an toàn với xe đang chạy phía trước tối thiểu là bao nhiêu? ', 'A', false, 1, '35 m.', '55 m.', '70 m.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(84, '
Tại nơi đường giao nhau, khi đèn điều khiển giao thông có tín hiệu màu vàng, người điều khiển phương tiện giao thông phải chấp hành như thế nào là đúng quy tắc giao thông? ', 'A', false, 1, 'Phải cho xe dừng lại trước vạch dừng, trường hợp đã đi quá vạch dừng hoặc đã quá gần vạch dừng nếu dừng lại thấy nguy hiểm thì được đi tiếp.', 'Trong trường hợp tín hiệu vàng nhấp nháy là được đi nhưng phải giảm tốc độ, chú ý quan sát nhường đường cho người đi bộ qua đường.', '', 'Cả ý 1 và ý 2.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(85, '
Người ngồi trên xe mô tô hai bánh, ba bánh, xe gắn máy khi tham gia giao thông có được mang, vác vật cồng kềnh hay không? ', 'A', false, 1, 'Được mang, vác tùy trường hợp cụ thể.', 'Không được mang, vác.', '', 'Được mang vác tùy theo sức khỏe của bản thân.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(86, '
Xe vận chuyển hàng nguy hiểm phải chấp hành những quy định nào dưới đây? ', 'A', false, 1, 'Phải có giấy phép do cơ quan có thẩm quyền cấp, xe vận chuyển hàng nguy hiểm không được dừng, đỗ nơi đông người, những nơi dễ xảy ra nguy hiểm.', 'Phải được chuyên chở trên xe chuyên dùng để vận chuyển hàng nguy hiểm; xe vận chuyển hàng nguy hiểm phải chạy liên tục không được dừng, đỗ trong quá trình vận chuyển.', 'Cả ý 1 và ý 2.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(87, '
Hình nào dưới đây đòi hỏi hàng phải xếp theo hướng thẳng đứng? 

', 'A', false, 1, 'Hình 1.', 'Hình 2.', 'Hình 3.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(88, '
Khái niệm về văn hóa giao thông được hiểu như thế nào là đúng? ', 'A', false, 1, 'Là sự hiểu biết và chấp hành nghiêm chỉnh pháp luật về giao thông; là ý thức trách nhiệm với cộng đồng khi tham gia giao thông.', 'Là ứng xử có văn hóa, có tình yêu thương con người trong các tình huống không may xảy ra khi tham gia giao thông.', 'Cả ý 1 và ý 2.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(89, '
Khi lái xe ô tô trên mặt đường có nhiều “ổ gà”, người lái xe phải thực hiện thao tác như thế nào để đảm bảo an toàn? ', 'A', false, 1, 'Giảm tốc độ, về số thấp và giữ đều ga.', 'Tăng tốc độ cho xe lướt qua nhanh.', 'Tăng tốc độ, đánh lái liên tục để tránh “ổ gà”.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(90, '
Chủ phương tiện cơ giới đường bộ có được tự ý thay đổi màu sơn, nhãn hiệu hoặc các đặc tính kỹ thuật của phương tiện so với chứng nhận đăng ký xe hay không? ', 'A', false, 1, 'Được phép thay đổi bằng cách dán đề can với màu sắc phù hợp.', 'Không được phép thay đổi.', 'Tùy từng loại phương tiện cơ giới đường bộ.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(91, '
Biển số 2 có ý nghĩa như thế nào? 

', 'A', false, 1, 'Cho phép ô tô có tải trọng trục lớn hơn 7 tấn đi qua.', 'Cho phép ô tô có tải trọng trên trục xe từ 7 tấn trở xuống đi qua.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(92, '
Xe ô tô chở hàng vượt quá phía trước và sau thùng xe, mỗi phía quá 10% chiều dài toàn bộ thân xe, tổng chiều dài xe (cả hàng) từ trước đến sau nhỏ hơn trị số ghi trên biển thì có được phép đi vào không? 

', 'A', false, 1, 'Không được phép', 'Được phép', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(93, '
Biển nào báo hiệu “Giao nhau với đường ưu tiên”? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', '', 'Cả ba biển.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(94, '
Biển nào báo hiệu “Đường hai chiều”? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', 'Biển 3.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(95, '
Biển nào sau đây là biển “Đường trơn”? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', 'Biển 3.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(96, '
Biển nào dưới đây cảnh báo nguy hiểm đoạn đường thường xảy ra tai nạn? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', 'Biển 3.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(97, '
Gặp biển nào người tham gia giao thông phải đi chậm và thận trọng đề phòng khả năng xuất hiện và di chuyển bất ngờ của trẻ em trên mặt đường? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(98, '
Biển báo này có ý nghĩa gì? 

', 'A', false, 1, 'Báo hiệu đường có ổ gà, lồi lõm.', 'Báo hiệu đường có gờ giảm tốc phía trước.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(99, '
Biển nào dưới đây báo hiệu hết cấm vượt? 

', 'A', false, 1, 'Biển 1.', 'Biển 2.', '', 'Biển 2 và 3.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(100, '
Vạch kẻ đường nào dưới đây là vạch phân chia các làn xe cùng chiều? 

', 'A', false, 1, 'Vạch 1.', 'Vạch 2.', '', 'Vạch 1 và 2.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(101, '
Thứ tự các xe đi như thế nào là đúng quy tắc giao thông? 

', 'A', false, 1, 'Xe tải, xe khách, xe con, mô tô.', 'Xe tải, mô tô, xe khách, xe con.', '', 'Mô tô, xe khách, xe tải, xe con.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(102, '
Trường hợp này xe nào được quyền đi trước? 

', 'A', false, 1, 'Mô tô.', 'Xe con.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(103, '
Xe nào đỗ vi phạm quy tắc giao thông? 

', 'A', false, 1, 'Cả hai xe.', 'Không xe nào vi phạm.', '', 'Chỉ xe tải vi phạm.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(104, '
Theo hướng mũi tên, những hướng nào ô tô không được phép đi? 

', 'A', false, 1, 'Hướng 1 và 2.', 'Hướng 3.', '', 'Hướng 2 và 3.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(105, '
Xe nào vượt đúng quy tắc giao thông? 

', 'A', false, 1, 'Cả 2 xe đều đúng.', 'Xe con.', 'Xe khách.', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(106, '
Ô tô con đi theo chiều mũi tên có vi phạm quy tắc giao thông không? 

', 'A', false, 1, 'Không vi phạm.', 'Vi phạm.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(107, '
Theo hướng mũi tên, xe nào được phép đi? 

', 'A', false, 1, 'Mô tô, xe con.', 'Xe con, xe tải.', '', 'Cả ba xe.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(108, '
Theo hướng mũi tên, những hướng nào xe mô tô được phép đi? 

', 'A', false, 1, 'Cả ba hướng.', 'Hướng 1 và 2.', '', 'Hướng 2 và 3.', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(110, '
Trong tình huống dưới đây, xe con màu vàng vượt xe con màu đỏ là đúng quy tắc giao thông hay không? 

', 'A', false, 1, 'Đúng.', 'Không đúng.', '', '', '');
INSERT INTO question
(id, "name", "result", paralysis, id_code_test, answera, answerb, answerc, answerd, img)
VALUES(113, 'phong', 'A', true, 2, 'A', 'B', 'C', '', 'DSC_774820210211120431.jpg');
