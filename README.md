# Student Portal Final DB

## Description

Student Portal ini adalah aplikasi Web Server API yang digunakan untuk mengakses data dari repositori student dan classes dengan melakukan autentikasi dengan data user. API ini memungkinkan pengguna untuk melakukan berbagai operasi CRUD (Create, Read, Update, Delete) pada data student.

API ini memiliki tiga endpoint yaitu: `/user`, `/student` dan `/class`.

- **Endpoint `/user`** digunakan untuk manajemen autentikasi user.
- **Endpoint `/student`** digunakan untuk mengakses data student.
- **Endpoint `/class`** digunakan untuk mengakses data class.

Pada setiap endpoint, API ini memiliki beberapa sub-endpoint yang berfungsi untuk melakukan operasi CRUD pada data. Sub-endpoint tersebut meliputi:

### /user

- `/register`: untuk mendaftarkan user baru di aplikasi
- `/login`: untuk masuk ke aplikasi menggunakan user yang telah terdaftar
- `/logout`: untuk keluar dari aplikasi

### /student

- `/get-all`: untuk mengambil semua data student
- `/get`: untuk mengambil data dengan ID tertentu
- `/add`: untuk menambahkan data baru
- `/update`: untuk memperbarui data yang sudah ada
- `/delete`: untuk menghapus data yang sudah ada
- `/get-with-class`: untuk mengambil semua data student beserta dengan detail class-nya

### /class

- `/get-all`: untuk mengambil semua data class

API ini dapat dijalankan dengan memanggil fungsi `Start()`, yang akan menampilkan pesan di console bahwa server sedang berjalan dan menjalankan server pada `http://localhost:8080`.

## Database Model and Schema

Aplikasi ini memiliki 4 tabel utama, yaitu `users`, `sessions`, `students` dan `classes`. Tabel `users` digunakan untuk menyimpan data-data user, tabel `sessions` digunakan untuk menyimpan data sesi token pada saat user login, tabel `students` digunakan untuk menyimpan data student dan tabel `classes` digunakan untuk menyimpan data-data class.

- Tabel `users` hanya dapat memiliki satu `sessions`, dan tabel `sessions` dapat memiliki banyak `users`. Tabel `users` dan `sessions` memiliki relasi one-to-many.
- Tabel `students` memiliki relasi one-to-many dengan tabel `classes`, dimana banyak siswa dapat terdaftar pada satu kelas. Kolom `class_id` pada tabel `students` merupakan foreign key yang mengacu pada primary key `id` pada tabel `classes`.

Note: aplikasi ini menggunakan GORM untuk manajemen data repository ke database PostgreSQL.

## Technologies Used

- **API**: Menyediakan endpoint untuk operasi CRUD pada data user, student, dan class.
- **ORM**: Menggunakan GORM untuk manajemen data repository.
- **Postman**: Digunakan untuk pengujian API.
- **JWT**: Digunakan untuk autentikasi dan otorisasi user.
- **PostgreSQL**: Database yang digunakan untuk menyimpan data.

### 📁 Repository

Ini adalah fungsi yang berinteraksi dengan database PostgreSQL.

#### user: repository/user.go

- **method `Add(user model.User)`**: menerima parameter bertipe `model.User` dan berfungsi menyimpan data sesuai parameter tersebut ke tabel `users`
- **method `CheckAvail(user model.User)`**: menerima parameter bertipe `model.User` dan berfungsi memeriksa ketersediaan data pada tabel `users` dengan ketentuan:
  - check berdasarkan field `username` dan `password` dari parameter yang diterima.
  - kembalikan error jika tidak ada
  - kembalikan nil jika ada

#### session: repository/session.go

- **method `AddSessions(session model.Session)`**: menerima parameter bertipe `model.Session` dan berfungsi menyimpan data sesuai parameter tersebut ke tabel `sessions`
- **method `UpdateSessions(session model.Session)`**: menerima parameter bertipe `model.Session` dan berfungsi mengubah data session sesuai parameter tersebut ke tabel `sessions` dengan kondisi sama antara username parameter dengan database.
- **method `DeleteSession(token string)`**: menerima parameter bertipe string dan berfungsi menghapus data tabel `sessions` sesuai dengan target token dari parameter yang diterima.
- **method `SessionAvailToken(token string)`**: menerima parameter bertipe string dan berfungsi memeriksa apakah token tersedia pada tabel `sessions` sesuai dengan kolom token sama dengan nilai dari parameter.
  - jika session ditemukan, maka kembalikan data session dalam bentuk `model.Session` dan error nil
  - jika session tidak ditemukan, maka kembalikan data session kosong dalam bentuk `model.Session{}` dan error message
- **method `SessionAvailName(name string)`**: menerima parameter bertipe string dan berfungsi memeriksa apakah token tersedia pada tabel `sessions` sesuai dengan kolom name sama dengan nilai dari parameter.
  - jika session ditemukan, maka kembalikan data session dalam bentuk `model.Session` dan error nil
  - jika session tidak ditemukan, maka kembalikan data session kosong dalam bentuk `model.Session{}` dan error message

#### student: repository/student.go

- **FetchAll**: Function ini akan mengambil semua data mahasiswa yang ada di dalam tabel `students` pada database. Selanjutnya, data mahasiswa tersebut akan di-scan dan dimasukkan ke dalam slice `[]model.Student`.
  - Jika proses tersebut berhasil, function akan mengembalikan slice tersebut beserta nilai nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan nil sebagai slice dan error yang terjadi.
- **FetchByID**: Function ini akan mengambil data mahasiswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query untuk mencari data mahasiswa dengan id yang sesuai. Hasil dari query tersebut akan di-scan ke dalam variabel `model.Student`.
  - Jika proses tersebut berhasil, function akan mengembalikan pointer `model.Student` beserta nilai nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan nil sebagai pointer dan error yang terjadi.
- **Store**: Function ini akan menyimpan data mahasiswa yang diberikan sebagai argumen ke dalam database. Pertama-tama, function akan mengeksekusi sebuah query `INSERT` untuk memasukkan data mahasiswa baru ke dalam tabel `students`. Query tersebut akan menggunakan nilai dari variabel `model.Student` yang diberikan sebagai argumen.
  - Jika proses tersebut berhasil, function akan mengembalikan nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan error yang terjadi.
- **Update**: Function ini akan mengupdate data mahasiswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query `UPDATE` untuk mengubah data mahasiswa dengan id yang sesuai. Query tersebut akan menggunakan nilai dari variabel `model.Student` dan id yang diberikan sebagai argumen.
  - Jika proses tersebut berhasil, function akan mengembalikan nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan error yang terjadi.
- **Delete**: Function ini akan menghapus data mahasiswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query `DELETE` untuk menghapus data mahasiswa dengan id yang sesuai dari tabel `students`.
  - Jika proses tersebut berhasil, function akan mengembalikan nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan error yang terjadi.
- **FetchWithClass**: Function ini akan mengambil semua data mahasiswa yang ada di dalam tabel `students` dan tabel `classes` dengan melakukan `JOIN` pada kedua tabel tersebut. Kemudian, data tersebut akan di-scan dan dimasukkan ke dalam slice `[]model.StudentClass`.
  - Jika proses tersebut berhasil, function akan mengembalikan slice tersebut beserta nilai nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan nil sebagai slice dan error yang terjadi.

#### class: repository/class.go

- **FetchAll**: Function ini akan mengambil semua data kelas yang ada di dalam tabel `classes` pada database. Selanjutnya, data kelas tersebut akan dimasukkan ke dalam slice `[]model.Class`.
  - Jika proses tersebut berhasil, function akan mengembalikan slice tersebut beserta nilai nil sebagai error.
  - Namun jika terjadi error pada proses tersebut, function akan mengembalikan nil sebagai slice dan error yang terjadi.

## Project Structure

```bash
.
├── README.md
├── api
│   ├── api.go
│   ├── class.go
│   ├── middleware.go
│   ├── student.go
│   └── user.go
├── assets
│   └── md
│       └── fcp-student-portal.png
├── assignment-config.json
├── db
│   └── postgres.go
├── go.mod
├── go.sum
├── golang_suite_test.go
├── main.go
├── main_test.go
├── model
│   └── model.go
├── repository
│   ├── class.go
│   ├── session.go
│   ├── student.go
│   └── user.go
└── service
    ├── class.go
    ├── session.go
    ├── student.go
    └── user.go

7 directories, 23 files

```
