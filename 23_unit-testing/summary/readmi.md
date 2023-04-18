========================= Go Testing Framework ================
Go memiliki built-in testing framework yang memungkinkan pengembang untuk menulis dan menjalankan unit test dengan mudah. Go testing framework dapat dijalankan dengan menggunakan perintah go test pada direktori yang berisi file-file test.

======================== Test Functions ====================
Test function adalah fungsi-fungsi yang dibuat untuk melakukan pengujian pada fungsi atau komponen tertentu. Test function harus diberi nama dengan awalan "Test" dan memiliki satu parameter bertipe \*testing.T sebagai objek testing yang akan digunakan untuk melakukan asserasi.

================ Assertion ===========
Assertion adalah cara untuk memastikan bahwa output dari fungsi atau kode yang diuji sesuai dengan yang diharapkan. Go menyediakan package testing yang memiliki beberapa metode asserasi bawaan seperti assert.Equal() dan assert.NoError(). Assertion yang dilakukan dengan benar memastikan bahwa program berjalan dengan benar dan membantu menemukan masalah dan bug dengan cepat.
