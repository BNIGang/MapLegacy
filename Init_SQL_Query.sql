/* 
create the bni_map_legacy database
*/
CREATE DATABASE IF NOT EXISTS bni_map_legacy;

/* 
switch to the bni_map_legacy database
*/
USE bni_map_legacy;

/* 
create the users table
*/
CREATE TABLE IF NOT EXISTS users (
  user_id CHAR(36) PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL
);

/* 
create the user_privileges table
*/
CREATE TABLE IF NOT EXISTS user_privileges (
  user_id CHAR(36) NOT NULL,
  wilayah_id INT NOT NULL,
  cabang_id INT NOT NULL,
  user_privilege VARCHAR(50) NOT NULL,
  PRIMARY KEY(user_id, wilayah_id),
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);

/* 
create the wilayah table
*/
CREATE TABLE IF NOT EXISTS wilayah (
  wilayah_id INT PRIMARY KEY,
  wilayah_name VARCHAR(50) NOT NULL
);

/* 
create the cabang table
*/
CREATE TABLE IF NOT EXISTS cabang (
  wilayah_id INT NOT NULL,
  cabang_id INT NOT NULL,
  cabang_name VARCHAR(50) NOT NULL,
  PRIMARY KEY(wilayah_id, cabang_id)
);

/* 
insert dummy data to users table
ALL encrpyed pass here is simply "pass"
*/
INSERT INTO users (user_id, username, password)
VALUES 
  (UUID(), 'user1', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user2', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user3', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user4', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user5', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user6', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy');

/* 
create the data_nasabah table
*/
CREATE TABLE IF NOT EXISTS data_nasabah (
  id CHAR(36) PRIMARY KEY,
  nama_pengusaha VARCHAR(50) NOT NULL,
  nomor_kontak VARCHAR(20) NOT NULL,
  alamat_tempat_tinggal VARCHAR(100) NOT NULL,
  bidang_usaha VARCHAR(50) NOT NULL,
  produk_usaha VARCHAR(50) NOT NULL,
  detail_bidang_usaha VARCHAR(100),
  kabupaten_kota VARCHAR(50) NOT NULL,
  cabang VARCHAR(50) NOT NULL,
  KCU_KCP_KK VARCHAR(50) NOT NULL,
  nasabah VARCHAR(50) NOT NULL,
  no_CIF INT NOT NULL,
  AUM_di_BNI DECIMAL(18,2) NOT NULL,
  debitur VARCHAR(50) NOT NULL,
  kredit_di_bni DECIMAL(18,2) NOT NULL,
  produk_bni_yang_dimiliki VARCHAR(50) NOT NULL,
  mitra_bank_dominan VARCHAR(50) NOT NULL,
  aum_di_bank_lain DECIMAL(18,2) NOT NULL,
  kredit_di_bank_lain DECIMAL(18,2) NOT NULL,
  afiliasi VARCHAR(50) NOT NULL,
  hubungan_afiliasi VARCHAR(50) NOT NULL,
  added_by CHAR(36) NOT NULL
);

/* 
insert dummy data to user_privileges table
*/
INSERT INTO user_privileges (user_id, wilayah_id, cabang_id, user_privilege)
VALUES 
  ((SELECT user_id FROM users WHERE username = 'user1'), 1, 1, 'admin'),
  ((SELECT user_id FROM users WHERE username = 'user2'), 2, 2, 'pemimpin_wilayah'),
  ((SELECT user_id FROM users WHERE username = 'user3'), 3, 3, 'pemimpin_cabang'),
  ((SELECT user_id FROM users WHERE username = 'user4'), 4, 4, 'pemimpin_cabang_pembantu'),
  ((SELECT user_id FROM users WHERE username = 'user5'), 5, 5, 'individu'),
  ((SELECT user_id FROM users WHERE username = 'user6'), 1, 2, 'individu');

/* 
insert dummy data to wilayah table
*/
INSERT INTO wilayah (wilayah_id, wilayah_name)
VALUES 
  (1, 'Palembang'),
  (2, 'Tanjungkarang'),
  (3, 'Musi Palembang'),
  (4, 'Jambi'),
  (5, 'Pangkalpinang'),
  (6, 'Bengkulu'),
  (7, 'Prabumulih'),
  (8, 'Kayuagung'),
  (9, 'Baturaja'),
  (10, 'Lubuklinggau'),
  (11, 'Bangko'),
  (12, 'Muarabungo'),
  (13, 'Kualatungkal'),
  (14, 'Metro'),
  (15, 'Kotabumi');

/* 
insert data to cabang table for Palembang
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES 
  (1, 1, 'Kota Palembang'),
  (1, 2, 'Kabupaten Banyuasin'),
  (1, 3, 'Kabupaten Ogan Ilir'),
  (1, 4, 'Kabupaten Ogan Komering Ilir'),
  (1, 5, 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Tanjung Karang
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (2, 1, 'Kota Bandar Lampung'),
  (2, 2, 'Kabupaten Lampung Selatan'),
  (2, 3, 'Kabupaten Pesawaran'),
  (2, 4, 'Kabupaten Pringsewu'),
  (2, 5, 'Kabupaten Tanggamus'),
  (2, 6, 'Kabupaten Tulang Bawang Barat');

/* 
insert data to cabang table for Musi Palembang
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (3, 1, 'Kota Palembang'),
  (3, 2, 'Kabupaten Musi Banyuasin'),
  (3, 3, 'Kabupaten Banyuasin'),
  (3, 4, 'Kabupaten Ogan Ilir'),
  (3, 5, 'Kabupaten Ogan Komering Ilir'),
  (3, 6, 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Jambi
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (4, 1, 'Kota Jambi'),
  (4, 2, 'Kabupaten Batanghari'),
  (4, 3, 'Kabupaten Muaro Jambi'),
  (4, 4, 'Kabupaten Banyuasin');

/* 
insert data to cabang table for Pangkal Pinang
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (5, 1, 'Kota Pangkalpinang'),
  (5, 2, 'Kabupaten Bangka'),
  (5, 3, 'Kabupaten Bangka Barat'),
  (5, 4, 'Kabupaten Bangka Selatan'),
  (5, 5, 'Kabupaten Bangka Tengah'),
  (5, 6, 'Kabupaten Belitung'),
  (5, 7, 'Kabupaten Belitung Timur');

/* 
insert data to cabang table for Bengkulu
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (6, 1, 'Kota Bengkulu'),
  (6, 2, 'Kabupaten Bengkulu Selatan'),
  (6, 3, 'Kabupaten Bengkulu Tengah'),
  (6, 4, 'Kabupaten Bengkulu Utara'),
  (6, 5, 'Kabupaten Kaur'),
  (6, 6, 'Kabupaten Kepahiang'),
  (6, 7, 'Kabupaten Lebong'),
  (6, 8, 'Kabupaten Mukomuko'),
  (6, 9, 'Kabupaten Rejang Lebong'),
  (6, 10, 'Kabupaten Seluma');

/* 
insert data to cabang table for Prabumulih
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (7, 1, 'Kabupaten Muara Enim'),
  (7, 2, 'Kota Prabumulih'),
  (7, 3, 'Kabupaten Penukal Abab Lematang Ilir');

/* 
insert data to cabang table for Kayuagung
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (8, 1, 'Kabupaten Ogan Ilir'),
  (8, 1, 'Kabupaten Ogan Komering Ilir'),
  (8, 1, 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Baturaja
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (9, 1, 'Kabupaten Ogan Komering Ulu'),
  (9, 2, 'Kabupaten Ogan Komering Ulu Selatan'),
  (9, 3, 'Kabupaten Ogan Komering Ulu Timur'),
  (9, 4, 'Kabupaten Penukal Abab Lematang Ilir');

/* 
insert data to cabang table for Lubuklinggau
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (10, 1, 'Kota Lubuklinggau'),
  (10, 2, 'Kota Pagar Alam'),
  (10, 3, 'Kabupaten Empat Lawang'),
  (10, 4, 'Kabupaten Lahat'),
  (10, 5, 'Kabupaten Musi Rawas'),
  (10, 6, 'Kabupaten Musi Rawas Utara');

/* 
insert data to cabang table for Bangko
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (11, 1, 'Kabupaten Merangin'),
  (11, 2, 'Kabupaten Sarolangun'),
  (11, 3, 'Kota Sungai Penuh'),
  (11, 4, 'Kabupaten Kerinci');

/* 
insert data to cabang table for Muarabungo
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (12, 1, 'Kabupaten Bungo'),
  (12, 2, 'Kabupaten Tebo');

/* 
insert data to cabang table for Kualatungkal
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (13, 1, 'Kabupaten Tanjung Jabung Barat'),
  (13, 2, 'Kabupaten Tanjung Jabung Timur');

/* 
insert data to cabang table for Metro
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (14, 1, 'Kota Metro'),
  (14, 2, 'Kabupaten Lampung Tengah'),
  (14, 3, 'Kabupaten Lampung Timur'),
  (14, 4, 'Kabupaten Mesuji'),
  (14, 5, 'Kabupaten Tulang Bawang');

/* 
insert data to cabang table for Kotabumi
*/
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (15, 1, 'Kabupaten Lampung Barat'),
  (15, 2, 'Kabupaten Lampung Utara'),
  (15, 3, 'Kabupaten Pesisir Barat'),
  (15, 4, 'Kabupaten Way Kanan');

/* 
create bidang_usaha
*/
CREATE TABLE bidang_usaha (
  bidang_id INT PRIMARY KEY,
  bidang VARCHAR(50)
);

/* 
insert values into bidang_usaha
*/
INSERT INTO bidang_usaha (bidang_id, bidang) VALUES
  (1, 'Perkebunan & Pabrik'),
  (2, 'Perkebunan'),
  (3, 'Pabrik (Pengolahan)'),
  (4, 'Prinsipal/Distributor Bangunan'),
  (5, 'Prinsipal/Distributor Consumer Goods'),
  (6, 'Perdagangan/Pengepul Hasil Bumi'),
  (7, 'Perdagangan Besar Lainnya'),
  (8, 'Perhotelan'),
  (9, 'Pertambangan'),
  (10, 'Jasa-Jasa Dunia Usaha'),
  (11, 'Konstruksi');

/* 
Create table for produk usaha
*/
CREATE TABLE produk_usaha (
  produk_id INT PRIMARY KEY,
  bidang_usaha_id INT,
  usaha VARCHAR(50),
  FOREIGN KEY (bidang_usaha_id) REFERENCES bidang_usaha(bidang_id)
);

/* 
Add produk usaha
*/
INSERT INTO produk_usaha (produk_id, bidang_usaha_id, usaha) VALUES
  -- Perkebunan & Pabrik
  (1, 1, 'Sawit'),
  (2, 1, 'Karet'),
  (3, 1, 'Kopi'),
  (4, 1, 'Pinang'),
  (5, 1, 'Jagung'),
  (6, 1, 'Singkong'),
  (7, 1, 'Padi'),
  (8, 1, 'Lada'),

  -- Perdagangan/Pengepul Hasil Bumi
  (9, 6, 'Sawit'),
  (10, 6, 'Karet'),
  (11, 6, 'Kopi'),
  (12, 6, 'Pinang'),
  (13, 6, 'Jagung'),
  (14, 6, 'Singkong'),
  (15, 6, 'Padi'),
  (16, 6, 'Lada'),

  -- Perhotelan
  (17, 8, 'Novotel'),
  (18, 8, 'Swisbel'),
  (19, 8, 'Mercure'),
  (20, 8, 'Amaris'),
  (21, 8, 'Santika'),
  (22, 8, 'The Zuri'),
  (23, 8, 'Aryaduta'),
  (24, 8, 'The Alts'),
  (25, 8, 'Favehotel'),
  (26, 8, 'Radisson'),
  (27, 8, 'Batiqa'),
  (28, 8, 'Sheraton'),
  (29, 8, 'Sahid'),
  (30, 8, 'Lainnya'),

  -- Prinsipal/Distributor Consumer Goods
  (31, 5, 'Unilever'),
  (32, 5, 'Garudafoods'),
  (33, 5, 'Artaboga'),
  (34, 5, 'Indofoods'),
  (35, 5, 'Nestle'),
  (36, 5, 'Wingsfood'),
  (37, 5, 'Mayora'),
  (38, 5, 'Orang Tua'),
  (39, 5, 'SoGoods Foods'),
  (40, 5, 'Lainnya'),

  -- Jasa-Jasa Dunia Usaha
  (41, 10, 'Batubara'),
  (42, 10, 'Timah'),
  (43, 10, 'Pasir'),
  (44, 10, 'Batu'),
  (45, 10, 'Lainnya'),

  -- Prinsipal/Distributor Bangunan
  (46, 4, 'Cat'),
  (47, 4, 'Keramik'),
  (48, 4, 'Seng'),
  (49, 4, 'Semen'),
  (50, 4, 'Pipa'),
  (51, 4, 'Besi'),

  -- Pertambangan
  (52, 9, 'Rumah Sakit'),
  (53, 9, 'Pengangkutan Tambang'),
  (54, 9, 'Pengangkutan Darat'),
  (55, 9, 'Pengangkutan Laut'),
  (56, 9, 'Ekspedisi'),

  -- Perdagangan Besar Lainnya
  (57, 7, 'Farmasi'),
  (58, 7, 'Sembako/Campuran'),
  (59, 7, 'Pecah Belah'),
  (60, 7, 'Elektronik'),
  (61, 7, 'Groceries/Swalayan/Minimarket'),
  (62, 7, 'Bahan Bangunan'),
  (63, 7, 'Lainnya'),

  -- Konstruksi
  (64, 11, 'Developer'),
  (65, 11, 'Konstruksi Jalan & Jembatan'),
  (66, 11, 'Konstruksi Gedung'),
  (67, 11, 'Konstruksi Kapal (Galangan)'),
  (68, 11, 'Konstruksi Elektrikal'),
  (69, 11, 'Lainnya'),

  -- Perkebunan
  (70, 2, 'Sawit'),
  (71, 2, 'Karet'),
  (72, 2, 'Kopi'),
  (73, 2, 'Pinang'),
  (74, 2, 'Jagung'),
  (75, 2, 'Singkong'),
  (76, 2, 'Padi'),
  (77, 2, 'Lada'),

  -- Pabrik (Pengolahan)
  (78, 3, 'Sawit'),
  (79, 3, 'Karet'),
  (80, 3, 'Kopi'),
  (81, 3, 'Pinang'),
  (82, 3, 'Jagung'),
  (83, 3, 'Singkong'),
  (84, 3, 'Padi'),
  (85, 3, 'Lada');

/* 
Create table for KCU_KCP_KK
*/
CREATE TABLE KCU_KCP_KK (
  kantor_id INT PRIMARY KEY,
  wilayah_id INT,
  kantor VARCHAR(50),
  FOREIGN KEY (wilayah_id) REFERENCES wilayah(wilayah_id)
);

/* 
Insert to KCU_KCP_KK table
*/

INSERT INTO KCU_KCP_KK (kantor_id, wilayah_id, kantor) VALUES
  -- Palembang
  (1, 1, 'KCU PALEMBANG'),
  (2, 1, 'KCP UNSRI'),
  (3, 1, 'KCP A. Yani'),
  (4, 1, 'KCP Komperta'),
  (5, 1, 'KCP UNSRI Indralaya'),
  (6, 1, 'KCP Kenten'),
  (7, 1, 'KCP KM 12'),
  (8, 1, 'KCP Jembatan Ampera'),
  (9, 1, 'KCP Pasar 16 Ilir'),
  (10, 1, 'KCP Palembang Square'),
  (11, 1, 'KK PIM Letkol Iskandar'),
  (12, 1, 'KK Dempo'),
  (13, 1, 'KK Lorong Basah'),
  (14, 1, 'KK Kertapati'),
  (15, 1, 'KK Plaju'),
  (16, 1, 'KK PUSRI'),
  (17, 1, 'KK Musi II'),
  (18, 1, 'KK Demang Lebar Daun'),

  -- Tanjungkarang
  (19, 2, 'KCU TANJUNGKARANG'),
  (20, 2, 'KCP Teuku Umar'),
  (21, 2, 'KCP Antasari'),
  (22, 2, 'KCP Panjang'),
  (23, 2, 'KCP Kalianda'),
  (24, 2, 'KCP Bandar Lampung'),
  (25, 2, 'KCP Unila'),
  (26, 2, 'KCP Pringsewu'),
  (27, 2, 'KCP Natar'),
  (28, 2, 'KCP Talang Padang'),
  (29, 2, 'KCP GEDONG TATAAN'),
  (30, 2, 'KCP SIDOMULYO'),
  (31, 2, 'KCP SUKARAME'),
  (32, 2, 'KK Majapahit'),
  (33, 2, 'KK Way Halim'),
  (34, 2, 'KK Universitas Malahayati'),

  -- Musi Palembang
  (35, 3, 'KCU MUSI PALEMBANG'),
  (36, 3, 'KCP Pasar Betung'),
  (37, 3, 'KCP Sungai Lilin'),
  (38, 3, 'KCP Sekayu'),
  (39, 3, 'KCP Bayung Lencir'),
  (40, 3, 'KCP Boom Baru'),
  (41, 3, 'KCP Pal Lima'),
  (42, 3, 'KCP Kalidoni'),
  (43, 3, 'KCP Lemabang'),
  (44, 3, 'KCP Rajawali'),
  (45, 3, 'KCP Sako'),
  (46, 3, 'KK KM9'),
  (47, 3, 'KK MP Mangkunegara'),
  (48, 3, 'KK OPI Mall'),
  (49, 3, 'KK Kenten Laut'),
  (50, 3, 'KK Tanjung Api-Api'),

  -- Jambi
  (51, 4, 'KCU JAMBI'),
  (52, 4, 'KCP Abadi'),
  (53, 4, 'KCP Simpang Sipin'),
  (54, 4, 'KCP The Hok'),
  (55, 4, 'KK Pattimura'),
  (56, 4, 'KK Selincah'),
  (57, 4, 'KK Talang Banjar'),
  (58, 4, 'KCP Sengeti'),
  (59, 4, 'KCP Muara Bulian'),

  -- Pangkalpinang
  (60, 5, 'KCU PANGKALPINANG'),
  (61, 5, 'KCP Sungailiat'),
  (62, 5, 'KCP Tanjung Pandan'),
  (63, 5, 'KCP Koba'),
  (64, 5, 'KCP Muntok'),
  (65, 5, 'KCP Toboali'),
  (66, 5, 'KCP Manggar'),
  (67, 5, 'KK Bangka Trade Centre'),

  -- Bengkulu
  (68, 6, 'KCU BENGKULU'),
  (69, 6, 'KCP Curup'),
  (70, 6, 'KCP Arga Makmur'),
  (71, 6, 'KCP Pasar Panorama'),
  (72, 6, 'KCP Kepahiang'),
  (73, 6, 'KCP Ketahun'),
  (74, 6, 'KK Universitas Bengkulu'),
  (75, 6, 'KK Pagar Dewa'),
  (76, 6, 'KCP Penarik'),
  (77, 6, 'KCP Bintuhan'),

  -- Prabumulih
  -- Soon

  -- Kayuagung (id: 8)
  (78, 8, 'KC KAYU AGUNG'),
  (79, 8, 'KCP Tugumulyo'),

  -- Baturaja
  (80, 9, 'KCU BATURAJA'),
  (81, 9, 'KCP Muara dua'),
  (82, 9, 'KCP Belitang'),
  (83, 9, 'KCP Martapura'),

  -- Lubuklinggau
  (84, 10, 'KCU LUBUKLINGGAU'),
  (85, 10, 'KK Simpang Periuk'),
  (86, 10, 'KCP Lahat'),
  (87, 10, 'KCP Pagar Alam'),
  (88, 10, 'KCP MURATARA'),
  (89, 10, 'KCP MUSI RAWAS'),
  (90, 10, 'KCP Empat Lawang'),

  -- Bangko
  (91, 11, 'KCU BANGKO'),
  (92, 11, 'KCP Hitam Ulu'),
  (93, 11, 'KCP Sarolangun'),
  (94, 11, 'KK Singkut'),

  -- Muarabungo
  (95, 12, 'KCU MUARA BUNGO'),
  (96, 12, 'KCP Kuamang Kuning'),
  (97, 12, 'KCP Jujuhan'),
  (98, 12, 'KCP Rimbo Bujang'),

  -- Kualatungkal
  (99, 13, 'KCU KUALA TUNGKAL'),
  (100, 13, 'KCP Muara Sabak'),

  -- Metro
  (101, 14, 'KCU METRO'),
  (102, 14, 'KCP Bandar Jaya'),
  (103, 14, 'KCP Tulang Bawang'),
  (104, 14, 'KCP Way Jepara'),
  (105, 14, 'KCP Mesuji'),

  -- Kotabumi
  (106, 15, 'KCU KOTABUMI'),
  (107, 15, 'KCP Bukit Kemuning'),
  (108, 15, 'KCP Liwa'),
  (109, 15, 'KCP Krui'),
  (110, 15, 'KK Bunga Mayang'),
  (111, 15, 'KK Daya Murni'),
  (112, 15, 'KCP Baradatu');

/* 
Create dummy data 
*/
INSERT INTO data_nasabah (id, nama_pengusaha, nomor_kontak, alamat_tempat_tinggal, bidang_usaha, produk_usaha, detail_bidang_usaha, kabupaten_kota, cabang, KCU_KCP_KK, nasabah, no_CIF, AUM_di_BNI, debitur, kredit_di_bni, produk_bni_yang_dimiliki, mitra_bank_dominan, aum_di_bank_lain, kredit_di_bank_lain, afiliasi, hubungan_afiliasi, added_by)
VALUES 
  (UUID(), 'Khidr Karawita', '081234567890', 'Jl. Contoh No. 1', 'Perkebunan & Pabrik', 'Sawit', 'Menanam Sawit :D', 'Palembang', 'Kota Palembang', 'KCU Palembang', 'Nasabah', 12345, 100000000, 'Debitur', 0, 'M-Banking', 'Mandiri', 50000000, 10000000, 'Muhammad Sumbul', 'Anak', (SELECT user_id FROM users WHERE username='user1')),
  (UUID(), 'Yaqub Qomarudin Dibizah', '081234567891', 'Jl. Contoh No. 2', 'Perhotelan', 'Novotel', 'Pemilik Novotel', 'Musi Palembang', 'Kabupaten Ogan Ilir', 'KCU Musi Palembang', 'Nasabah', 12346, 200000000, 'Debitur', 0, 'Giro', 'BCA', 10000000, 5000000, '', '', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Khalid Kashmiri', '081234567892', 'Jl. Contoh No. 3', 'Pertambangan', 'Batubara', 'Saya menambang batubara', 'Jambi', 'Kota Jambi', 'KCU Jambi', 'Non Nasabah', 12347, 300000000, 'Non Debitur', 0, 'Deposito', 'BTN', 20000000, 10000000, '', '', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Ismail Ahmad Khan Kabawi', '081234567893', 'Jl. Contoh No. 4', 'Prinsipal/Distributor Bangunan', 'Keramik', 'Penjual Keramik Nomor 1', 'Baturaja', 'Kabupaten Ogan Komering Ulu Selatan', 'KCP BATURAJA', 'Non Nasabah', 12348, 400000000, 'Debitur', 0, 'Kredit BB', 'Danamon', 30000000, 20000000, 'Joseph Joestar', 'Adik', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Utsman Abdul Jalil Shisha', '081234567894', 'Jl. Contoh No. 5', 'Perkebunan', 'Kopi', 'Pemilik Pt Kopi Indonesia', 'Bengkulu', 'Kota Bengkulu', 'KCU BENGKULU', 'Nasabah', 12349, 500000000, 'Non Debitur', 0, 'Tabungan', 'BSI', 40000000, 30000000, '', '', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Dio Brando', '081234567890', 'Jl. Contoh No. 3', 'Jasa-Jasa Dunia Usaha', 'Rumah Sakit', 'Pemilik Rumah Sakit Nusantara', 'Tanjungkarang', 'Kota Bandar Lampung', 'KCU TANJUNGKARANG', 'Nasabah', 12347, 500000000, 'Non Debitur', 0, 'Giro', 'Maybank', 40000000, 30000000, '', '', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Jotaro Kujo', '081234567896', 'Jl. Contoh No. 5', 'Konstruksi', 'Developer', 'Pengdevelop handal', 'Lubuklinggau', 'Kota Pagar Alam', 'KCU LUBUKLINGGAU', 'Non Debitur', 12349, 600000000, 'Tidak', 0, 'Kredit BB', 'Mega', 40000000, 10000000, '', '', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Giorno Giovanna', '081234567899', 'Jl. Contoh No. 6', 'Perdagangan Besar Lainnya', 'Farmasi', 'Pemilik PT Farmasi Sriwijaya', 'Palembang', 'Kota Palembang', 'KCU Palembang', 'Debitur', 12350, 700000000, 'Tidak', 0, 'Tapenas', 'Panin', 20000000, 15000000, '', '', (SELECT user_id FROM users WHERE username='user4'));