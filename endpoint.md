# Endpoint API

## Endpoint Auth
Endpoint ini bertanggung jawab membuat generate token untuk token ke endpoint pengelolaan.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***POST*** | *`/api/users/login`* | Men-generate token untuk mengakses endpoint yang berfungsi sebagai pengelolaan. Token akan didapatkan setelah pengguna mengirim request json berupa email dan password yang sudah terdaftar
***POST*** | *`/api/users/register`* | Membuat akun user untuk akses endpoint

## Endpoint User
Endpoint ini bertanggung jawab mengelola user.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/users/:id`* | Mengakses data pengguna berdasakan id pengguna | token
***PUT*** | *`/api/users/:id`* | Mengubah data pengguna berdasakan id pengguna | token
***DELETE*** | *`/api/users/:id`* | Menghapus data pengguna berdasakan id pengguna | token

## Endpoint Photo
Endpoint ini bertanggung jawab mengelola photo.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/photos`* | Mengakses data foto | token
***GET*** | *`/api/photos/:id`* | Mengakses data pengguna berdasakan id foto | token
***POST*** | *`/api/photos`* | Membuat data foto baru | token
***PUT*** | *`/api/photos/:id`* | Mengubah data foto berdasakan id foto | token
***DELETE*** | *`/api/photos/:id`* | Menghapus data foto berdasakan id foto | token