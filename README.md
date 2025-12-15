# Aplikasi Inventaris CLI
### Mini Project App Inventaris CLI oleh Fathoni

**Video Demo:** [`  Video-Demo-App-Inventaris-CLI   `](https://youtu.be/7EKqL2QxGkg)

## Tentang Aplikasi
Aplikasi Inventaris CLI merupakan aplikasi untuk membantu pengguna mengelola data kategori dan barang inventaris melalui terminal. Sistem ini dilengkapi dengan fitur CRUD, sehingga pengguna dapat menambah, melihat, mengupdate, dan menghapus data kategori maupun barang pada database. Fitur utama aplikasi ini mencakup pengelolaan kategori, manajemen data barang inventaris, pengecekan barang yang perlu diganti, laporan nilai investasi dan depresiasi, serta pencarian barang berdasarkan keyword yang dimasukkan.

#### Home Page
Halaman Home akan menampilkan 6 menu utama yang dapat dipilih oleh pengguna. 6 menu terdiri dari fitur kategori, fitur manajemen barang, daftar barang yang harus diganti, laporan nilai investasi dan depresiasi, pencarian barang, dan keluar dari program.
````bash
=== INVENTORY APP MENU ===
1) Item Category Features
2) Inventory Management Features
3) List Items Must Be Replaced
4) Investment and Depreciation Report
5) Search Items by Keyword
6) Exit
````

#### Item Category Features
Fitur kategori berguna untuk pengelolaan daftar kategori yang tersedia. Memiliki 5 submenu, yaitu lihat semua daftar kategori, menambah kategori, melihat detail kategori berdasarkan ID, mengedit kategori, dan menghapus kategori.
````bash
===  Item Category Features  ===
1. View List Categories
2. Add New Category
3. View Detail Category by ID
4. Edit Category
5. Delete Category
6. Back to Home
````

#### Inventory Management Features
Fitur manajemen barang berguna untuk pengelolaan daftar barang yang tersedia. Memiliki 5 submenu, yaitu lihat semua daftar barang, menambah barang, melihat detail barang berdasarkan ID, mengedit barang, dan menghapus barang.
````bash
=== Inventory Management Features ===
1. View List Items
2. Add New Item
3. View Detail Item by ID
4. Edit Item
5. Delete Item
6. Back to Home
````

#### List Items Must Be Replaced
Menu ini akan menampilkan daftar barang yang harus diganti, dengan ketentuan total penggunaannya sudah lebih dari 100 hari dan status 'replaced' bernilai True.

#### Investment and Depreciation Report
Fitur laporan investasi dan depresiasi berguna untuk menampilkan total nilai investasi dan depresiasi semua barang. Selain itu fitur ini juga memiliki submenu tampilkan nilai investasi dan depresiasi setiap barang. Perhitungan investasi dan depresiasi menggunakan metode depresiasi saldo menurun, sebesar 20% per tahun.
````bash
=== Investment and Depreciation Report ===
1. View Total Invest
2. View Invest Value All items
3. View Invest Value Item by ID
4. Back to Home
````

#### Search Items by Keyword
Menu ini berguna untuk melakukan pencarian barang berdasarkan keyword yang diinputkan

Terima kasih telah menggunakan Aplikasi Inventaris CLI
`Happy Coding ❤️🩵`