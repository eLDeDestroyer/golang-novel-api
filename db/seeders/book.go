package seeders

import (
	"e-novel/model"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

func SeedBook(db *gorm.DB) {
	books := []model.Book{
		{
			Title: "Kala Itu Langit Biru",
			Description: `Di sebuah kota kecil yang tenang, Awan kembali setelah sepuluh tahun merantau. Hanya ada satu alasan: kenangan. Kenangan tentang langit biru di suatu siang yang sunyi—dan tentang Biru, gadis yang dulu mengajarkannya bagaimana mencintai tanpa kata, dan melepaskan tanpa air mata.
			Saat masa lalu mulai mengetuk kembali lewat tempat-tempat yang mereka lalui bersama, Awan menemukan bahwa tak semua hal bisa dikembalikan ke tempat semula. Namun, langit tetap biru—seperti janjinya yang tak pernah diucap.
			“Kala Itu Langit Biru” adalah kisah tentang pertemuan, kehilangan, dan keberanian untuk melihat ke belakang agar bisa benar-benar melangkah ke depan.`,
			ImagePath: "https://imgs.search.brave.com/xF8cFw0Ja5b8lY0VwPGlzbR2sfCxOlJNtF_BZyVxhaw/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tYXJr/ZXRwbGFjZS5jYW52/YS5jb20vRUFHRU5V/VTZQeU0vMS8wLzEw/MDN3L2NhbnZhLWJp/cnUtbGFuZ2l0LW5v/dmVsLWJvb2stY292/ZXItMmU0aWdOeFVJ/Y1kuanBn",
			UserId: 1,
		},
		{
			Title: "Bersama Kamu",
			Description: `Bagi Langit, hidup selalu berjalan lurus dan teratur, sampai ia bertemu dengan Dara—gadis dengan senyum musim semi dan cara pandang yang mengubah segalanya. Di antara hujan sore, musik indie, dan tumpukan catatan kuliah yang tak selesai-selesai, tumbuhlah perasaan yang tak pernah ia rencanakan.
			“Bersama Kamu” adalah kisah tentang persahabatan yang menjelma jadi cinta, tentang memilih untuk tetap tinggal meski dunia meminta pergi. Sebuah perjalanan emosional untuk mengerti bahwa kadang, pulang bukanlah tempat—tapi seseorang.`,
			ImagePath: "https://imgs.search.brave.com/Pt-AbgfkF-uA9yjOilnp87y95jWNMj5vaeVey5_pR-c/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tYXJr/ZXRwbGFjZS5jYW52/YS5jb20vRUFGMzNG/Z05ydEEvMS8wLzEw/MDN3L2NhbnZhLWty/ZW0taGlqYXUtb3Jn/YW5pay1lc3RldGlr/LWNlcml0YS1jaW50/YS1yZW1hamEtYlVF/QzRGSUxyZkkuanBn",
			UserId: 1,
		},
		{
			Title: "Merajut Mimpi",
			Description: `Nara adalah anak penjual kain yang tinggal di pinggir pasar kota. Hidupnya sederhana, penuh warna, dan kadang diselimuti kekhawatiran tentang masa depan. Namun, ia punya satu impian yang tak pernah ia lepaskan: menjadi perancang busana yang bisa mempersembahkan karyanya ke panggung dunia.
			Suatu hari, ia bertemu dengan Elang, pemuda perantau yang diam-diam sedang mencari arti hidup setelah meninggalkan pekerjaan kantoran. Dari kain-kain yang dijahit dengan harapan hingga percakapan larut malam tentang hidup dan kehilangan, mereka mulai menenun sesuatu yang lebih besar dari sekadar mimpi.
			“Merajut Mimpi” adalah kisah tentang kerja keras, keberanian mengejar cita, dan menemukan makna hidup di sela-sela benang takdir.`,
			ImagePath: "https://imgs.search.brave.com/nhy4DaodPRfTlES-6FSc1cQ2SxbvrJMIOtnFCQmKcxA/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tYXJr/ZXRwbGFjZS5jYW52/YS5jb20vRUFHT1hD/Z2cxOW8vMS8wLzEw/MDN3L2NhbnZhLWJp/cnUta3VuaW5nLWx1/Y3UtYW5pbWFzaS1h/bmFrLWtlY2lsLXRl/bnRhbmctbWltcGkt/a292ZXItYnVrdS1W/eTFQaUxtR05uSS5q/cGc",
			UserId: 1,
		},
		{
			Title: "Petualangan Anna",
			Description: `Anna adalah gadis kecil berusia sembilan tahun yang gemar membaca buku-buku fantasi dan menjelajah alam sekitar rumahnya di desa pelangi. Suatu hari, saat mengikuti seekor kupu-kupu berwarna-warni, ia tersesat di sebuah hutan misterius yang tak pernah ia lihat sebelumnya—Hutan Pelangi. Tapi hutan ini bukan hutan biasa. Pohonnya bisa bicara, sungainya berkilau seperti kaca pelangi, dan hewan-hewannya menyimpan rahasia kuno.`,
			ImagePath: "https://imgs.search.brave.com/diXVa7dno23UIACj4jzDPPHTrENzX2NM_uhfTNPv12g/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tYXJr/ZXRwbGFjZS5jYW52/YS5jb20vRUFHcTJn/RC1RNTAvMS8wLzEw/MDN3L2NhbnZhLW1l/cmFoLWt1bmluZy1i/aXJ1LXdhcm5hLXdh/cm5pLWx1Y3UtbWVu/YXJpay1idWt1LWNl/cml0YS1hbmFrLXBl/dHVhbGFuZ2FuLWJv/b2stY292ZXItc1Mt/MVV4NUlLQjguanBn",
			UserId: 1,
		},
		{
			Title: "Hai, Nak",
			Description: `Sinta, seorang ibu muda yang baru saja kehilangan anaknya karena kecelakaan tragis, memutuskan pindah ke rumah tua di pinggir kota untuk memulihkan diri. Namun sejak malam pertama, ia mendengar suara-suara aneh dari kamar anak yang kosong—suara mainan bergemerincing, bisikan lirih, dan suara anak kecil yang berkata, "Hai, Nak."
			Awalnya Sinta mengira semua itu adalah halusinasinya sendiri. Namun perlahan, suara-suara itu menjadi nyata. Mainan bergerak sendiri. Lukisan keluarga bergeser. Dan yang paling mengerikan: suara anak kecil itu memanggilnya setiap malam, semakin dekat... semakin nyata.
			Apakah itu arwah anaknya? Ataukah sesuatu yang lebih gelap sedang mencoba masuk ke hidupnya melalui celah kehilangan?`,
			ImagePath: "https://imgs.search.brave.com/wsCC8eFWPsVK-P81dOnWlbTQ7bZYl6iOwRICjJQKDPY/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tYXJr/ZXRwbGFjZS5jYW52/YS5jb20vRUFHQ2p6/OWlIYTQvMS8wLzEx/MzF3L2NhbnZhLXBv/c3Rlci1maWxtLWhv/cm9yLXNlcmFtLW1v/ZGVybi1oaXRhbS1k/YW4tbWVyYWgtSEYw/THd0QngyTE0uanBn",
			UserId: 1,
		},
		{
			Title: "Indgigo Tapi Penakut",
			Description: `Aira, seorang gadis indigo berusia 15 tahun, pindah ke rumah tua warisan neneknya di desa terpencil Lereng Wana. Dengan kemampuan melihat makhluk gaib sejak kecil, ia terbiasa dengan dunia lain. Tapi rumah ini berbeda. Ada lorong yang tak pernah berujung, pintu terkunci dengan segel gaib, dan makhluk-makhluk yang tak hanya ingin dilihat—tapi juga ingin keluar. Dalam upayanya mengungkap rahasia rumah itu, Aira menemukan bahwa kemampuan indigonya bukan sekadar anugerah, melainkan kunci ke dalam petualangan gelap yang berbahaya dan menantang nyawanya.`,
			ImagePath: "https://imgs.search.brave.com/2q6BPnpwoHPSMIR31JqnVwHW9iNE-3nnHdjVzIni22A/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9jZG4u/Z3JhbWVkaWEuY29t/L3VwbG9hZHMvaXRl/bXMvSU5ESUdPX1RB/UElfUEVOQUtVVC5w/bmc",
			UserId: 1,
		},
		{
			Title: "Penjelajah Antariksa: Bencana di Planet Poe",
			Description: `Tahun 3127, manusia telah menjelajahi berbagai sudut galaksi. Kapten Elara, pemimpin muda dari armada ekspedisi bintang Orionis VII, ditugaskan menjelajahi planet misterius bernama Poe—planet yang dulunya dihuni peradaban tinggi, tapi kini sunyi tak berpenghuni. Namun sesampainya di sana, mereka menemukan bahwa Poe tidak mati… ia tertidur. Dan sesuatu sedang menunggu untuk terbangun. Dengan krunya yang terdiri dari ilmuwan, penjelajah, dan makhluk alien sahabat, Elara harus menghadapi misteri, bencana alam kosmik, dan kekuatan kuno yang bisa menghancurkan lebih dari sekadar planet. Ia harus memilih antara menyelamatkan misi… atau menyelamatkan alam semesta.`,
			ImagePath: "https://imgs.search.brave.com/tAiqMomYsFjkP3U1UogW58wTylkhFLzbk_4YA0hGJLc/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly93d3cu/Z3JhbWVkaWEuY29t/L2Jsb2cvY29udGVu/dC9pbWFnZXMvMjAy/MC8wNi9wZW5qZWxh/amFoLWFudGFyaWtz/YS0xLWJlbmNhbmEt/ZGktcGxhbmV0LXBv/YV9ncmFtZWRpYS5q/cGc",
			UserId: 1,
		},
		{
			Title: "Skripsick",
			Description: `Bayangkan hidup tenang di semester tua, saat semua teman sudah wisuda... kecuali kamu. Inilah kisah Gani, mahasiswa abadi jurusan Ilmu Komunikasi, yang sudah lima kali ganti judul skripsi dan tiga kali ganti dosen pembimbing. Dari perjuangan menghadapi dosen killer, printer ngadat, sampai drama dengan teman seperjuangan yang mulai buka bisnis kopi, skripsi bukan lagi soal penelitian — tapi soal bertahan hidup. Satu-satunya tujuan Gani: lulus tahun ini... atau kabur ke luar negeri.`,
			ImagePath: "https://imgs.search.brave.com/XTKV-RoHYr9i8YoecYivAAQ4LjtAfKlVivQLoVmO6MQ/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9jZG4x/LXByb2R1Y3Rpb24t/aW1hZ2VzLWtseS5h/a2FtYWl6ZWQubmV0/LzBUcEVBZ0VCZlg2/S2Ywb0NVcnFNSGVz/UXNucz0vNjQweDg1/My9zbWFydC9maWx0/ZXJzOnF1YWxpdHko/NzUpOnN0cmlwX2lj/YygpOmZvcm1hdCh3/ZWJwKS9rbHktbWVk/aWEtcHJvZHVjdGlv/bi9tZWRpYXMvNDY3/Njc0Mi9vcmlnaW5h/bC8wNzI5ODMxMDBf/MTcwMTg3Mjk4OC1O/b3ZlbF9sdWN1XzQu/anBn",
			UserId: 1,
		},
	}


	for _,book := range books {
		err := db.Table("books").Create(&book).Error
		if err != nil {
			spew.Dump(err)
		}
	}

	fmt.Println("Seeder Book")
}