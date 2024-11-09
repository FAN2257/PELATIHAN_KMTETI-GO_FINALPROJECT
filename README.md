# Fadel Aulia Naldi (23/519144/TK/57236)

## Penjelasan API (Database name = `toko_buku`)
1. Book (`\api\book`) [Database collection = `book`]
    1) Get `?displayAll`
        API ini berfungsi untuk memanggil semua data buku yang ada dengan yang ditampilkan hanya `title`, `author`, dan `price`.
    2) Get `?displayDetails&id={id}`
        API ini berfungsi untuk melihat informasi secara detail pada salah satu buku yang dipilih sesuai `id`.
    3) Put book `?id={id}`
        API ini berfungsi untuk mengganti value `price` dan `stock` pada buku yang dipilih sesuai `id`.
    4) Post book
        API ini berfungsi untuk membuat data buku baru untuk disimpan di database.
    5) Delete book `?id={id}`
        API ini berfungsi untuk menghapus salah satu buku yang ingin dihapus dari database.

2. Employee (`\api\employee`) [Database collection = `employee`]
    1) Get employee
        API ini berfungsi untuk menampilkan seluruh data karyawan yang ada di database, namun hanya bagian `name`, `date_release`, dan `status`.
    2) Post employee
        API ini berfungsi untuk membuat data karyawan baru untuk dimasukkan di database.