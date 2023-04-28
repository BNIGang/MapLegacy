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
create the cabang table
*/
CREATE TABLE IF NOT EXISTS cabang (
  cabang_id CHAR(36) PRIMARY KEY,
  cabang_name VARCHAR(50) NOT NULL
);

/* 
insert cabang data to wilayah table
*/
INSERT INTO cabang (cabang_id, cabang_name)
VALUES 
  (UUID(), 'Palembang'),
  (UUID(), 'Tanjungkarang'),
  (UUID(), 'Musi Palembang'),
  (UUID(), 'Jambi'),
  (UUID(), 'Pangkalpinang'),
  (UUID(), 'Bengkulu'),
  (UUID(), 'Prabumulih'),
  (UUID(), 'Kayuagung'),
  (UUID(), 'Baturaja'),
  (UUID(), 'Lubuklinggau'),
  (UUID(), 'Bangko'),
  (UUID(), 'Muarabungo'),
  (UUID(), 'Kualatungkal'),
  (UUID(), 'Metro'),
  (UUID(), 'Kotabumi');

/* 
create the user_privileges table
*/
CREATE TABLE IF NOT EXISTS user_privileges (
  user_id CHAR(36) NOT NULL,
  wilayah_id INT NOT NULL,
  cabang_id CHAR(36) NOT NULL,
  user_privilege VARCHAR(50) NOT NULL,
  PRIMARY KEY(user_id, wilayah_id)
);

/* 
insert dummy data to user_privileges table
TODO: wilayah list not yet found
*/
INSERT INTO user_privileges (user_id, wilayah_id, cabang_id, user_privilege)
VALUES 
  ((SELECT user_id FROM users WHERE username='user1'), 1, (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'admin'),
  ((SELECT user_id FROM users WHERE username='user2'), 2, (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'pemimpin_wilayah'),
  ((SELECT user_id FROM users WHERE username='user3'), 3, (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'pemimpin_cabang'),
  ((SELECT user_id FROM users WHERE username='user4'), 4, (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'pemimpin_cabang_pembantu'),
  ((SELECT user_id FROM users WHERE username='user5'), 5, (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'individu'),
  ((SELECT user_id FROM users WHERE username='user6'), 1, (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'individu');

/* 
create the kota_kabupaten table
*/
CREATE TABLE IF NOT EXISTS kota_kabupaten (
  cabang_id CHAR(36) NOT NULL,
  kota_kabupaten_id CHAR(36) NOT NULL,
  kota_kabupaten_name VARCHAR(50) NOT NULL,
  PRIMARY KEY(cabang_id, kota_kabupaten_id)
);

/* 
insert data to cabang table for Palembang
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES 
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), UUID(), 'Kota Palembang'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), UUID(), 'Kabupaten Banyuasin'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), UUID(), 'Kabupaten Ogan Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), UUID(), 'Kabupaten Ogan Komering Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), UUID(), 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Tanjung Karang
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kota Bandar Lampung'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kabupaten Lampung Selatan'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kabupaten Pesawaran'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kabupaten Pringsewu'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kabupaten Tanggamus'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), UUID(), 'Kabupaten Tulang Bawang Barat');

/* 
insert data to cabang table for Musi Palembang
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kota Palembang'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kabupaten Musi Banyuasin'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kabupaten Banyuasin'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kabupaten Ogan Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kabupaten Ogan Komering Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), UUID(), 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Jambi
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), UUID(), 'Kota Jambi'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), UUID(), 'Kabupaten Batanghari'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), UUID(), 'Kabupaten Muaro Jambi'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), UUID(), 'Kabupaten Banyuasin');

/* 
insert data to cabang table for Pangkal Pinang
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kota Pangkalpinang'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Bangka'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Bangka Barat'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Bangka Selatan'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Bangka Tengah'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Belitung'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), UUID(), 'Kabupaten Belitung Timur');

/* 
insert data to cabang table for Bengkulu
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kota Bengkulu'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Bengkulu Selatan'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Bengkulu Tengah'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Bengkulu Utara'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Kaur'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Kepahiang'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Lebong'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Mukomuko'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Rejang Lebong'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), UUID(), 'Kabupaten Seluma');

/* 
insert data to cabang table for Prabumulih
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Prabumulih"), UUID(), 'Kabupaten Muara Enim'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Prabumulih"), UUID(), 'Kota Prabumulih'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Prabumulih"), UUID(), 'Kabupaten Penukal Abab Lematang Ilir');

/* 
insert data to cabang table for Kayuagung
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kayuagung"), UUID(), 'Kabupaten Ogan Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kayuagung"), UUID(), 'Kabupaten Ogan Komering Ilir'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kayuagung"), UUID(), 'Kabupaten Muara Enim');

/* 
insert data to cabang table for Baturaja
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), UUID(), 'Kabupaten Ogan Komering Ulu'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), UUID(), 'Kabupaten Ogan Komering Ulu Selatan'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), UUID(), 'Kabupaten Ogan Komering Ulu Timur'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), UUID(), 'Kabupaten Penukal Abab Lematang Ilir');

/* 
insert data to cabang table for Lubuklinggau
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kota Lubuklinggau'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kota Pagar Alam'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kabupaten Empat Lawang'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kabupaten Lahat'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kabupaten Musi Rawas'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), UUID(), 'Kabupaten Musi Rawas Utara');

/* 
insert data to cabang table for Bangko
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), UUID(), 'Kabupaten Merangin'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), UUID(), 'Kabupaten Sarolangun'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), UUID(), 'Kota Sungai Penuh'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), UUID(), 'Kabupaten Kerinci');

/* 
insert data to cabang table for Muarabungo
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), UUID(), 'Kabupaten Bungo'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), UUID(), 'Kabupaten Tebo');

/* 
insert data to cabang table for Kualatungkal
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kualatungkal"), UUID(), 'Kabupaten Tanjung Jabung Barat'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kualatungkal"), UUID(), 'Kabupaten Tanjung Jabung Timur');

/* 
insert data to cabang table for Metro
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), UUID(), 'Kota Metro'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), UUID(), 'Kabupaten Lampung Tengah'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), UUID(), 'Kabupaten Lampung Timur'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), UUID(), 'Kabupaten Mesuji'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), UUID(), 'Kabupaten Tulang Bawang');

/* 
insert data to cabang table for Kotabumi
*/
INSERT INTO kota_kabupaten (cabang_id, kota_kabupaten_id, kota_kabupaten_name)
VALUES
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), UUID(), 'Kabupaten Lampung Barat'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), UUID(), 'Kabupaten Lampung Utara'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), UUID(), 'Kabupaten Pesisir Barat'),
  ((SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), UUID(), 'Kabupaten Way Kanan');

/* 
create bidang_usaha
*/
CREATE TABLE bidang_usaha (
  bidang_id CHAR(36) PRIMARY KEY,
  bidang VARCHAR(50)
);

/* 
insert values into bidang_usaha
*/
INSERT INTO bidang_usaha (bidang_id, bidang) VALUES
  (UUID(), 'Perkebunan & Pabrik'),
  (UUID(), 'Perkebunan'),
  (UUID(), 'Pabrik (Pengolahan)'),
  (UUID(), 'Prinsipal/Distributor Bangunan'),
  (UUID(), 'Prinsipal/Distributor Consumer Goods'),
  (UUID(), 'Perdagangan/Pengepul Hasil Bumi'),
  (UUID(), 'Perdagangan Besar Lainnya'),
  (UUID(), 'Perhotelan'),
  (UUID(), 'Pertambangan'),
  (UUID(), 'Jasa-Jasa Dunia Usaha'),
  (UUID(), 'Konstruksi');

/* 
Create table for produk usaha
*/
CREATE TABLE produk_usaha (
  produk_id CHAR(36) PRIMARY KEY,
  bidang_usaha_id CHAR(36),
  usaha VARCHAR(50)
);

/* 
Add produk usaha
*/
INSERT INTO produk_usaha (produk_id, bidang_usaha_id, usaha) VALUES
  -- Perkebunan & Pabrik
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Sawit'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Karet'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Kopi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Pinang'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Jagung'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Singkong'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Padi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan & Pabrik"), 'Lada'),

  -- Perdagangan/Pengepul Hasil Bumi
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Sawit'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Karet'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Kopi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Pinang'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Jagung'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Singkong'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Padi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan/Pengepul Hasil Bumi"), 'Lada'),

  -- Perhotelan
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Novotel'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Swisbel'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Mercure'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Amaris'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Santika'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'The Zuri'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Aryaduta'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'The Alts'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Favehotel'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Radisson'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Batiqa'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Sheraton'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Sahid'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perhotelan"), 'Lainnya'),

  -- Prinsipal/Distributor Consumer Goods
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Unilever'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Garudafoods'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Artaboga'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Indofoods'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Nestle'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Wingsfood'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Mayora'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Orang Tua'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'SoGoods Foods'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Consumer Goods"), 'Lainnya'),

  -- Pertambangan
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pertambangan"), 'Batubara'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pertambangan"), 'Timah'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pertambangan"), 'Pasir'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pertambangan"), 'Batu'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pertambangan"), 'Lainnya'),

  -- Prinsipal/Distributor Bangunan
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Cat'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Keramik'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Seng'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Semen'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Pipa'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Prinsipal/Distributor Bangunan"), 'Besi'),

  -- Jasa-Jasa Dunia Usaha
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Jasa-Jasa Dunia Usaha"), 'Rumah Sakit'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Jasa-Jasa Dunia Usaha"), 'Pengangkutan Tambang'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Jasa-Jasa Dunia Usaha"), 'Pengangkutan Darat'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Jasa-Jasa Dunia Usaha"), 'Pengangkutan Laut'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Jasa-Jasa Dunia Usaha"), 'Ekspedisi'),

  -- Perdagangan Besar Lainnya
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Farmasi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Sembako/Campuran'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Pecah Belah'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Elektronik'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Groceries/Swalayan/Minimarket'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Bahan Bangunan'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perdagangan Besar Lainnya"), 'Lainnya'),

  -- Konstruksi
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Developer'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Konstruksi Jalan & Jembatan'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Konstruksi Gedung'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Konstruksi Kapal (Galangan)'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Konstruksi Elektrikal'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Konstruksi"), 'Lainnya'),

  -- Perkebunan
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Sawit'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Karet'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Kopi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Pinang'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Jagung'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Singkong'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Padi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Perkebunan"), 'Lada'),

  -- Pabrik (Pengolahan)
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Sawit'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Karet'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Kopi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Pinang'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Jagung'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Singkong'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Padi'),
  (UUID(), (SELECT bidang_id FROM bidang_usaha WHERE bidang="Pabrik (Pengolahan)"), 'Lada');

/* 
Create table for KCU_KCP_KK
*/
CREATE TABLE kantor (
  kantor_id CHAR(36) PRIMARY KEY,
  cabang_id CHAR(36),
  kantor VARCHAR(50)
);

/* 
Insert to KCU_KCP_KK table
*/
INSERT INTO kantor (kantor_id, cabang_id, kantor) VALUES
  -- Palembang
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCU PALEMBANG'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP UNSRI'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP A. Yani'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP Komperta'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP UNSRI Indralaya'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP Kenten'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP KM 12'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP Jembatan Ampera'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP Pasar 16 Ilir'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KCP Palembang Square'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK PIM Letkol Iskandar'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Dempo'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Lorong Basah'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Kertapati'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Plaju'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK PUSRI'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Musi II'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Palembang"), 'KK Demang Lebar Daun'),

  -- Tanjungkarang
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCU TANJUNGKARANG'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Teuku Umar'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Antasari'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Panjang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Kalianda'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Bandar Lampung'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Unila'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Pringsewu'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Natar'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP Talang Padang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP GEDONG TATAAN'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP SIDOMULYO'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KCP SUKARAME'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KK Majapahit'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KK Way Halim'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Tanjungkarang"), 'KK Universitas Malahayati'),

  -- Musi Palembang
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCU MUSI PALEMBANG'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Pasar Betung'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Sungai Lilin'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Sekayu'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Bayung Lencir'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Boom Baru'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Pal Lima'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Kalidoni'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Lemabang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Rajawali'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KCP Sako'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KK KM9'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KK MP Mangkunegara'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KK OPI Mall'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KK Kenten Laut'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Musi Palembang"), 'KK Tanjung Api-Api'),

  -- Jambi
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCU JAMBI'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCP Abadi'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCP Simpang Sipin'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCP The Hok'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KK Pattimura'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KK Selincah'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KK Talang Banjar'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCP Sengeti'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Jambi"), 'KCP Muara Bulian'),

  -- Pangkalpinang
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCU PANGKALPINANG'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Sungailiat'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Tanjung Pandan'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Koba'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Muntok'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Toboali'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KCP Manggar'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Pangkalpinang"), 'KK Bangka Trade Centre'),

  -- Bengkulu
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCU BENGKULU'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Curup'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Arga Makmur'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Pasar Panorama'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Kepahiang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Ketahun'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KK Universitas Bengkulu'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KK Pagar Dewa'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Penarik'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bengkulu"), 'KCP Bintuhan'),

  -- Prabumulih
  -- TODO: Soon

  -- Kayuagung
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kayuagung"), 'KC KAYU AGUNG'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kayuagung"), 'KCP Tugumulyo'),

  -- Baturaja
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), 'KCU BATURAJA'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), 'KCP Muara dua'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), 'KCP Belitang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Baturaja"), 'KCP Martapura'),

  -- Lubuklinggau
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCU LUBUKLINGGAU'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KK Simpang Periuk'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCP Lahat'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCP Pagar Alam'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCP MURATARA'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCP MUSI RAWAS'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Lubuklinggau"), 'KCP Empat Lawang'),

  -- Bangko
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), 'KCU BANGKO'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), 'KCP Hitam Ulu'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), 'KCP Sarolangun'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Bangko"), 'KK Singkut'),

  -- Muarabungo
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), 'KCU MUARA BUNGO'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), 'KCP Kuamang Kuning'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), 'KCP Jujuhan'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Muarabungo"), 'KCP Rimbo Bujang'),

  -- Kualatungkal
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kualatungkal"), 'KCU KUALA TUNGKAL'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kualatungkal"), 'KCP Muara Sabak'),

  -- Metro
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'KCU METRO'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'KCP Bandar Jaya'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'KCP Tulang Bawang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'KCP Way Jepara'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Metro"), 'KCP Mesuji'),

  -- Kotabumi
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KCU KOTABUMI'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KCP Bukit Kemuning'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KCP Liwa'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KCP Krui'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KK Bunga Mayang'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KK Daya Murni'),
  (UUID(), (SELECT cabang_id FROM cabang WHERE cabang_name="Kotabumi"), 'KCP Baradatu');

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
  kcu_kcp_kk VARCHAR(50) NOT NULL,
  nasabah VARCHAR(50) NOT NULL,
  no_CIF INT NOT NULL,
  aum_di_bni DECIMAL(18,2) NOT NULL,
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
Create dummy data 
*/
INSERT INTO data_nasabah (id, nama_pengusaha, nomor_kontak, alamat_tempat_tinggal, bidang_usaha, produk_usaha, detail_bidang_usaha, kabupaten_kota, cabang, kcu_kcp_kk, nasabah, no_CIF, aum_di_bni, debitur, kredit_di_bni, produk_bni_yang_dimiliki, mitra_bank_dominan, aum_di_bank_lain, kredit_di_bank_lain, afiliasi, hubungan_afiliasi, added_by)
VALUES 
  (UUID(), 'Khidr Karawita', '081234567890', 'Jl. Contoh No. 1', 'Perkebunan & Pabrik', 'Sawit', 'Menanam Sawit :D', 'Palembang', 'Kota Palembang', 'KCU Palembang', 'Nasabah', 12345, 100000000, 'Debitur', 0, 'M-Banking', 'Mandiri', 50000000, 10000000, 'Muhammad Sumbul', 'Anak', (SELECT user_id FROM users WHERE username='user1')),
  (UUID(), 'Yaqub Qomarudin Dibizah', '081234567891', 'Jl. Contoh No. 2', 'Perhotelan', 'Novotel', 'Pemilik Novotel', 'Musi Palembang', 'Kabupaten Ogan Ilir', 'KCU Musi Palembang', 'Nasabah', 12346, 200000000, 'Debitur', 0, 'Giro', 'BCA', 10000000, 5000000, '', '', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Khalid Kashmiri', '081234567892', 'Jl. Contoh No. 3', 'Pertambangan', 'Batubara', 'Saya menambang batubara', 'Jambi', 'Kota Jambi', 'KCU Jambi', 'Non Nasabah', 12347, 300000000, 'Non Debitur', 0, 'Deposito', 'BTN', 20000000, 10000000, '', '', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Ismail Ahmad Khan Kabawi', '081234567893', 'Jl. Contoh No. 4', 'Prinsipal/Distributor Bangunan', 'Keramik', 'Penjual Keramik Nomor 1', 'Baturaja', 'Kabupaten Ogan Komering Ulu Selatan', 'KCP BATURAJA', 'Non Nasabah', 12348, 400000000, 'Debitur', 0, 'Kredit BB', 'Danamon', 30000000, 20000000, 'Joseph Joestar', 'Adik', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Utsman Abdul Jalil Shisha', '081234567894', 'Jl. Contoh No. 5', 'Perkebunan', 'Kopi', 'Pemilik Pt Kopi Indonesia', 'Bengkulu', 'Kota Bengkulu', 'KCU BENGKULU', 'Nasabah', 12349, 500000000, 'Non Debitur', 0, 'Tabungan', 'BSI', 40000000, 30000000, '', '', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Dio Brando', '081234567890', 'Jl. Contoh No. 3', 'Jasa-Jasa Dunia Usaha', 'Rumah Sakit', 'Pemilik Rumah Sakit Nusantara', 'Tanjungkarang', 'Kota Bandar Lampung', 'KCU TANJUNGKARANG', 'Nasabah', 12347, 500000000, 'Non Debitur', 0, 'Giro', 'Maybank', 40000000, 30000000, '', '', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Jotaro Kujo', '081234567896', 'Jl. Contoh No. 5', 'Konstruksi', 'Developer', 'Pengdevelop handal', 'Lubuklinggau', 'Kota Pagar Alam', 'KCU LUBUKLINGGAU', 'Non Debitur', 12349, 600000000, 'Tidak', 0, 'Kredit BB', 'Mega', 40000000, 10000000, '', '', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Giorno Giovanna', '081234567899', 'Jl. Contoh No. 6', 'Perdagangan Besar Lainnya', 'Farmasi', 'Pemilik PT Farmasi Sriwijaya', 'Palembang', 'Kota Palembang', 'KCU Palembang', 'Debitur', 12350, 700000000, 'Tidak', 0, 'Tapenas', 'Panin', 20000000, 15000000, '', '', (SELECT user_id FROM users WHERE username='user4'));