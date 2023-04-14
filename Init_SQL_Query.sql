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
  username VARCHAR(50) NOT NULL,
  wilayah_id INT NOT NULL,
  cabang_id INT NOT NULL,
  user_privilege VARCHAR(50) NOT NULL,
  PRIMARY KEY(user_id, wilayah_id)
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
