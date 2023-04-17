-- create the bni_map_legacy database
CREATE DATABASE IF NOT EXISTS bni_map_legacy;

-- switch to the bni_map_legacy database
USE bni_map_legacy;

-- create the users table
CREATE TABLE IF NOT EXISTS users (
  user_id CHAR(36) PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL
);

-- create the user_privileges table
CREATE TABLE IF NOT EXISTS user_privileges (
  user_id CHAR(36) NOT NULL,
  wilayah_id INT NOT NULL,
  cabang_id INT NOT NULL,
  user_privilege VARCHAR(50) NOT NULL,
  PRIMARY KEY(user_id, wilayah_id),
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);

-- create the wilayah table
CREATE TABLE IF NOT EXISTS wilayah (
  wilayah_id INT PRIMARY KEY,
  wilayah_name VARCHAR(50) NOT NULL
);

-- create the cabang table
CREATE TABLE IF NOT EXISTS cabang (
  wilayah_id INT NOT NULL,
  cabang_id INT NOT NULL,
  cabang_name VARCHAR(50) NOT NULL,
  PRIMARY KEY(wilayah_id, cabang_id)
);

-- insert dummy data to users table
-- encrpyed pass here is simply "pass"
INSERT INTO users (user_id, username, password)
VALUES 
  (UUID(), 'user1', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user2', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user3', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user4', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user5', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy'),
  (UUID(), 'user6', '$2a$10$.vnkTrMayTCzju1JXniwFe6vPkaKD/yQxatrwpW/DZqJRPY4/srPy');

-- create the data_nasabah table
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

-- insert dummy data to user_privileges table
INSERT INTO user_privileges (user_id, wilayah_id, cabang_id, user_privilege)
VALUES 
  ((SELECT user_id FROM users WHERE username = 'user1'), 1, 1, 'admin'),
  ((SELECT user_id FROM users WHERE username = 'user2'), 2, 2, 'pemimpin_wilayah'),
  ((SELECT user_id FROM users WHERE username = 'user3'), 3, 3, 'pemimpin_cabang'),
  ((SELECT user_id FROM users WHERE username = 'user4'), 4, 4, 'pemimpin_cabang_pembantu'),
  ((SELECT user_id FROM users WHERE username = 'user5'), 5, 5, 'individu'),
  ((SELECT user_id FROM users WHERE username = 'user6'), 1, 2, 'individu');

-- insert dummy data to wilayah table
INSERT INTO wilayah (wilayah_id, wilayah_name)
VALUES 
  (1, 'Palembang'),
  (2, 'Tanjungkarang'),
  (3, 'Musi Palembang'),
  (4, 'Jambi'),
  (5, 'Pangkalpinang');

-- insert data to cabang table for Palembang
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES 
  (1, 1, 'Kota Palembang'),
  (1, 2, 'Kabupaten Banyuasin'),
  (1, 3, 'Kabupaten Ogan Ilir'),
  (1, 4, 'Kabupaten Ogan Komering Ilir'),
  (1, 5, 'Kabupaten Muara Enim');

-- insert data to cabang table for Tanjung Karang
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (2, 1, 'Kota Bandar Lampung'),
  (2, 2, 'Kabupaten Lampung Selatan'),
  (2, 3, 'Kabupaten Pesawaran'),
  (2, 4, 'Kabupaten Pringsewu'),
  (2, 5, 'Kabupaten Tanggamus'),
  (2, 6, 'Kabupaten Tulang Bawang Barat');

-- insert data to cabang table for Musi Palembang
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (3, 1, 'Kota Palembang'),
  (3, 2, 'Kabupaten Musi Banyuasin'),
  (3, 3, 'Kabupaten Banyuasin'),
  (3, 4, 'Kabupaten Ogan Ilir'),
  (3, 5, 'Kabupaten Ogan Komering Ilir'),
  (3, 6, 'Kabupaten Muara Enim');

-- insert data to cabang table for Jambi
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (4, 1, 'Kota Jambi'),
  (4, 2, 'Kabupaten Batanghari'),
  (4, 3, 'Kabupaten Muaro Jambi'),
  (4, 4, 'Kabupaten Banyuasin');

-- insert data to cabang table for Pangkal Pinang
INSERT INTO cabang (wilayah_id, cabang_id, cabang_name)
VALUES
  (5, 1, 'Kota Pangkalpinang'),
  (5, 2, 'Kabupaten Bangka'),
  (5, 3, 'Kabupaten Bangka Barat'),
  (5, 4, 'Kabupaten Bangka Selatan'),
  (5, 5, 'Kabupaten Bangka Tengah'),
  (5, 6, 'Kabupaten Belitung'),
  (5, 7, 'Kabupaten Belitung Timur');

-- create bidang_usaha
CREATE TABLE bidang_usaha (
  id INT PRIMARY KEY,
  bidang VARCHAR(50)
);

-- insert values into bidang_usaha
INSERT INTO bidang_usaha (id, bidang) VALUES
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
  (11, 'Konsutruksi');

-- create produk usaha
CREATE TABLE produk_usaha (
  id INT PRIMARY KEY,
  bidang_usaha_id INT,
  usaha VARCHAR(50),
  FOREIGN KEY (bidang_usaha_id) REFERENCES bidang_usaha(id)
);

-- Add produk_usaha
INSERT INTO produk_usaha (id, bidang_usaha_id, usaha) VALUES
  (1, 1, 'Sawit'),
  (2, 1, 'Karet'),
  (3, 1, 'Kopi'),
  (4, 1, 'Pinang'),
  (5, 1, 'Jagung'),
  (6, 1, 'Singkong'),
  (7, 1, 'Padi'),
  (8, 1, 'Lada'),
  (9, 6, 'Sawit'),
  (10, 6, 'Karet'),
  (11, 6, 'Kopi'),
  (12, 6, 'Pinang'),
  (13, 6, 'Jagung'),
  (14, 6, 'Singkong'),
  (15, 6, 'Padi'),
  (16, 6, 'Lada'),
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
  (41, 10, 'Batubara'),
  (42, 10, 'Timah'),
  (43, 10, 'Pasir'),
  (44, 10, 'Batu'),
  (45, 10, 'Lainnya'),
  (46, 4, 'Cat'),
  (47, 4, 'Keramik'),
  (48, 4, 'Seng'),
  (49, 4, 'Semen'),
  (50, 4, 'Pipa'),
  (51, 4, 'Besi'),
  (52, 9, 'Rumah Sakit'),
  (53, 9, 'Pengangkutan Tambang'),
  (54, 9, 'Pengangkutan Darat'),
  (55, 9, 'Pengangkutan Laut'),
  (56, 9, 'Ekspedisi'),
  (57, 7, 'Farmasi'),
  (58, 7, 'Sembako/Campuran'),
  (59, 7, 'Pecah Belah'),
  (60, 7, 'Elektronik'),
  (61, 7, 'Groceries/Swalayan/Minimarket'),
  (62, 7, 'Bahan Bangunan'),
  (63, 7, 'Lainnya'),
  (64, 11, 'Developer'),
  (65, 11, 'Konstruksi Jalan & Jembatan'),
  (66, 11, 'Konstruksi Gedung'),
  (67, 11, 'Konstruksi Kapal (Galangan)'),
  (68, 11, 'Konstruksi Elektrikal'),
  (69, 11, 'Lainnya');

-- Change this later idk lmao
INSERT INTO data_nasabah (id, nama_pengusaha, nomor_kontak, alamat_tempat_tinggal, bidang_usaha, produk_usaha, detail_bidang_usaha, kabupaten_kota, cabang, KCU_KCP_KK, nasabah, no_CIF, AUM_di_BNI, debitur, kredit_di_bni, produk_bni_yang_dimiliki, mitra_bank_dominan, aum_di_bank_lain, kredit_di_bank_lain, afiliasi, hubungan_afiliasi, added_by)
VALUES 
  (UUID(), 'Nasabah 1', '081234567890', 'Jl. Contoh No. 1', 'Pertanian', 'Pupuk', 'Pertanian Umum', 'Kabupaten A', 'Cabang A', 'KCU A', 'Perusahaan', 12345, 100000000, 'Tidak', 0, 'Tabungan BNI', 'Mandiri', 50000000, 10000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user1')),
  (UUID(), 'Nasabah 2', '081234567891', 'Jl. Contoh No. 2', 'Pertanian', 'Pupuk', 'Pertanian Umum', 'Kabupaten B', 'Cabang B', 'KCU B', 'Perusahaan', 12346, 200000000, 'Tidak', 0, 'Tabungan BNI', 'BCA', 10000000, 5000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Nasabah 3', '081234567892', 'Jl. Contoh No. 3', 'Perdagangan', 'Elektronik', 'Pusat Perbelanjaan', 'Kota C', 'Cabang C', 'KCU C', 'Perorangan', 12347, 300000000, 'Tidak', 0, 'Kartu Kredit BNI', 'Mandiri', 20000000, 10000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user5')),
  (UUID(), 'Nasabah 4', '081234567893', 'Jl. Contoh No. 4', 'Pertanian', 'Tanaman Obat', 'Pertanian Umum', 'Kabupaten A', 'Cabang A', 'KCU A', 'Perusahaan', 12348, 400000000, 'Tidak', 0, 'Tabungan BNI', 'BCA', 30000000, 20000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Nasabah 5', '081234567894', 'Jl. Contoh No. 5', 'Pertanian', 'Pupuk', 'Pertanian Umum', 'Kabupaten B', 'Cabang B', 'KCU B', 'Perorangan', 12349, 500000000, 'Tidak', 0, 'Kartu Kredit BNI', 'Mandiri', 40000000, 30000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user3')),
  (UUID(), 'Nasabah 6', '081234567890', 'Jl. Contoh No. 3', 'Pertanian', 'Padi', 'Pertanian Umum', 'Kabupaten A', 'Cabang A', 'KCU A', 'Perusahaan', 12347, 500000000, 'Tidak', 0, 'Tabungan BNI', 'Mandiri', 40000000, 30000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Nasabah 7', '081234567896', 'Jl. Contoh No. 5', 'Pendidikan', 'Bimbingan Belajar', 'Bimbingan Belajar', 'Kabupaten B', 'Cabang B', 'KCU B', 'Perusahaan', 12349, 600000000, 'Tidak', 0, 'Tabungan BNI', 'BCA', 40000000, 10000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user2')),
  (UUID(), 'Nasabah 8', '081234567899', 'Jl. Contoh No. 6', 'Pendidikan', 'Kursus Bahasa', 'Kursus Bahasa', 'Kabupaten B', 'Cabang B', 'KCU B', 'Perusahaan', 12350, 700000000, 'Tidak', 0, 'Tabungan BNI', 'Mandiri', 20000000, 15000000, 'Tidak', 'Tidak', (SELECT user_id FROM users WHERE username='user4'));