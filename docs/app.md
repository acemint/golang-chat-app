## Pengembangan Aplikasi Chat dengan Golang

Sebuah perusahaan Chat yang berkembang pesat ingin membangun aplikasi web baru yang lebih cepat, skalabel, dan mudah dipelihara untuk menangani jutaan pengguna dan transaksi setiap hari. Setelah mengevaluasi berbagai bahasa pemrograman, mereka memilih Go (Golang) karena keunggulannya dalam hal performa dan dukungan bawaan untuk konkurensi.

## Tujuan
Membangun backend API yang cepat dan responsif.
Mengelola transaksi dan data pengguna dengan efisien.
Menjamin skalabilitas untuk menangani beban pengguna yang tinggi.

## Tools yang diwajibkan:
`Gin` digunakan sebagai framework web karena performanya yang tinggi dan kemudahan penggunaannya.
`GORM` dipilih sebagai ORM (Object-Relational Mapping) untuk memudahkan interaksi dengan database.


-Aplikasi dibagi menjadi beberapa layanan mikro (microservices) seperti autentikasi, manajemen produk, dan pemrosesan pembayaran.
-HTTP/2 digunakan untuk komunikasi antar layanan guna meningkatkan efisiensi.

## Management Memory:
`pprof`

Dependency Management:
`Go Modules` digunakan untuk mengelola paket-paket dan dependensi proyek dengan lebih baik.

Konkurensi:
`Goroutine` digunakan untuk menangani permintaan pengguna secara paralel, seperti mengelola banyak permintaan pencarian produk secara bersamaan.

Testing:
`GoConvey`

HTTP Server:
`net/http` package

Apakah ada penggunaan `Channel` dalam project ini?
Apa itu `Channel`?
