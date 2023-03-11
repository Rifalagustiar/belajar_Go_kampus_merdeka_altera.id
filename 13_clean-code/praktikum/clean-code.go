package main

type user struct {
	id       int // tidak mengikuti konvebsi, biasanya dalam bahasa go id digunakan dengan hruf besar
	username int // tipe data pada username bukan int, tetapi string
	password int // tipe data pada password bukan int, tetapi string
}

// Pada userservice tidak disebutkan nama variabel atau field sehingga perlu diberikan nama yang jelas.
type userservice struct {
	t []user
}

/*
Pada method getallusers() dan getuserbyid(), receiver tidak menggunakan pointer,
sehingga ketika dilakukan perubahan pada userservice di luar method,
perubahan tidak akan terlihat pada instance userservice yang digunakan di dalam method.
*/
func (u userservice) getallusers() []user {
	return u.t
}

/*
Pada method getuserbyid(), ketika user tidak ditemukan,
method mengembalikan struct user dengan nilai default yang bukanlah tindakan yang baik,
karena bisa saja nilai default dari field-field pada user memberikan hasil yang salah.
*/
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}
