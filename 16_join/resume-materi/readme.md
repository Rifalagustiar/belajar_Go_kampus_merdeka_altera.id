Join: Join digunakan untuk menggabungkan baris dari dua atau lebih tabel berdasarkan kolom yang terkait di antara mereka. Terdapat beberapa jenis join, seperti inner join, left join, right join, dan full outer join.

Union: Union digunakan untuk menggabungkan hasil dari dua atau lebih query dengan jumlah kolom yang sama menjadi satu set data yang berbeda.

Agregasi: Agregasi digunakan untuk menghitung nilai statistik pada kumpulan data, seperti jumlah, rata-rata, maksimum, minimum, dan lain sebagainya. Beberapa fungsi agregasi umum yang digunakan adalah COUNT, SUM, AVG, MAX, dan MIN.

Subquery: Subquery digunakan untuk mengeksekusi sebuah query di dalam query lainnya. Subquery dapat digunakan dalam berbagai bentuk, seperti WHERE, HAVING, atau FROM clause.

Function: Fungsi dalam DBMS dapat digunakan untuk melakukan operasi pada nilai atau kumpulan data dalam database. Beberapa contoh fungsi yang umum digunakan adalah fungsi matematika, fungsi string, fungsi tanggal dan waktu, dan lain sebagainya.
seperti contoh

SELECT _ FROM transactions WHERE user_id IN (1, 2);
SELECT SUM(total_price) FROM transactions WHERE user_id = 1;
SELECT COUNT(_) FROM transactions t
JOIN transaction_details td
ON t.id = td.transaction_id
JOIN products p
ON td.product_id = p.id WHERE p.product_type_id = 2;

    SELECT products.*, product_types.name
    FROM products
    JOIN product_types
    ON products.product_type_id = product_types.id;

    SELECT transactions.*, products.name AS product_name, users.name AS user_name
    FROM transactions
    INNER JOIN products ON transactions.id = products.id
    INNER JOIN users ON transactions.user_id = users.id;

    DELIMITER $$
    CREATE FUNCTION delete_transaction_detail_by_transaction_id(p_transaction_id INT)
    RETURNS BOOLEAN
    BEGIN
    DELETE FROM transaction_details WHERE transaction_id = p_transaction_id;
    RETURN TRUE;
    END$$
    DELIMITER ;
    SELECT delete_transaction_detail_by_transaction_id(4);

    DROP FUNCTION IF EXISTS `update_total_qty_by_transaction_detail_id`;
    DELIMITER $$
    CREATE FUNCTION update_total_qty_by_transaction_detail_id(p_transaction_detail_id INT)
    RETURNS BOOLEAN
    BEGIN
    DECLARE transaction_id INT;
    DECLARE qty INT;
    SELECT transaction_id, qty INTO transaction_id, qty FROM transaction_details WHERE transaction_id = p_transaction_detail_id;
    UPDATE transactions SET total_qty = total_qty - qty WHERE id = transaction_id;
    RETURN TRUE;
    END$$
    DELIMITER ;
    SELECT update_total_qty_by_transaction_detail_id(10);


    SELECT * FROM products WHERE id NOT IN (SELECT DISTINCT product_id FROM transaction_detail);
