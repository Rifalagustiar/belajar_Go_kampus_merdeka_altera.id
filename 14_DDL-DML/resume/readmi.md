DDL adalah bahasa SQL yang digunakan untuk mendefinisikan struktur tabel database, seperti membuat tabel baru, mengubah struktur tabel, dan menghapus tabel. Beberapa perintah DDL yang dapat digunakan pada Golang antara lain:

CREATE TABLE: perintah ini digunakan untuk membuat tabel baru pada database:
sql

CREATE TABLE users (
id INT PRIMARY KEY,
name VARCHAR(50) NOT NULL,
email VARCHAR(50) NOT NULL
);
ALTER TABLE: perintah ini digunakan untuk mengubah struktur tabel, seperti menambah kolom baru, mengubah tipe data kolom, atau menghapus kolom:
sql

ALTER TABLE users
ADD phone VARCHAR(15);
DROP TABLE: perintah ini digunakan untuk menghapus tabel dari database:
sql

DROP TABLE users;
DML adalah bahasa SQL yang digunakan untuk memanipulasi data pada tabel database, seperti menambah, mengubah, atau menghapus data. Beberapa perintah DML yang dapat digunakan pada Golang antara lain:

INSERT INTO: perintah ini digunakan untuk menambah data baru ke dalam tabel:
sql

INSERT INTO users (id, name, email, phone)
VALUES (1, 'rifal', 'rifal@gmail.com.com', '1234567890');
UPDATE: perintah ini digunakan untuk mengubah data yang sudah ada pada tabel:
sql

UPDATE users
SET phone = '082286272340'
WHERE id = 1;
DELETE FROM: perintah ini digunakan untuk menghapus data dari tabel:
sql

DELETE FROM users
WHERE id = 1;
