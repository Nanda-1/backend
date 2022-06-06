package structs

import "time"

type Buku struct {
	IDBuku          int           `grom:"primary_key;auto_increment" json:"id_buku"`
	ISBN            string        `grom:"size:255" json:"isbn"`
	IDKategoriJenis int           `grom:"size:255" json:"id_kategori_jenis"`
	KategoriJenis   Jenis_Buku    `grom:"foreignkey:IDKategoriJenis" json:"kategori_jenis"`
	Judul           string        `grom:"size:255" json:"judul"`
	IDPenulisBuku   int           `grom:"size:255" json:"id_penulis_buku"`
	PenulisBuku     Penulis_Buku  `grom:"foreignkey:IDPenulisBuku" json:"penulis_buku"`
	IDPenerbitBuku  int           `grom:"size:255" json:"id_penerbit_buku"`
	PenerbitBuku    Penerbit_Buku `grom:"foreignkey:IDPenerbitBuku" json:"penerbit_buku"`
	TahunTerbit     string        `grom:"size:255" json:"tahun_terbit"`
	StokBuku        int           `grom:"size:255" json:"stok_buku"`
	RakBuku         string        `grom:"size:255" json:"rak_buku"`
	Deskripsi       string        `grom:"size:255" json:"deskripsi"`
	Gambar          string        `grom:"size:255" json:"gambar"`
	Status          string        `grom:"size:255" json:"status"`
	CreatedAT       time.Time     `grom:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAT       time.Time     `grom:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Detail_Buku struct {
	IDDetailBuku int    `grom:"primary_key;auto_increment" json:"id_detail_buku"`
	IDBuku       int    `grom:"size:255" json:"id_buku"`
	Buku         []Buku `grom:"foreignkey:IDBuku" json:"buku"`
	Gambarbuku   string `grom:"size:255" json:"gambar_buku"`
	Kondisi      string `grom:"size:255" json:"kondisi"`
}

type Jenis_Buku struct {
	IDJenis   int    `grom:"primary_key;auto_increment" json:"id_jenis"`
	JenisBuku string `grom:"size:255" json:"jenis_buku"`
	Deskripsi string `grom:"type:text;null" json:"deskripsi"`
}

type Penulis_Buku struct {
	IDPenulis     int    `grom:"primary_key;auto_increment" json:"id_penulis"`
	PenulisBuku   string `grom:"size:255" json:"penulis_buku"`
	AlamatPenulis string `grom:"size:255" json:"alamat_penulis"`
	EmailPenulis  string `grom:"size:255" json:"email_penulis"`
	Deskripsi     string `grom:"type:text;null" json:"deskripsi"`
}

type Penerbit_Buku struct {
	IDPenerbit     int    `grom:"primary_key;auto_increment" json:"id_penerbit"`
	PenerbitBuku   string `grom:"size:255" json:"penerbit_buku"`
	AlamatPenerbit string `grom:"size:255" json:"alamat_penerbit"`
	TelpPenerbit   string `grom:"size:255" json:"telp_penerbit"`
	EmailPenerbit  string `grom:"size:255" json:"email_penerbit"`
	Deskripsi      string `grom:"type:text;null" json:"deskripsi"`
}
