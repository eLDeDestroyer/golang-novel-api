package seeders

import (
	"e-novel/model"
	"fmt"

	"gorm.io/gorm"
)

func PageSeeder(db *gorm.DB) error {
	data := []model.Page{
		{
			Page: 1,
			Text: `Kereta berhenti perlahan di stasiun kota kecil itu. Awan melangkah turun, menenteng koper kecil berwarna cokelat yang sudah mulai terkelupas di sudut-sudutnya. Angin sore menyambutnya dengan bau tanah dan suara jangkrik yang tak pernah benar-benar ia lupakan.
Langkahnya terhenti sejenak di pintu keluar stasiun. Ia mengangkat kepalanya, memandang langit yang masih biru seperti dulu. Ada sesuatu yang menggantung di sana—bukan awan, bukan juga cahaya. Mungkin kenangan.
Kota ini tak banyak berubah. Toko kelontong di pojok jalan masih berdiri, meski catnya mulai pudar. Anak-anak kecil berlari di tepi jalan sambil tertawa, seperti versi kecil dari dirinya dan Biru dulu.
Ia menelusuri jalanan pelan-pelan, setiap langkahnya seperti membuka kembali halaman lama buku hidupnya. Ia melewati taman kecil tempat dulu mereka biasa duduk sambil berbagi mimpi. Pohon flamboyan di tengah taman itu masih berdiri, kini jauh lebih besar dari ingatannya.
Saat malam mulai turun, Awan tiba di rumah kosong peninggalan orang tuanya. Ia membuka pintu perlahan, mencium bau kayu tua dan debu yang tertahan waktu. Di ruang tamu, jam dinding masih menggantung—berhenti di pukul dua siang, seolah menunggu waktu yang sama datang kembali.
“Selamat datang kembali, Awan,” bisiknya pada dirinya sendiri, pada rumah, pada masa lalu yang belum usai`,
			BookId: 1,
		},
		{
			Page: 2,
			Text: `Langkah kaki Awan membawanya ke bangunan tua dengan papan nama kusam: "Perpustakaan Daerah Nusa." Pintu kayunya berderit saat dibuka, seperti menyapa pengunjung lama yang pulang dari pengembaraan.

Ia berjalan di antara rak-rak buku yang berdebu, menyentuh punggung buku dengan jari-jarinya yang gemetar pelan. Di sudut ruangan, masih ada meja kayu tua tempat ia dan Biru dulu duduk berjam-jam, membaca tanpa suara.

Tiba-tiba, matanya tertumbuk pada sebuah buku puisi karya Sapardi Djoko Damono. Ia menariknya, membuka halaman pertama—dan di sana, pada halaman ketiga puluh dua, ada tulisan tangan yang sangat ia kenal:

"Jika suatu hari kamu kembali, lihatlah langit dari tempat ini. Aku mungkin tidak ada, tapi biru akan selalu menyapa kita."

Tanda tangan di bawahnya hanya satu kata: Biru.

Awan menutup buku itu perlahan. Ia menatap langit dari jendela perpustakaan yang retak. Masih biru. Masih sama. Tapi dadanya terasa berat, seperti ada yang belum selesai.

Ia duduk lama di kursi itu, mencoba merangkai waktu yang sudah pecah. Ia tahu, ini bukan sekadar perjalanan pulang. Ini perjalanan untuk menyusun kembali serpihan yang dulu ia tinggalkan`,
			BookId: 1,
		},
		{
			Page: 3,
			Text: `Warung nasi Bu Rini berdiri di pojok jalan seperti dulu. Aroma nasi goreng dan sambal terasi menyambutnya saat ia melangkah masuk. Ruangan kecil itu penuh dengan kursi plastik biru dan meja kayu yang dilapisi taplak bunga-bunga.

"Mas Awan?" Bu Rini menyapanya, matanya membelalak mengenali wajah yang lama tak ia lihat.

"Iya, Bu. Lama nggak ke sini."

"Sudah sepuluh tahun, ya? Duduk, duduk. Saya buatkan teh hangat dan nasi rawon favoritmu."

Awan tersenyum kecil, duduk di meja yang biasa ia dan Biru tempati. Ia bisa melihat bayangan gadis itu di seberang meja—senyumnya yang cerah, cara ia menyendok nasi perlahan, dan tawanya saat mereka berbicara tentang puisi atau langit.

Ia menyentuh permukaan meja, berharap waktu bisa benar-benar mundur. Tapi tidak. Hanya ingatan yang terus bermain seperti film usang di benaknya.

Saat makanan datang, Awan makan pelan-pelan. Seolah setiap sendok membawa kembali rasa yang dulu, rasa yang tak berubah meski segalanya telah berlalu. Dan di tiap tegukan teh, ia merasa semakin dekat dengan masa lalu yang ingin ia pahami kembali.`,
			BookId: 1,
		},
		{
			Page: 4,
			Text: `Hari itu cerah. Langit membentang biru tanpa cela, seperti kanvas yang belum disentuh kuas. Awan berdiri di taman kota, memandangi langit dari tempat Biru pernah mengajaknya duduk diam.

"Kalau langit ini masih biru, aku masih di sini," begitu kata Biru, sepuluh tahun lalu. Kata-kata itu terus membayang dalam pikirannya, seperti gema yang tak mau pergi.

Awan duduk di bangku yang sama. Ia memejamkan mata, membiarkan cahaya matahari menembus kelopak matanya yang tertutup. Dalam keheningan, ia mencoba mendengar suara Biru. Tapi yang terdengar hanya suara angin dan dedaunan yang gugur pelan.

Ia meraih buku catatan kecil dari sakunya. Di sana, ia menulis satu kalimat: “Langit masih biru, tapi aku tak tahu di mana kamu.”`,
			BookId: 1,
		},
		{
			Page: 5,
			Text: `Di rumah lamanya, Awan menemukan kotak kayu kecil di bawah tempat tidur. Ia membukanya perlahan. Di dalamnya ada tumpukan kertas, foto-foto lama, dan satu amplop berwarna krem dengan namanya tertulis di depan: "Untuk Awan."

Tangannya gemetar saat membuka surat itu. Tertanggal dua minggu setelah ia pergi meninggalkan kota. Tulisan Biru rapi dan halus, seperti suaranya yang lembut.

“Awan, jika kamu membaca ini, berarti kamu akhirnya pulang. Aku tidak tahu apakah kamu masih mengingatku, atau langit biru itu. Tapi aku ingin kamu tahu bahwa aku menunggumu. Bukan untuk kembali, tapi untuk mengerti.”

Surat itu diakhiri dengan satu kalimat yang membuat Awan menutup mata, menahan air yang menggantung di ujung pelupuk:

“Kita tak perlu kembali untuk bersama, cukup mengenang tanpa kehilangan.”`,
			BookId: 1,
		},
		{
			Page: 6,
			Text: `Di toko buku kecil yang baru direnovasi, Awan membeli buku secara acak. Penjaga toko, perempuan seusia Biru, menatapnya lama.

"Kamu Awan? Temannya Biru, ya?"

"Iya... kamu siapa?"

"Aku Rani. Teman kuliah Biru. Dia sering cerita tentang kamu. Tentang langit biru."

Awan tersenyum kaku. Hatinya bergetar mendengar nama itu disebut oleh orang asing. Rani mengajaknya duduk, lalu bercerita. Tentang Biru yang memilih tetap tinggal, tentang mimpinya membangun taman baca untuk anak-anak, dan tentang satu janji yang tak pernah ia lupakan.

"Biru selalu bilang, kalau Awan kembali, dia pasti akan tahu lewat langit."`,
			BookId: 1,
		},
		{
			Page:   7,
			Text:   `Rani mengajak Awan ke sebuah taman kecil di pinggir kota. Ada papan bertuliskan "Taman Baca Biru Langit". Di sana, anak-anak duduk membaca, tertawa, berlari. Di sudut taman, ada bangku kayu dengan ukiran kecil:

“Untuk Awan, yang mengajarkanku melihat langit.”

Awan berdiri lama di situ. Dadanya sesak, tapi hangat. Ia merasa dekat, seolah Biru tak benar-benar pergi. Ia duduk, mengambil buku anak-anak, dan mulai membaca bersama mereka. Dalam tawa dan cerita, ia merasa utuh kembali.`,
			BookId: 1,
		},
		{
			Page:   8,
			Text:   `Hujan turun saat malam. Awan berjalan di bawah payung, menyusuri jalanan sepi yang dulu mereka lewati bersama. Di sebuah gang sempit, ia berhenti. Dindingnya penuh coretan puisi, sebagian dari Biru.

Ia membaca satu per satu. Lalu menambahkan satu baris miliknya:

“Aku pulang, Biru. Dan langit masih sama.”`,
			BookId: 1,
		},
		{
			Page:   9,
			Text:   `Malam itu, Awan bermimpi. Ia melihat Biru duduk di bangku taman, mengenakan gaun putih, tersenyum padanya. Mereka tidak bicara, hanya saling pandang.

Saat Awan hendak meraih tangannya, Biru berkata lirih:

"Jangan berhenti hanya karena aku tak di sini. Lihat langit, dan kamu akan tahu jawabannya."

Ia terbangun dengan mata basah. Tapi kali ini, ia tersenyum.`,
			BookId: 1,
		},
		{
			Page:   10,
			Text:   `Pagi itu, Awan menulis surat. Surat balasan untuk Biru, yang akan ia tinggalkan di bangku taman tempat mereka dulu duduk.

“Aku sudah kembali. Aku sudah tahu. Terima kasih telah menungguku lewat langit.”

Ia menatap ke atas. Langit masih biru. Dan untuk pertama kalinya, ia merasa bisa benar-benar melangkah maju.

Karena meski tak lagi bersama, mereka telah menyatu dalam langit yang sama.`,
			BookId: 1,
		},

		{
			Page: 1,
			Text: `Langit tidak menyukai hujan. Baginya, hujan hanya memperlambat langkah, mengaburkan pandangan, dan mengingatkan pada masa kecil yang sepi. Tapi hari itu berbeda. Hujan turun deras di luar gedung kampus saat ia sedang membaca buku sendirian di kantin. Suasana dingin, gemericik di jendela, aroma kopi yang basi. Ia sendirian, seperti biasa.

Tiba-tiba seorang gadis datang, duduk di depannya tanpa permisi, mengibaskan air hujan dari jaket jeans lusuhnya, dan tersenyum.

"Boleh nebeng duduk? Semua kursi penuh."

Langit menoleh, sedikit terkejut. Ia mengangguk pelan. Tidak tahu bahwa senyum itu akan tinggal jauh lebih lama dari hujan yang saat itu mengguyur kota.

Gadis itu memperkenalkan diri. Dara. Mahasiswi jurusan seni. Satu angkatan di bawah Langit. Pembicaraan mereka ringan, soal kopi yang hambar, soal dosen killer, soal buku yang belum sempat dibaca sampai habis.

Tapi ada sesuatu dalam cara Dara tertawa yang membuat waktu terasa lebih lambat. Sesuatu yang membuat Langit ingin bertanya lebih banyak, tapi memilih diam dan mengamati.

“Nama kamu Langit?”

Ia mengangguk. “Iya.”

Dara tersenyum lebar. “Nama yang bagus. Aku suka langit.”

Langit menatapnya lama. Entah kenapa, hatinya ikut hangat meski tubuhnya masih kedinginan. Hujan berhenti, tapi mereka belum ingin beranjak. Saat akhirnya Dara berdiri, ia berkata,

“Semoga kita ketemu lagi, Langit.”

Itu kalimat sederhana. Tapi cukup untuk membuat Langit menutup bukunya dan tersenyum kecil. Ia tahu, ia akan mengingat hari ini.`,
			BookId: 2,
		},
		{
			Page: 2,
			Text: `Hari-hari setelah itu, Langit kembali ke kantin di jam yang sama. Kadang sambil membaca, kadang hanya duduk tanpa tujuan. Tapi Dara tidak datang.

Lalu, beberapa minggu kemudian, mereka bertemu lagi. Kali ini di perpustakaan. Dara mengenali Langit duluan. Ia datang sambil membawa tumpukan buku tentang seni dan filsafat, lalu duduk di seberang meja.

“Langit, si penyendiri dari kantin!” katanya sambil tersenyum.

Mereka tertawa. Dan sejak saat itu, mereka mulai sering bertemu. Di kelas, di taman kampus, di toko roti favorit Langit. Mereka mulai berbagi makanan, cerita, dan tawa.

Suatu malam, mereka duduk berdampingan di taman kampus, berbagi earphone sambil memutar lagu-lagu yang mereka suka. Langit memutar lagu dari ponselnya.

“Ternyata kamu suka Efek Rumah Kaca juga?”

“Banget,” jawab Dara. “Apalagi lagu yang ini…”

Lagu pun mengalun. Pelan. Tenang. Dan untuk beberapa menit, dunia terasa tidak penting. Hanya ada musik, mereka berdua, dan senyap yang nyaman.

“Rasanya seperti lagi dengerin isi kepala sendiri,” kata Dara sambil menatap langit malam.

Langit mengangguk. “Dan kamu bagian yang bikin suaranya lebih tenang.”

Dara tertawa kecil. Malam itu, di bawah langit berbintang, Langit tahu bahwa hidupnya perlahan sedang berubah. Ada seseorang yang mulai mengambil tempat di ruang paling tenang dalam dirinya.`,
			BookId: 2,
		},
		{
			Page: 3,
			Text: `Minggu sore. Dara mengajak Langit ke ruang seni, tempat ia sering melukis. Ruangan itu seperti dunia kecil yang asing tapi hangat bagi Langit. Dindingnya penuh kanvas, sebagian masih kosong, sebagian penuh warna yang kacau dan indah. Ada aroma cat minyak yang khas.

“Aku nggak melukis untuk jadi bagus,” kata Dara sambil mencelupkan kuas. “Aku melukis supaya aku bisa tetap waras.”

Langit duduk di lantai, memperhatikan Dara bekerja. Ia tidak bicara banyak. Ia hanya menikmati caranya menggores warna, caranya diam tapi hidup.

Dara mulai menggambar sesuatu yang tak bisa dijelaskan dengan kata-kata. Sosok dua orang di tengah hujan, berdiri berhadapan, saling memayungi dengan tangan mereka sendiri.

“Apa itu kita?” tanya Langit akhirnya.

Dara tersenyum. “Mungkin.”`,
			BookId: 2,
		},
		{
			Page: 4,
			Text: `Tiap Jumat malam, mereka punya kebiasaan: naik ke atap gedung lama kampus. Di sana mereka berbagi cerita, melihat langit, dan membicarakan hal-hal kecil yang tak pernah sempat disampaikan di bawah cahaya lampu.

Malam itu langit bersih, hanya ada sedikit awan dan banyak bintang. Dara bersandar di dinding, memeluk lututnya.

“Langit, aku pernah takut banget sama kesepian.”

Langit menatap wajahnya yang hanya disinari lampu jalan jauh di bawah. “Sekarang?”

“Aku masih takut. Tapi sejak ada kamu... takutnya jadi berkurang.”

Langit menggenggam tangannya. Diam. Tapi erat.`,
			BookId: 2,
		},
		{
			Page: 5,
			Text: `Mereka memutuskan untuk membolos sehari penuh dari rutinitas. Tanpa ponsel. Tanpa jam tangan. Tanpa rencana. Mereka naik angkot entah ke mana, turun di pasar lama, makan bakso kaki lima, lalu naik bus ke pantai.

Di tepi pantai, mereka berjalan tanpa sepatu. Air laut menyentuh kaki mereka. Dara menggambar di pasir, Langit menulis puisi.

Saat matahari hampir tenggelam, mereka duduk di bebatuan, diam dan damai.

“Kalau waktu bisa berhenti,” kata Langit, “aku mau di sini, sama kamu.”

Dara tidak menjawab. Ia hanya menyandarkan kepala di bahu Langit.`,
			BookId: 2,
		},
		{
			Page: 6,
			Text: `Langit menulis surat untuk Dara. Bukan karena ia ingin romantis. Tapi karena ada terlalu banyak hal yang sulit diucap.

Ia menulis tentang ketakutan. Tentang rasa kagum. Tentang keinginan untuk tetap bersama tapi juga kecemasan bahwa semuanya bisa berubah kapan saja.

“Aku jatuh cinta. Diam-diam. Dalam-dalam. Tapi aku takut kamu akan pergi kalau aku bilang.”

Surat itu tidak pernah dikirim. Tapi ia simpan di buku harian—dan di dada.

`,
			BookId: 2,
		},
		{
			Page:   7,
			Text:   `Perasaan Langit tumbuh semakin dalam. Ia mulai memerhatikan cara Dara mengeja kata. Cara ia tertawa saat gugup. Cara ia menatap matahari sore seolah sedang berbicara dengan masa lalunya sendiri.

Tapi ia tetap diam. Karena Dara pernah bilang, “Aku takut kalau semuanya berubah hanya karena satu kata.”

Dan Langit mengerti. Maka ia memilih menjaga. Meski dari kejauhan.`,
			BookId: 2,
		},
		{
			Page:   8,
			Text:   `Tiba-tiba, Dara menghilang.

Tidak ada pesan. Tidak ada kabar. Tidak di kelas. Tidak di taman.

Langit menunggu. Satu hari. Dua hari. Tujuh hari. Ia mulai resah. Ia menulis pesan tapi tidak dikirim. Ia mendatangi ruang seni, tapi kosong.

Lalu suatu sore, Dara muncul. Wajahnya letih. Matanya sembab.

“Maaf aku pergi. Aku... butuh waktu.”

Langit tidak bertanya. Ia hanya menatapnya. Dan dalam diam, Dara tahu: Langit tetap ada.`,
			BookId: 2,
		},
		{
			Page:   9,
			Text:   `Dara menghilang lagi. Lebih lama. Kali ini Langit khawatir sungguhan. Ia mencari tahu lewat teman-teman. Lewat dosen. Lewat Raka, sahabat Dara.

Dari Raka, ia tahu: Dara sedang menjalani terapi. Depresi. Tekanan batin dari masa lalu yang belum selesai.

Langit menulis surat lagi. Dan kali ini, ia kirimkan.

“Kalau kamu merasa dunia terlalu berat, kamu boleh istirahat di pundakku. Aku tidak akan menuntut apa-apa. Hanya ingin kamu tahu: kamu nggak sendirian.`,
			BookId: 2,
		},
		{
			Page:   10,
			Text:   `Beberapa minggu kemudian, Dara datang lagi. Pelan. Tidak langsung tertawa. Tidak langsung bercerita. Tapi cukup untuk duduk di sebelah Langit dan berkata,

“Terima kasih karena nggak pergi.”

Langit menatapnya. “Aku nggak bisa pergi. Karena kamu adalah tempat aku pulang.”

Dara menghela napas, lalu tersenyum. “Kalau begitu, kita pulang bareng.”

Dan sore itu, di bawah langit senja yang temaram, mereka tahu satu hal pasti:

Mereka tidak perlu saling memiliki untuk saling tinggal. Karena yang paling penting adalah...

...bersama kamu.`,
			BookId: 2,
		},



		{
			Page: 1,
			Text: `Mesin jahit tua di pojok ruang berderak pelan, nyaris tenggelam oleh riuh suara pasar pagi yang menggema dari luar. Nara duduk bersila di lantai ruang belakang kios kecil ibunya, tangannya cekatan menuntun jarum dan benang di atas kain polos berwarna hijau zaitun. Cahaya matahari menerobos kisi jendela, menyinari wajahnya yang serius dan kening yang mulai berkeringat meski hari masih pagi.

Ia bukan anak desain. Ia tidak pernah sekolah mode. Ia hanya anak dari penjual kain yang sejak kecil bermain di antara gulungan katun dan sutra, mendengar kisah hidup pelanggan dari balik tirai kios. Tapi satu hal yang ia warisi dari mendiang ayahnya adalah semangat: bahwa mimpi besar bisa dimulai dari benang yang paling kecil.

“Kalau kamu sabar, benang yang kusut pun bisa jadi pola yang indah,” kata ayahnya suatu kali, sambil memperbaiki mesin jahit yang sudah tua itu.

Mesin itulah yang kini ia pakai untuk menyelesaikan kebaya pesanan Bu Retno, langganan lama mereka yang akan merayakan ulang tahun pernikahan ke-30. Nara tak dibayar banyak. Tapi tiap jahitan yang ia buat selalu ia niatkan dengan sungguh-sungguh, seolah tiap helai kain adalah pijakan menuju panggung impiannya sendiri.

“Ra, kamu nggak telat kuliah?” suara Ibu terdengar dari depan kios.

Nara menghentikan mesin, melirik jam dinding. “Astaga! Sudah hampir jam delapan!”

Ia buru-buru menggunting sisa benang, menata kain jadi rapi, lalu menggenggam map sketsa yang selalu ia bawa ke mana-mana. Di dalamnya, ada puluhan rancangan gaun yang belum pernah dilirik siapa pun—rancangan dari hati, dari impian yang belum punya nama.

Sebelum pergi, ia mencium tangan ibunya. “Nanti aku bantu lagi sore, ya.”

“Jangan lupa makan siang,” pesan Ibu, sambil menyelipkan uang dua puluh ribu ke tas kecil Nara.

Di jalan menuju halte, Nara berjalan cepat di antara kaki-kaki orang dewasa yang sibuk berdagang. Angin pagi membawa bau rempah dari warung seberang dan suara tawa anak-anak yang berlarian dengan seragam kusut. Ini bukan kota besar. Tapi bagi Nara, setiap sudutnya menyimpan alasan untuk bermimpi lebih jauh.

Di dalam angkot, Nara membuka map-nya. Salah satu sketsa terbaru bergambarkan gaun dengan motif awan dan langit senja. Ia menulis di bawahnya: “Untuk mereka yang berjalan perlahan, tapi tak pernah berhenti.”

Dan ia tahu, meski ia belum punya panggung, belum punya studio, belum punya modal—ia punya benang. Dan benang itu cukup untuk memulai.`,
			BookId: 3,
		},
		{
			Page: 2,
			Text: `Elang kembali datang ke kios dua hari kemudian. Kali ini ia membawa beberapa potong kain yang sudah dijahitnya—hasil latihan pertamanya. Jahitannya masih miring, banyak simpul yang kurang rapi, tapi mata Nara menangkap sesuatu yang lebih penting: ketekunan.

“Mau belajar serius?” tanya Nara sambil menatap jahitan itu.

Elang mengangguk. “Aku ingin mulai dari bawah. Aku sudah lelah kerja di belakang meja, duduk dari pagi sampai malam tanpa tahu apa yang benar-benar aku buat.”

Sejak hari itu, Elang sering muncul. Kadang sekadar duduk di kursi kecil dekat pintu kios, kadang membantu Ibu Nara menata kain, kadang datang membawa hasil belajar dari buku-buku yang ia pinjam dari perpustakaan kota.

Di antara tumpukan kain dan suara mesin jahit, mereka mulai bicara banyak hal. Tentang impian, tentang kegagalan, dan tentang betapa kerasnya dunia saat seseorang mencoba membangun jalan sendiri.`,
			BookId: 3,
		},
		{
			Page: 3,
			Text: `Nara kini punya rutinitas baru. Sepulang dari kampus, ia langsung membantu ibunya. Lalu malam hari, ia menghabiskan waktu bersama Elang di ruang belakang kios. Mereka membuat pola, menjahit ulang, memperbaiki kesalahan.

“Kenapa kamu nggak ikut lomba desain yang diselenggarakan kampus?” tanya Elang suatu malam.

Nara tertawa lirih. “Pendaftaran mahal. Dan aku bukan anak desain, jadi pasti kalah dari mereka yang punya guru dan fasilitas.”

“Tapi kamu punya sesuatu yang mereka nggak punya. Kamu kerja lebih keras.”

Elang membuat Nara mulai percaya pada dirinya sendiri. Sedikit demi sedikit, sketsanya berubah jadi prototipe. Gaun sederhana dengan pola yang ia pelajari sendiri. Kemeja dengan motif tenun yang ia kembangkan dari pola tradisional daerah.`,
			BookId: 3,
		},
		{
			Page: 4,
			Text: `Satu malam, kios mereka kemalingan. Beberapa gulung kain mahal hilang. Ibunya hampir pingsan karena syok. Nara menatap mesin jahit tua mereka yang hampir dibawa kabur—hanya selamat karena ia menguncinya di lemari kayu.

Hari-hari berikutnya berat. Mereka kehilangan banyak pelanggan karena tidak bisa segera mengisi ulang stok. Nara harus meminjam uang ke koperasi pasar demi menyambung hidup. Tapi ia tidak menyerah.

Ia mulai menjahit dari kain sisa, membuat aksesoris sederhana untuk dijual online. Elang membantunya membuka akun media sosial dan memfoto produknya. Satu demi satu, pesanan kecil mulai masuk.`,
			BookId: 3,
		},
		{
			Page: 5,
			Text: `Dua minggu menjelang batas pendaftaran lomba desain di kampus, Elang datang membawa kabar baik.

“Ada beasiswa khusus untuk peserta dari luar jurusan. Aku sudah bantu isi formulirnya. Kamu tinggal tanda tangan.”

Nara terkejut, lalu meneteskan air mata. Ia tahu ini kesempatannya. Ia bekerja siang malam menyelesaikan gaun yang terinspirasi dari kisah ibunya—kain batik yang diwariskan dari generasi ke generasi, dipadukan dengan potongan modern yang berani.

Hari lomba tiba. Nara gemetar saat mempresentasikan karyanya. Tapi saat ia melihat Elang dan ibunya duduk di deretan kursi belakang, ia merasa kuat..`,
			BookId: 3,
		},
		{
			Page: 6,
			Text: `Nara tidak menang. Namanya bahkan tidak disebut dalam nominasi.

Ia menangis malam itu. Bukan karena kalah, tapi karena kecewa pada dirinya sendiri.

Tapi Elang mendekatinya dan berkata pelan, “Kamu mungkin tidak menang hari ini, tapi kamu baru saja membuat dirimu dikenal. Lihat akun media sosialmu. Ratusan orang mulai bertanya tentang desainmu.”

Dan benar. Keesokan harinya, akun kecil mereka dibanjiri pesan. Ada yang tertarik membeli desain, ada yang ingin bekerja sama.

`,
			BookId: 3,
		},
		{
			Page:   7,
			Text:   `Hari-hari berikutnya penuh perjuangan. Pesanan datang, tapi mereka masih terbatas peralatan. Mesin jahit mereka rusak dua kali, dan Nara sempat sakit karena kelelahan.

Namun, ia terus bekerja. Ia bangun jam lima, bantu ibunya, kuliah, menjahit, tidur hanya empat jam sehari. Elang juga mulai menjual kemeja rancangannya secara online.

Satu malam, saat mereka duduk menatap langit dari atap kios, Elang berkata, “Ternyata, merajut mimpi itu butuh jarum yang kuat. Tapi juga tangan yang sabar.”`,
			BookId: 3,
		},
		{
			Page:   8,
			Text:   `Sebuah brand kecil dari Jakarta menghubungi mereka lewat media sosial. Mereka tertarik dengan kombinasi desain tradisional dan modern yang dibuat Nara.

Mereka menawarkan kerja sama: Nara diminta mengirim tiga sampel dan akan dikurasi untuk acara pameran desain muda nasional.

Kembali, Nara berjibaku. Ia pinjam mesin jahit dari tetangga, membeli bahan dari hasil jualan online, dan menyelesaikan tiga gaun dalam empat hari.`,
			BookId: 3,
		},
		{
			Page:   9,
			Text:   `Hari keberangkatan ke Jakarta, Nara berdiri lama di depan kios kecil mereka. Ia mencium tangan ibunya, menggenggam erat map berisi desain.

“Bukan soal menang atau kalah,” kata ibunya. “Tapi soal berani maju.”

Elang ikut mengantarnya ke stasiun. Di sana, Nara menatap kereta yang mulai bergerak, lalu menoleh pada Elang.

“Terima kasih sudah percaya.”

“Selalu,” jawab Elang. “Jahit mimpimu. Aku tunggu hasilnya.”`,
			BookId: 3,
		},
		{
			Page:   10,
			Text:   `Pameran di Jakarta berlangsung meriah. Nara berdiri di antara desainer muda dari seluruh Indonesia. Gaunnya dipajang di pojok ruangan, tapi menarik banyak perhatian karena ceritanya yang mengalir dari tiap jahitan.

Seorang juri menghampirinya. “Desainmu sederhana, tapi jujur. Itu kekuatanmu.”

Nara tersenyum. Ia tidak tahu apakah akan menang. Tapi untuk pertama kalinya, ia merasa pantas berada di sini.

Dan malam itu, saat lampu sorot panggung menyala ke arah gaunnya, Nara tahu: kerja kerasnya tidak sia-sia. Ia sudah memulai. Ia sedang menjahit jalannya sendiri.

Karena mimpi, seperti kain, tak akan berarti tanpa usaha untuk menjahitnya.`,
			BookId: 3,
		},



		{
			Page: 1,
			Text: `Pagi itu, langit di atas Desa Pelangi tampak lebih cerah dari biasanya. Awan-awan putih menggantung pelan seperti kapas, dan angin semilir meniupkan aroma bunga liar dari bukit utara. Di antara padang rumput yang membentang, Anna berlari kecil sambil membawa buku cerita kesayangannya—tentang kerajaan yang terapung di langit dan naga yang menjaga pelangi.

Namun hari itu berbeda. Saat ia duduk di bawah pohon apel tua untuk membaca, seekor kupu-kupu besar dengan sayap berwarna-warni mendarat di halaman bukunya. Sayapnya memantulkan cahaya seperti kaca mozaik.

"Wow... kamu dari mana?" bisik Anna.

Kupu-kupu itu mengepakkan sayapnya dan terbang pelan, seakan mengundang Anna untuk mengikutinya. Tanpa ragu, ia berlari mengikuti kupu-kupu itu, melintasi pagar kayu, menyeberangi ladang, dan akhirnya... masuk ke dalam semak lebat yang belum pernah ia lihat sebelumnya.

Ia terus berjalan, ranting-ranting menyentuh pipinya, tanahnya lembap, dan udara terasa harum. Hingga ia sadar—ini bukan hutan biasa. Cahaya di dalamnya seperti pelangi yang berpendar di udara. Daun-daunnya berwarna merah muda, ungu, biru laut. Ada suara denting lonceng entah dari mana.

“Ini… Hutan Pelangi?” gumam Anna.

Ia tersenyum lebar. Petualangannya baru saja dimulai.`,
			BookId: 4,
		},
		{
			Page: 2,
			Text: `Anna berjalan pelan di antara pepohonan berwarna-warni yang daunnya seperti kristal. Saat ia mendongak, langit di atas hutan ini tak seperti langit biasa. Awan berwarna ungu muda, dan sinar matahari berpendar seperti pelangi.
			Tiba-tiba, suara lembut memecah keheningan. “Kamu bukan dari sini, kan?”
			Anna menoleh dan melihat seekor burung ungu besar bertengger di dahan. Burung itu memakai semacam syal dari benang emas, dan matanya cerdas seperti manusia.
			“Aku Sora. Penjaga langit Hutan Pelangi,” katanya.
			Anna tersenyum kaget. “Kamu bisa bicara?”
			“Tentu saja. Semua makhluk di sini bisa, kalau kamu cukup berani untuk mendengar.”
			Sora lalu terbang melingkari Anna dan menceritakan bahwa warna-warna di hutan perlahan menghilang. “Kabut hitam itu datang dari Utara. Kalau tidak dihentikan, hutan ini akan menjadi abu-abu untuk selamanya.”
			Anna mengepalkan tangan. “Aku mau bantu.”
			Sora menatapnya dengan kagum. “Kalau begitu, ikuti aku. Kita mulai petualanganmu.”`,
			BookId: 4,
		},
		{
			Page: 3,
			Text: `Perjalanan pertama mereka membawa Anna ke sebuah danau berkilau seperti cermin pelangi. Airnya tenang, tapi tidak memantulkan bayangan. Di tengah danau, ada batu mengapung—tempat pertama yang menyimpan petunjuk menuju Inti Pelangi.
			Sora menjelaskan bahwa Anna harus menatap ke dalam air dan menjawab pertanyaan dari dirinya sendiri. “Ini ujian pertama: Kejujuran.”
			Anna melihat ke dalam air. Alih-alih bayangan, ia melihat dirinya menangis sendirian di ruang kelas, saat teman-temannya mengolok-olok mimpinya menjadi penjelajah dunia.
			“Apakah kamu takut?” suara dari air bertanya.
			Anna menarik napas dalam. “Aku takut. Tapi aku tidak mau berhenti.”
			Air danau berkilau lebih terang. Batu di tengah danau pecah menjadi pelangi kecil, lalu membentuk jalan menuju sisi lain hutan.`,
			BookId: 4,
		},
		{
			Page: 4,
			Text: `Di tengah padang berumput merah, mereka bertemu rusa bermata emas. Rusa itu terlihat tenang namun waspada.
			“Kenapa kalian datang ke wilayahku?” tanya sang rusa.
			“Kami mencari Inti Pelangi,” jawab Sora.
			Rusa menunduk. “Aku bisa bantu. Tapi hanya jika kalian bisa menyeberangi Jembatan Bayangan tanpa ketakutan.”
			Jembatan Bayangan terbuat dari cahaya. Di bawahnya jurang gelap tak berdasar. Saat Anna melangkah, bayangan ketakutannya muncul: orang tuanya marah, dirinya gagal, sendirian di hutan.
			Tapi ia terus melangkah. Ia percaya.
			Saat ia mencapai sisi lain, rusa tersenyum dan menyerahkan sebutir batu kecil yang bersinar. “Bagian dari Inti Pelangi. Kumpulkan tujuh.”`,
			BookId: 4,
		},
		{
			Page: 5,
			Text: `Perjalanan membawa mereka ke lembah bunga yang dihuni peri kecil. Tapi peri ini tak berbicara biasa. Ia hanya menyampaikan pesan lewat teka-teki:
			"Tanpa suara, aku bisa bicara. Tanpa tubuh, aku bisa terbang. Aku adalah...?"
			Anna berpikir keras. “Mimpi?”
			Peri tersenyum. “Benar.”
			Ia menyerahkan kelopak bunga perak yang berubah menjadi kristal warna-warni. Batu kedua berhasil didapat.`,
			BookId: 4,
		},
		{
			Page: 6,
			Text: `Tiba-tiba, hujan turun. Tapi bukan air biasa. Ini hujan warna. Saat hujan menyentuh kulit, warnanya berubah mengikuti suasana hati.
			Anna yang sedang ragu berubah biru. Tapi perlahan, saat ia mulai percaya diri, tubuhnya bersinar hijau cerah.
			Hujan ini adalah ujian ketiga: Emosi. Anna belajar bahwa mengenali dan mengendalikan perasaan adalah bagian penting dari perjalanan.
			Setelah hujan reda, mereka menemukan batu ketiga di bawah pelangi ganda.
			`,
			BookId: 4,
		},
		{
			Page:   7,
			Text:   `Akhirnya mereka memasuki wilayah kabut. Semua warna menghilang. Sora menjadi lemah.
			Anna harus berjalan sendiri. Di dalam kabut, ia bertemu bayangannya sendiri yang mengolok-olok: “Kamu cuma anak kecil. Tidak ada yang percaya padamu.”
			Tapi Anna melawan. Ia mengingat semua yang sudah ia lakukan. Ia berteriak, “Aku bukan takut. Aku hanya belum selesai.”
			Kabut terbelah. Batu keempat muncul dari dalam tanah.`,
			BookId: 4,
		},
		{
			Page:   8,
			Text:   `Di tengah hutan paling dalam, berdiri Pohon Tertua. Dahan-dahannya menyentuh langit.
			Pohon ini memberi ujian terakhir: Pengorbanan.
			Anna diminta menyerahkan sesuatu yang paling ia sayangi. Ia membuka tas kecil dan mengeluarkan buku cerita favoritnya.
			“Aku ingin menyelamatkan dunia ini lebih dari membaca tentang dunia lain.”
			Pohon menyerap buku itu dan memberikan tiga batu terakhir—melengkapi Inti Pelangi.`,
			BookId: 4,
		},
		{
			Page:   9,
			Text:   `Dengan semua batu di tangan, Anna dan Sora kembali ke tengah hutan. Mereka menyatukan ketujuh batu di atas altar cahaya.
			Pelangi meledak di langit. Warna kembali. Kabut hilang.
			Hutan bersorak. Peri menari. Rusa berlari bebas. Sora mengepak tinggi ke langit.`,
			BookId: 4,
		},
		{
			Page:   10,
			Text:   `Anna terbangun di bawah pohon apel. Kupu-kupu masih di atas bukunya.
			Ia pulang ke rumah sambil membawa bunga kristal dari hutan.
			Saat ibunya bertanya, “Kamu dari mana?”
			Anna hanya tersenyum. “Aku baru saja menjahit pelangi.”
			Dan meski tak ada yang percaya sepenuhnya, Anna tahu satu hal:
			Petualangan itu nyata. Dan ia tidak akan pernah sama lagi.`,
			BookId: 4,
		},


		{
			Page: 1,
			Text: `Kabut turun cepat sore itu saat Sinta berhenti di depan rumah barunya—sebuah bangunan tua bergaya kolonial yang berdiri sendirian di ujung jalan berbatu. Rumah itu besar, terlalu besar untuk ia tinggali sendiri, tapi sepi dan murah. Agen properti tak banyak bicara saat ia membayar tunai dua bulan lalu. Ia tidak peduli. Ia hanya ingin pergi jauh dari kota, jauh dari ingatan akan kecelakaan itu.

Tangannya menggenggam erat kunci tua yang baru saja ia ambil dari bawah pot bunga depan pintu. Ia sempat mengira tak ada orang yang akan mempercayakan rumah sebesar ini untuk ditinggali sendirian, tapi anehnya, proses pembelian berjalan sangat cepat—terlalu cepat. Tanpa banyak pertanyaan, tanpa banyak syarat.

Sinta membawa dua koper, satu kotak mainan anak, dan sebuah bingkai foto kecil yang selalu ia letakkan di sisi tempat tidur. Di dalamnya, foto putranya, Dion, tersenyum dengan gigi ompong dan tangan kecilnya menggenggam boneka kelinci. Hatinya teriris setiap kali menatap foto itu, tapi ia belum sanggup menyimpannya.

Ruangan dalam rumah dipenuhi debu dan aroma kayu tua yang lembap. Lantai kayunya berderit setiap kali diinjak, dan cat di dinding mengelupas seperti kulit yang menua. Tapi Sinta merasa, setidaknya di tempat ini ia bisa sendiri. Menangis sendiri. Rindu sendiri.

Malam pertama terasa tenang. Terlalu tenang. Sinta duduk di kursi rotan tua di dekat jendela, menatap angin menggerakkan tirai tipis. Ia menyeduh teh hangat dan membuka buku yang tak pernah ia sentuh sejak Dion meninggal. Di luar, suara jangkrik samar-samar terdengar, diselingi deru angin yang sesekali mengetuk jendela.

Saat jarum jam mendekati tengah malam, Sinta hampir tertidur di kursinya. Tapi tiba-tiba...

“Kriiiing... krijiing...”

Bunyi mainan bergemerincing terdengar dari lantai atas.

Sinta membuka matanya lebar. Jantungnya berdegup keras. Ia mengenal suara itu. Itu suara dari bola mainan Dion—bola warna-warni dengan lonceng di dalamnya yang akan berbunyi jika terguling.

Tidak mungkin. Bola itu disimpan di dalam kotak mainan yang belum ia buka.

Ia berdiri perlahan, menahan napas, menaiki anak tangga satu per satu sambil memegang senter dari dapur. Setiap anak tangga berderit pelan, seperti mengeluh karena beban waktu. Lorong di lantai dua tampak lebih gelap dari seharusnya, meski bulan purnama bersinar terang di luar.

Di ujung lorong gelap, pintu kamar yang ia siapkan sebagai kamar Dion terbuka sedikit. Ia yakin telah menutupnya sore tadi. Lampu di langit-langit tidak menyala. Ia hanya mengandalkan cahaya senter kecil yang berkedip-kedip.

Tangannya gemetar saat mendorong pintu itu. Ruangan itu kosong. Dindingnya polos, hanya ada kotak mainan yang belum dibuka di pojok ruangan.

Dan di tengah lantai kayu... bola warna-warni Dion terguling pelan ke arahnya.

“Tidak...” bisik Sinta.

Ia mundur perlahan. Tapi sebelum ia sempat keluar dari ruangan, terdengar suara pelan di belakangnya.

“...Hai, Nak...”

Suara itu bukan suara Dion. Terlalu dalam. Terlalu hampa. Dan terlalu... dekat.`,
			BookId: 5,
		},
		{
			Page: 2,
			Text: `Keesokan paginya, Sinta duduk di dapur dengan secangkir kopi yang sudah dingin. Bola Dion masih berada di atas meja makan, bukti bahwa suara semalam bukan mimpi. Ia memandang ke arah kotak mainan yang belum dibuka sejak pindahan.

Dengan tangan gemetar, ia membuka tutupnya. Di dalamnya tersusun rapi mainan-mainan Dion: mobil-mobilan kecil, buku gambar, boneka kelinci kesayangannya. Tapi yang membuatnya tercekat adalah lembaran kertas lusuh di bawah boneka.

Sebuah catatan dengan tulisan tangan kecil yang sangat ia kenal.

"Aku di sini, Bu. Jangan menangis."

Air matanya jatuh tanpa suara. Ia meremas kertas itu, memeluk boneka kelinci Dion, dan menangis lama. Tapi saat ia mengangkat kepalanya, boneka itu seolah... tersenyum.

Dan dari atas, terdengar suara langkah kaki kecil—berlari. Padahal rumah itu hanya dihuni oleh dirinya seorang.`,
			BookId: 5,
		},
		{
			Page: 3,
			Text: `Matahari belum sepenuhnya muncul ketika Sinta terbangun karena suara derit pelan dari lantai bawah. Ia mengira itu suara rumah yang mengembang karena suhu malam, tapi perasaannya mengatakan lain. Ia turun perlahan, menyusuri lorong menuju ruang tengah, dan berhenti di depan dinding tempat ia menggantung dua lukisan besar—satu foto keluarga terakhir bersama Dion, satu lagi lukisan tua yang sudah menempel sejak rumah itu dibeli.

Tapi kini posisi keduanya berubah.

Foto keluarganya yang sebelumnya menggantung lurus di tengah kini bergeser ke kanan, nyaris menempel pada lukisan tua berbingkai emas yang menggambarkan seorang ibu dan anak dengan wajah yang samar, hampir memudar. Sinta yakin benar, semalam ia memastikan posisi semua bingkai sudah rapi.

Ia mendekat. Lukisan tua itu kini terlihat berbeda. Warna wajah anak kecil di dalam lukisan itu memucat, seolah-olah semakin menyerupai Dion. Ia mengerjap beberapa kali, mencoba menyangkal ilusi matanya sendiri. Tapi saat ia mendekat, ia melihat sesuatu terjepit di antara bingkai dan dinding.

Sebuah kertas kecil, kekuningan, hampir rapuh.

Dengan hati-hati, ia menariknya keluar. Tulisannya buram, seperti bekas tinta lama yang mulai luntur. Namun ia masih bisa membaca kalimat pendek yang tertulis:

"Jangan gantikan aku, Bu."

Sinta berdiri terpaku. Tangan yang memegang kertas itu gemetar. Ia menoleh ke arah lukisan tua itu, dan kini... mata anak kecil dalam lukisan itu seolah menatap langsung ke arahnya.

Bahkan saat ia menunduk dan menjauh, tatapan itu seperti terus mengikutinya. Untuk pertama kalinya, Sinta merasa rumah itu bukan tempat yang aman, melainkan tempat yang sedang mengamatinya.

`,
			BookId: 5,
		},
		{
			Page: 4,
			Text: `Malam itu, langit mendung pekat tanpa bintang. Sinta sedang menutup jendela ketika angin kencang menerpa rumah tua itu, menyebabkan kaca bergetar. Tak lama kemudian, terdengar petir menyambar. Lampu di seluruh rumah mati seketika.

Sinta mengambil lilin dari dapur dan menyalakannya. Nyala kecil itu menari-nari, memantulkan bayangan di dinding. Ia mencoba mengalihkan pikiran dengan membaca, tapi suara ketukan dari arah atap mengganggu konsentrasinya.

Dug. Dug. Dug.

Suara itu semakin keras. Lalu sunyi.

Tiba-tiba, lilin padam tanpa sebab. Ruangan gelap gulita. Sinta berdiri, jantungnya berdegup cepat. Ia mendengar napas berat—tapi bukan napasnya sendiri.

Dari arah lemari... terdengar suara bisikan:

"Main yuk, Bu."

Sinta menyalakan senter. Perlahan-lahan ia arahkan cahayanya ke lemari. Pintu lemari terbuka sedikit. Di dalamnya, boneka kelinci Dion duduk diam, menghadap keluar. Di dada boneka itu tertempel secarik kertas:

"Kamu janji akan main lagi."`,
			BookId: 5,
		},
		{
			Page: 5,
			Text: `Esok paginya, saat membersihkan kamar Dion, lantai di pojok ruangan mengeluarkan suara aneh saat diinjak. Sinta membongkar lantai kayu yang longgar dan menemukan sesuatu yang dibungkus kain.

Ia membuka bungkusan itu—sebuah buku harian.

Ia mengenali tulisan tangan Dion. Awalnya berisi gambar-gambar lucu dan cerita kecil. Tapi di halaman terakhir, coretannya berubah.

"Temanku bilang ibu tidak sayang aku lagi. Tapi aku lihat Ibu menangis. Aku akan tetap di sini. Sampai Ibu percaya."

Tangan Sinta gemetar. Ia menutup buku itu dan berbisik, "Ibu percaya, Nak... Ibu percaya."

Dari cermin, bayangan anak kecil tersenyum—tapi tidak ada siapa-siapa di ruangan itu.`,
			BookId: 5,
		},
		{
			Page: 6,
			Text: `Setiap malam, lorong lantai dua terasa lebih panjang. Tapi malam ini, ia merasa lorong itu berubah. Dindingnya lebih sempit, langit-langitnya lebih rendah. Lampu redup di langit-langit berkedip pelan.

Di ujung lorong, berdiri sosok anak kecil. Tak bergerak. Menghadap ke dinding.

"Dion...?"

Sosok itu berlari masuk ke kamar Dion. Sinta mengejar, tapi kamar kosong.

Namun di cermin, ia melihat sosok itu masih berdiri di belakangnya. Perlahan, ia mengangkat tangannya... menunjuk ke arah perut Sinta.

“Ada yang harus keluar…”
			`,
			BookId: 5,
		},
		{
			Page:   7,
			Text:   `Pagi hari, Sinta turun tangga dan melihat jejak kaki kecil berwarna merah di anak tangga—berdarah. Ia panik, mengikuti jejak itu hingga dapur. Kompor menyala sendiri, air mendidih tanpa sebab.

Di atas meja dapur, ada tulisan embun di permukaan kaca:

“Kami lapar.”

Panci jatuh sendiri. Sinta berlari keluar rumah, terengah-engah. Tapi rumah tetangga terdekat terlalu jauh, dan sinyal ponsel hilang.

Saat ia kembali ke dalam, semua jejak darah menghilang. Tapi di cermin dapur, ada tangan kecil menempel dari dalam.`,
			BookId: 5,
		},
		{
			Page:   8,
			Text:   `Dalam loteng yang hampir rubuh, Sinta menemukan koper tua milik pemilik sebelumnya. Di dalamnya ada surat-surat, dan salah satunya tertulis dengan tergesa:

“Anak kami menghilang. Ia sering bicara dengan seseorang di balik lemari. Ia bilang itu Dion. Kami tak pernah punya anak bernama itu.”

Di sisi surat, terdapat foto lama seorang anak perempuan. Di baliknya tertulis:

“Rina, 1987 – menghilang tanpa jejak.”

Sinta jatuh terduduk. Rina... adalah nama tengahnya sendiri.`,
			BookId: 5,
		},
		{
			Page:   9,
			Text:   `Telepon rumah berbunyi keras. Padahal tidak ada sambungan aktif. Sinta angkat dengan tangan gemetar.

"Halo?"

"Ibu... mau Dion kembali?"

"Siapa ini?!"

"Berikan tubuhmu. Maka kami beri dia kembali."

"Tidak... jangan..."

Telepon mati. Tapi suara itu tetap terdengar dari ruang keluarga. Sinta berlari ke arah sumbernya. Televisi menyala sendiri. Menampilkan video Dion tertawa di taman, lalu perlahan... layar dipenuhi bayangan hitam.`,
			BookId: 5,
		},
		{
			Page:   10,
			Text:   `Malam terakhir. Sinta mengambil boneka Dion dan kembali ke kamar anaknya. Ia menyalakan semua lilin yang ia punya, duduk di tengah lantai dan berdoa.

Angin tiba-tiba berhembus dari jendela yang tertutup rapat. Lilin-lilin padam satu per satu.

Dalam kegelapan, terdengar suara:

"Ibu… dingin. Aku sendiri."

Sosok anak kecil muncul dari bayangan, tubuhnya samar, wajahnya... bukan milik Dion. Suara-suara lain mulai berbisik dari dinding, dari lantai, dari langit-langit.

Sinta menjerit, "Aku cuma ingin Dion kembali! Bukan yang lain!"

Tiba-tiba semua senyap.

Ia terbangun pagi harinya di atas lantai kayu. Sinar matahari menyinari kamar. Semua tampak normal.

Kecuali satu hal: bingkai foto Dion di meja kini kosong.

Dan dari ruang tamu, terdengar suara pelan:

"...Hai, Nak..."`,
			BookId: 5,
		},

		{
			Page: 1,
			Text: `Namanya Aira. Seorang gadis 15 tahun dengan mata tajam yang mampu melihat sesuatu yang tak kasatmata. Sejak kecil, ia tahu ia berbeda. Orang-orang menyebutnya anak indigo—dikaruniai kemampuan lebih untuk melihat dan merasakan dunia lain. Tapi Aira bukan tipe yang bersembunyi ketakutan. Justru ia tertarik. Ingin tahu. Sampai hari itu, saat ia dan keluarganya pindah ke rumah tua warisan nenek di desa terpencil bernama Lereng Wana.

Rumah itu besar dan gelap. Penuh debu dan bayangan. Lorong-lorongnya panjang dan berliku seperti labirin. Di malam pertama, Aira mendengar suara-suara dari lorong belakang rumah. Suara berbisik. Tertawa kecil. Dan satu suara berat berbisik namanya: “Aira...”

Ia bangkit dari tempat tidur, membawa senter, dan menelusuri lorong yang dingin itu. Jantungnya berdetak cepat tapi ia terus berjalan. Di ujung lorong, ia melihat pintu tua yang tertutup rantai. Neneknya dulu pernah berkata, “Jangan pernah buka pintu belakang, apa pun yang terjadi.” Tapi malam itu, rantainya jatuh ke lantai.

Dan di balik celah pintu yang perlahan terbuka, Aira melihat sepasang mata merah menyala. Mata itu tak hanya melihat. Ia merasa... dia sedang dipanggil.`,
			BookId: 6,
		},
		{
			Page: 2,
			Text: `Pagi harinya, Aira bangun dengan perasaan aneh. Saat bercermin, ia melihat garis merah di lehernya—seperti bekas jari. Ibunya bilang mungkin Aira menggaruk saat tidur. Tapi Aira tahu, semalam dia tidak tidur nyenyak. Mimpi buruk membayangi, dan perasaan tertindih membuatnya terbangun berkali-kali.

Saat sarapan, ayahnya menyebut tentang suara hewan di malam hari. Aira tidak menjawab. Ia tahu itu bukan suara hewan. Itu suara... napas. Dalam dan berat. Ia kembali ke lorong tempat ia melihat mata merah itu, tapi pintunya kini tertutup rapat dan rantainya terkunci lagi. Seolah tak pernah disentuh.

Namun lantai di depan pintu itu meninggalkan bekas jejak kaki... kaki kecil dan telanjang, basah seperti baru keluar dari sungai. Dan dindingnya basah. Seperti baru saja dilewati sesuatu yang dingin dan berlendir.`,
			BookId: 6,
		},
		{
			Page: 3,
			Text: `Hari ketiga di rumah itu, Aira mulai melihat sesuatu yang lebih aneh. Cermin di kamarnya kadang menampilkan bayangan yang tidak sinkron dengan geraknya. Ia menoleh ke kanan, bayangan menoleh ke kiri. Ia tersenyum, bayangan tetap diam. Semakin sering terjadi, semakin jelas bahwa cermin itu bukan pantulan biasa.

Suatu malam, sosok anak kecil muncul dalam pantulan cermin itu. Rambutnya panjang, kulitnya pucat, dan matanya kosong. Ia hanya berdiri, menatap Aira dari dunia seberang cermin. Anak itu tak berbicara. Tapi selalu menunjuk ke arah lantai kamarnya.

Setelah beberapa malam, Aira memberanikan diri menggali di bawah papan lantai. Ia menemukan sebuah boneka kayu dan catatan kecil yang hanya bertuliskan: "Dia masih di sini. Jangan buka pintunya."

`,
			BookId: 6,
		},
		{
			Page: 4,
			Text: `Di pekarangan belakang, Aira menemukan sebuah sumur tua yang tertutup semak. Tertarik oleh suara aneh yang terus ia dengar tiap malam, ia memberanikan diri mendekat. Saat merapatkan telinganya, ia mendengar suara samar dari dalamnya—suara anak-anak menangis dan kadang tertawa. Kadang suara itu berubah menjadi teriakan.

Ia melempar batu ke dalam, tapi tak pernah mendengar bunyinya jatuh. Sumur itu terlalu dalam... atau mungkin tidak punya dasar sama sekali. Saat berdiri di sisi sumur, Aira merasa angin hangat dari dalamnya menyentuh wajahnya—angin yang berbau amis.

Malam-malam berikutnya, Aira terus bermimpi tentang sumur itu. Dalam mimpinya, ada tali yang tergantung... dan bayangan yang menariknya turun. Suatu saat, ia tahu ia akan ditarik masuk.`,
			BookId: 6,
		},
		{
			Page: 5,
			Text: `Aira tidak pernah menyukai loteng. Tangga kayunya sempit dan setiap pijakan mengeluarkan bunyi mencicit. Tapi suatu malam, saat listrik mati, ia mendengar langkah kaki di atas langit-langit—berjalan perlahan, lalu berhenti tepat di atas kamarnya.

Ia membawa lilin, menaiki tangga loteng sambil menahan napas. Pintu loteng terbuka sendiri saat ia menyentuh gagangnya. Di dalam, ruangan penuh debu dan sarang laba-laba, tapi ada satu hal yang tak seharusnya ada: boneka kayu yang sama seperti yang ia temukan di bawah lantai.

Boneka itu duduk menghadap tembok, dan di belakangnya tertempel selembar kertas: "Aku hanya ingin bermain. Tapi dia mengurungku."`,
			BookId: 6,
		},
		{
			Page: 6,
			Text: `Setiap pagi, Aira menemukan tanda aneh di dinding kamarnya. Awalnya seperti noda. Tapi lama-lama membentuk pola: tangan. Jejak tangan kecil yang muncul perlahan, dari lantai menuju langit-langit. Saat malam tiba, tangan-tangan itu mengeluarkan cahaya redup.

Suatu malam, Aira melihatnya bergerak. Perlahan, tangan-tangan itu menjulur ke bawah, menelusuri dinding. Salah satunya sempat menyentuh rambutnya.

Ia menjerit. Tapi saat orang tuanya masuk, semua menghilang. Dinding bersih. Tapi di bawah bantalnya, tertinggal satu pesan:

"Kita belum selesai."
			`,
			BookId: 6,
		},
		{
			Page:   7,
			Text:   `Di ruang kerja tua milik nenek, Aira menemukan jurnal tua tersembunyi di balik tumpukan buku. Isinya adalah catatan harian neneknya sendiri—tentang makhluk yang terjebak di balik pintu rantai. Makhluk yang dulu datang lewat mimpi, menyamar sebagai anak kecil, dan memohon dibebaskan.

Nenek menuliskan bahwa ia berhasil menyegelnya dengan mantra kuno dan simbol-simbol di seluruh rumah. Tapi jika salah satu pintu terbuka, semuanya bisa kembali.

Halaman terakhir hanya berisi satu kalimat besar, ditulis dengan tinta merah:

"JANGAN BUKA PINTU BELAKANG."`,
			BookId: 6,
		},
		{
			Page:   8,
			Text:   `Aira mulai bermimpi buruk setiap malam. Dalam mimpinya, ia berlari di lorong rumah yang tak berujung, dikejar oleh makhluk tinggi, kurus, tanpa wajah. Makhluk itu bisa merayap di dinding dan langit-langit, dan mengeluarkan suara seperti tangisan bayi.

Suatu malam, ia melihatnya tidak dalam mimpi. Makhluk itu berdiri di ujung lorong nyata. Tidak bergerak. Hanya mengamati. Tapi setiap kali Aira berpaling, makhluk itu lebih dekat.

Ia mencoba menyegelnya dengan simbol dari jurnal nenek. Tapi makhluk itu menertawakannya. Simbol-simbol itu sudah pudar.`,
			BookId: 6,
		},
		{
			Page:   9,
			Text:   `Dengan pengetahuan dari jurnal, Aira menyusun ulang segel di sekitar rumah. Ia menggunakan kapur, air garam, dan benda-benda milik Dion, saudara jauhnya yang dulu hilang secara misterius. Semua simbol diaktifkan kembali.

Tapi pada malam ketujuh, makhluk itu mengamuk. Rumah bergetar. Jendela terbuka sendiri. Pintu belakang bergetar hebat. Suara-suara anak-anak menangis terdengar dari semua sudut rumah.

Aira berdiri di depan pintu belakang, menahan tubuhnya agar tidak jatuh, dan berteriak mantra terakhir yang diajarkan dalam jurnal. Pintu berhenti bergetar. Semua menjadi gelap.`,
			BookId: 6,
		},
		{
			Page:   10,
			Text:   `Pagi menjelang. Rumah itu sunyi. Tidak ada lagi suara. Tidak ada lagi makhluk. Aira merasa lelah, tapi lega. Ia duduk di depan cermin kamarnya, menatap wajahnya yang pucat dan mata yang sembab.

Tapi cermin kembali mempermainkannya.

Bayangannya tidak meniru gerakannya. Bayangan itu... tersenyum.

Dan dari belakang, suara yang sangat dikenalnya kembali berbisik:

"Aira... ayo main lagi."`,
			BookId: 6,
		},



		{
			Page: 1,
			Text: `Langit malam di stasiun luar angkasa Astra Prime tidak pernah benar-benar gelap. Kilatan dari bintang, satelit, dan kapal pengangkut silih berganti melintasi pandangan. Di sinilah Kapten Elara menerima misi paling berbahaya dalam kariernya: mengeksplorasi Planet Poe. Sebuah planet di tepi sistem Galarion yang selama ratusan tahun dinyatakan "mati". Namun dua minggu lalu, stasiun pemantau menerima sinyal kuno dari sana. Sinyal dalam bahasa yang sudah ribuan tahun tak digunakan.

Elara berdiri di dek utama kapal Orionis VII, matanya menatap peta bintang holografik. Ia tahu misi ini bukan sekadar pencarian, tapi sebuah misteri galaksi yang menuntut jawaban. Kru yang ia bawa bukan orang biasa: Yarin si teknisi alien dari ras Khraltor, Lio si ahli astrobiologi, dan Syra, navigator yang bisa membaca medan magnet planet seperti membaca peta daratan.

Sesaat sebelum keberangkatan, layar komunikasi bergetar. Sebuah gambar muncul—potongan visual dari permukaan Poe. Sebuah kota, reruntuhan, dan... bayangan besar bergerak di antara bangunan yang rubuh. "Apa itu...?" tanya Syra pelan. Elara mengatur helmnya dan menjawab, "Itulah yang akan kita cari tahu."`,
			BookId: 7,
		},
		{
			Page: 2,
			Text: `Pendaratan di Poe tidak berjalan mulus. Atmosfernya tidak stabil, menyebabkan turbulensi hebat. Awan ungu pekat menutupi langit, dan petir menyambar secara acak. Kapal mereka berhasil mendarat di dataran tinggi yang menghadap ke kota tua yang terlihat dalam gambar.

Begitu keluar dari kapal, udara Poe terasa berat meskipun layak hirup. Tanahnya berdenyut perlahan, seolah planet ini memiliki denyut nadi sendiri. Suasana sunyi. Terlalu sunyi. Tidak ada hewan, tidak ada angin. Hanya reruntuhan dan patung-patung batu raksasa yang mengawasi mereka dalam diam.

Saat mereka mulai menelusuri jalanan kota kuno, Lio menunjuk ke satu mural di dinding. Lukisan yang menggambarkan makhluk bersayap keluar dari dasar planet, membawa api dan kehancuran. Di bawahnya tertulis kalimat dalam bahasa Poe: "Jangan bangunkan dewa yang tidur."

Elara mulai menyadari bahwa ada alasan mengapa planet ini ditinggalkan. Rasanya seperti berjalan di atas tanah suci yang dikutuk oleh waktu dan kenangan. Satu langkah salah, dan seluruh sejarah kelam Poe bisa bangkit kembali.`,
			BookId: 7,
		},
		{
			Page: 3,
			Text: `Di pusat kota tua, mereka menemukan struktur setengah runtuh yang disebut sebagai Makam Cahaya oleh para arkeolog dahulu. Saat memasuki ruang dalamnya, langit-langit gua berpendar biru, seolah dipenuhi fosfor alami. Yarin menyentuh salah satu dinding, dan hologram kuno menyala, memperlihatkan gambaran tentang makhluk bersayap emas yang terbang meninggalkan planet.

Syra menemukan jalan rahasia yang menuju ke ruangan bawah tanah. Di sana terdapat kristal besar, yang tampak seperti sumber energi. Tapi saat didekati, kristal itu berdenyut… dan mengeluarkan suara berbisik. Suara yang hanya Elara bisa dengar: "Pewaris cahaya telah tiba."

Seluruh ruangan tiba-tiba bergetar. Lantai di bawah kaki mereka membentuk pola melingkar yang bergerak. Lio panik dan mencoba merekam data energi dari kristal, tetapi alatnya mati mendadak. Sumber kekuatan ini tidak ingin disentuh.

Elara merasa ada sesuatu dalam dirinya berubah. Seperti resonansi antara jiwanya dan kristal itu. Ia belum menyadari, tapi kristal mengenalinya. Dan dalam waktu dekat, mereka akan terikat lebih dalam dari yang ia kira.	

`,
			BookId: 7,
		},
		{
			Page: 4,
			Text: `Reruntuhan kota Poe bukan sekadar tumpukan batu dan logam tua. Saat malam menjelang, perubahan aneh terjadi. Dinding-dinding mulai memantulkan cahaya lembut berwarna biru dan ungu, seperti napas yang keluar dari paru-paru planet yang tertidur. Pola-pola kuno muncul di lantai, mengalir seperti sungai energi yang hidup.

Elara memimpin ekspedisi kecil bersama Syra dan Lio menyusuri jalan utama kota menuju pusat reruntuhan. Aroma logam dan ozon memenuhi udara. Tiba-tiba, suara langkah mereka mulai bergema… bukan sekali, tapi dua kali lipat. Seolah ada yang meniru jejak mereka dari kejauhan. Mereka berhenti sejenak, saling menatap dengan cemas, tapi tidak menemukan apa pun—hanya bayangan mereka sendiri yang tampak memanjang dan bergerak aneh di dinding.

Di persimpangan antara dua bangunan roboh, mereka menemukan sebuah ukiran yang tidak ada sebelumnya. Ukiran itu menggambarkan empat penjaga raksasa yang menjaga gerbang cahaya. Tapi yang membuat Lio panik adalah bahwa salah satu penjaga di ukiran itu—berbentuk makhluk kristal transparan—sama persis seperti yang mereka temui sebelumnya. “Apakah makhluk itu yang membuat ukiran ini… semalam?” gumamnya pelan.

Tiba-tiba, tanah di bawah mereka bergetar ringan, dan suara gemuruh pelan terdengar dari bawah permukaan. Bangunan di sekitar mereka memancarkan getaran halus, dan reruntuhan perlahan bergerak, membuka jalur baru ke bawah tanah. Elara mengaktifkan pelindung energinya, menyadari bahwa kota ini merespons kehadiran mereka. Poe, rupanya, memiliki sistem pertahanan otomatis yang belum sepenuhnya mati.

Mereka mundur ke kapal untuk berdiskusi, tapi Elara tidak bisa mengabaikan firasat bahwa kota ini ingin menunjukkan sesuatu. Mungkin bukan untuk menyakiti mereka… tapi memperingatkan. Dan dalam hati kecilnya, ia tahu: reruntuhan ini masih bernyawa—dan ia sedang mencoba berbicara.

`,
			BookId: 7,
		},
		{
			Page: 5,
			Text: `Setelah kejadian di reruntuhan, kru Orionis VII kembali ke kapal untuk mengevaluasi penemuan mereka. Elara tidak bisa berhenti memikirkan apa yang dilihatnya dalam kilatan ingatan saat menyentuh makhluk kristal. Sebuah perang purba, kehancuran, dan makhluk bersayap kegelapan yang dikurung oleh leluhur Poe di pusat planet. Apakah mereka benar-benar terbangunkan?

Yarin mulai menyelidiki pola energi dari kristal dan menemukan bahwa planet Poe menunjukkan aktivitas anomali—denyut energi dari pusat planet naik drastis setiap dua jam. Tidak hanya itu, medan magnet Poe berubah arah, seolah-olah ada kekuatan besar yang sedang bangkit dari dalam.

Lio menyarankan penggalian ke bawah reruntuhan untuk menemukan sumber utama energi tersebut. Mereka menemukan sebuah lorong spiral batu yang mengarah ke bawah tanah. Saat mereka menyusuri lorong itu, mereka mulai mendengar suara—ratapan, nyanyian dalam bahasa kuno, dan suara dentuman logam yang tidak berasal dari alat mereka.

Di ujung lorong, mereka menemukan gerbang besar yang tertutup oleh pelat emas, dihiasi ukiran makhluk bersayap dan matahari hitam. Di tengah-tengah gerbang, terukir peringatan dalam bahasa Poe: “Siapa yang membangunkannya, menanggung kehancurannya.”

Elara menyadari sesuatu yang mengerikan: kristal yang mereka sentuh mungkin bukan hanya penjara. Ia adalah kunci. Dan mereka telah membukanya.`,
			BookId: 7,
		},
		{
			Page: 6,
			Text: `Keesokan harinya, mereka dikejutkan oleh gempa hebat. Tanah Poe bergetar, menciptakan celah besar di dataran tempat mereka mendarat. Kapal hampir terguling. Dari retakan bumi, semburan cahaya ungu dan partikel kristal keluar seperti geyser.

Yarin dan Syra segera menuju sumber retakan untuk menstabilkan kapal. Tapi Syra terpeleset dan jatuh ke dalam celah. Elara dan Yarin berhasil menariknya kembali, namun Syra berubah. Tatapannya kosong dan kulitnya mulai menyerap kilau kristal. Lio mendiagnosisnya terinfeksi oleh semacam organisme energi—mungkin bagian dari makhluk yang tertidur.

Dalam kekacauan, makhluk kristal pertama yang mereka temui muncul kembali. Kali ini ia membentuk tubuh humanoid yang lebih jelas dan memperingatkan Elara: “Dia telah membuka matanya.” Elara berteriak, “Siapa DIA?” Tapi makhluk itu menghilang menjadi debu bercahaya.

Kru kini dihadapkan pada dua pilihan: melarikan diri atau mencari cara menutup kembali celah yang telah terbuka. Tapi Elara tahu satu hal—melarikan diri sekarang berarti membiarkan Poe bangkit dan menyeret galaksi ke dalam kegelapan.
			`,
			BookId: 7,
		},
		{
			Page:   7,
			Text:   `Saat mencoba menghubungi markas di Astra Prime, mereka menerima sinyal balasan dalam bentuk aneh—gelombang simbol cahaya. Lio berhasil mendekripsinya menjadi pesan visual dari ras kuno Poe. Dalam pesan holografik, seorang tetua Poe menjelaskan bahwa makhluk bernama Vhalagor disegel di inti Poe karena kekuatannya dapat memusnahkan sistem bintang hanya dengan kehendak.

Tetua itu mengatakan hanya satu garis keturunan yang bisa mengendalikan kembali segel: keturunan pembawa cahaya. Dan kristal mengenali Elara sebagai salah satunya. Ia tidak hanya pemimpin misi. Ia adalah kunci terakhir pertahanan galaksi.

Elara terpukul. Bagaimana mungkin ia memiliki hubungan darah dengan bangsa kuno ini? Tapi bukti dalam tubuh Syra, dalam pantulan matanya, dan dalam mimpi-mimpi Elara yang mulai muncul sejak pendaratan—semua mengarah ke satu kenyataan. Darah Poe mengalir dalam dirinya.

Dengan kekuatan dan kutukan itu, Elara memutuskan untuk menyusuri pusat planet dan mencari altar penyegelan. Ia tahu waktu tidak berpihak pada mereka. Vhalagor semakin dekat untuk bangkit sepenuhnya.

`,
			BookId: 7,
		},
		{
			Page:   8,
			Text:   `Perjalanan ke pusat planet memakan waktu lebih dari 20 jam melalui koridor alami dan medan lava yang hampir tak bisa dilewati. Di kedalaman, mereka menemukan ruangan besar berbentuk kubah, dengan altar mengambang di tengah ruangan, dikelilingi cahaya yang terus berubah warna.

Altar itu dijaga oleh empat makhluk transparan seperti penjaga gerbang dimensi. Makhluk-makhluk ini menguji Elara—bukan secara fisik, tetapi secara mental dan spiritual. Mereka menggali ingatan Elara, ketakutannya, keraguannya, dan niat sebenarnya.

Dalam salah satu ujian, Elara dihadapkan pada bayangan ibunya yang telah lama meninggal. “Apakah kamu akan mengorbankan dirimu untuk dunia yang tak kamu kenal?” tanya sang bayangan. Elara menjawab, “Jika dunia itu akan menghancurkan semua yang aku cintai, maka ya.”

Penjaga pun membuka jalan. Elara melangkah ke altar, dan cahaya dari kristal utama menyerap ke dalam tubuhnya. Kini ia tidak hanya memimpin. Ia adalah bagian dari Poe. Dan ia adalah satu-satunya harapan menyegel kembali Vhalagor.`,
			BookId: 7,
		},
		{
			Page:   9,
			Text:   `Meski Elara sudah terhubung dengan energi planet, proses penyegelan tidak berjalan cepat. Sementara itu, dari celah yang terbuka, Vhalagor bangkit. Makhluk raksasa yang tak berbentuk, dengan suara yang bergema di seluruh langit Poe. Cakar-cakar bayangannya merusak atmosfer, dan planet mulai terpecah dari inti.

Orionis VII mencoba mengaktifkan sistem pertahanan orbit, tapi Vhalagor menembus medan perlindungan. Satelit hancur, dan langit berubah merah. Suara Vhalagor menggemakan kesedihan, kemarahan, dan rasa lapar akan kehancuran.

Di altar, Elara berjuang melawan pengaruh Vhalagor yang mencoba mengambil alih tubuhnya. Kristal di jantung altar pecah menjadi serpihan cahaya, menyatu dengan darah Elara, memberi kekuatan yang tak manusiawi. Ia mulai mengucapkan mantra terakhir dari hologram Poe—mantra penyegelan yang hanya bisa diaktifkan oleh pewaris cahaya.

Yarin dan Lio bertarung untuk menahan Vhalagor secara fisik, menggunakan teknologi dan kekuatan medan frekuensi tinggi. Tapi Vhalagor terus tumbuh. Satu-satunya harapan tinggal pada Elara.`,
			BookId: 7,
		},
		{
			Page:   10,
			Text:   `Elara melayang di atas altar, tubuhnya menyatu dengan cahaya planet. Ia mengangkat kedua tangannya, dan cahaya dari semua reruntuhan, dari semua kristal, dan dari makhluk-makhluk penjaga mengalir ke arahnya. Suara mantra memenuhi udara, dan langit Poe pecah menjadi ribuan cahaya.

Vhalagor meraung. Tubuhnya pecah-pecah, mencoba membentuk ulang, tapi cahaya itu lebih kuat. Dalam momen terakhir, Elara melihat wajah Vhalagor—wajah penuh kesepian dan kesalahan. Ia bukan hanya monster. Ia makhluk yang dibuang dan ingin kembali.

Dengan air mata mengalir, Elara mengucapkan kata terakhir. Ledakan cahaya menghabisi Vhalagor dan menutup kembali inti planet. Poe tenang. Dan Elara jatuh tak sadarkan diri.

Saat ia terbangun di kapal, Syra memegang tangannya dan berkata, “Kau berhasil.” Tapi Elara tahu, harga dari keberhasilan itu adalah ikatan abadi dengan planet itu. Ia bukan lagi hanya manusia. Ia adalah penjaga baru Poe.

Dan di langit malam, sebuah bintang baru bersinar—cahaya dari planet yang dulu terlupakan, kini hidup kembali.`,
			BookId: 7,
		},



		{
			Page: 1,
			Text: `Pagi itu Gani terbangun dengan keringat dingin. Bukan karena mimpi buruk, tapi karena notifikasi dari grup Line "Pejuang Skripsi 2025" yang isinya sudah 157 chat. Sebagian besar update tentang sidang, ACC bab lima, dan template revisi. Sementara Gani masih bertanya-tanya: "Judul gue apaan, ya?"

Ia duduk di meja belajarnya, memandangi tumpukan kertas draft yang isinya coretan dari empat tahun lalu. Ada bekas kopi, bekas mie instan, dan highlight warna-warni yang kini tidak relevan lagi. "Hari ini gue harus ke kampus. HARUS!" katanya sambil memakai hoodie yang sama sejak tiga hari lalu.

Sesampainya di kampus, ia bertemu dengan Pak Bagus, dosen pembimbingnya yang ketiga. Pak Bagus menatap Gani dari atas kacamata bulatnya dan berkata, "Mas Gani, saya senang kamu masih hidup." Kalimat itu bukan sarkasme—lebih kepada keheranan tulus karena Gani terakhir konsultasi setahun lalu.

Gani tersenyum kaku. Ia duduk, membuka laptop (yang bunyinya seperti pesawat mau lepas landas), dan berkata, "Pak, saya mau lanjut skripsi..." Tapi dalam hati, dia justru bertanya, "Gimana caranya bilang saya lupa password email buat login OJS?"`,
			BookId: 8,
		},
		{
			Page: 2,
			Text: `Tiga hari kemudian, Gani resmi mengajukan judul baru. Lagi. Judul ke-lima. Dan untuk yang kelima kalinya pula, dia bingung mau nulis apa setelah kata “analisis”. Judul terakhirnya adalah “Analisis Semiologi dalam Iklan Mie Instan: Studi pada Kalimat ‘Kenyangnya Melekat’”. Pak Bagus hanya mengangkat alis.

“Kalau kamu ganti judul terus, nanti malah jadi koleksi, bukan skripsi,” ujar Pak Bagus sambil menyeruput kopinya. Gani nyengir, mencoba melucu: “Minimal bisa jadi buku, Pak. ‘101 Judul Skripsi yang Tidak Jadi.’” Tidak lucu. Bahkan satpam di depan fakultas pun kelihatan lebih stres.

Gani akhirnya menyusun ulang judul dengan bantuan ChatGPT dan tiga teman seperjuangannya yang semuanya sudah wisuda. Judulnya sekarang: “Pengaruh Konten TikTok terhadap Perilaku Konsumtif Mahasiswa Generasi Z.” Pak Bagus membaca pelan-pelan lalu mengangguk, “Judul ini... bisa. Asal kamu jangan ngilang lagi.”

Pulang ke kosan dengan perasaan campur aduk, Gani merasa seperti baru saja memenangkan medali perunggu dalam olimpiade skripsi. Belum menang, tapi setidaknya nggak gugur di babak awal. Ia langsung merayakan dengan gorengan dan teh manis hangat, sambil berkata pada dirinya sendiri, “Besok gue nulis Bab 1!”

Tapi seperti biasa, “besok” adalah kata yang bisa ditunda selamanya.`,
			BookId: 8,
		},
		{
			Page: 3,
			Text: `Hari Minggu, Gani akhirnya duduk di depan laptop. File Word berjudul “Bab1_fix_beneran_fix.docx” terbuka. Tapi layar itu kosong selama satu jam. Ia mengetik satu kalimat: “Bab ini membahas mengenai latar belakang.” Lalu berhenti.

Ia membuka YouTube sebentar untuk cari inspirasi. Tiga jam kemudian, ia sudah tahu cara masak tteokbokki, sejarah perang dunia ke-2, dan kenapa alien mungkin menyukai kentang goreng. Tapi Bab 1? Masih tetap satu kalimat.

Ia putuskan buat nulis di warung kopi. Katanya suasana bisa bantu fokus. Tapi baru lima menit di sana, dia sudah ketemu temannya, Arul, yang sekarang jualan kopi di situ. “Bro, lo belum lulus juga? Wah, gue kira lo udah kerja di agensi!” kata Arul sambil menyodorkan cappuccino gratis.

Malu, Gani berusaha fokus. Akhirnya, ia nulis dua paragraf. Tapi ketika listrik di warung padam, laptopnya mati dan ia lupa nge-save. Ia menatap layar hitam sambil berkata pelan, “Kenapa Tuhan mengujiku lewat file Word?”	

`,
			BookId: 8,
		},
		{
			Page: 4,
			Text: `Setelah seminggu bergulat dengan Bab 1, akhirnya Gani berhasil menyelesaikannya. Kali ini dia ingin mencetak dan menunjukkannya langsung ke Pak Bagus. Tapi printer kosan? Ah, printer legendaris itu lebih tua dari niat Gani lulus kuliah.

Gani menghidupkan printer dengan doa dan sedikit tamparan sayang. Tiga lembar pertama keluar mulus. Tapi lembar keempat... macet. Kertas terlipat, tinta muncrat seperti darah di film horor. Ia mencoba mencabut kertas, hasilnya printer error permanen.

“Tenang, tenang, masih bisa ke warnet,” pikirnya. Tapi di warnet, ia lupa bawa flashdisk. File-nya ada di email, dan seperti biasa... password email-nya lupa.

Akhirnya dia login dengan verifikasi lewat nomor lama. Gagal. Opsi terakhir: minta kirim kode ke email cadangan, yang... juga lupa passwordnya. Dalam keputusasaan, Gani menatap layar dan berteriak lirih, “Kenapa Gmail kayak mantan? Nggak bisa dilupain, tapi juga nggak bisa dibuka.”

`,
			BookId: 8,
		},
		{
			Page: 5,
			Text: `Pak Bagus mulai curiga. Gani memang sudah menyerahkan Bab 1, tapi kenapa setelah itu hilang lagi? Tiga minggu tak ada kabar, tak ada chat, tak ada email. Pak Bagus pun iseng membuat instastory: “Ada yang lihat Gani? Terakhir katanya nulis Bab 2.” Mahasiswa lain ikut komentar, “Wah Pak, itu udah jadi urban legend.”

Sebenarnya Gani tidak hilang. Ia hanya terjebak dalam siklus nonton ulang serial Korea dan meyakinkan dirinya bahwa memahami psikologi karakter itu bagian dari riset. Tapi yang dia pahami justru bahwa cinta segitiga lebih mudah dipahami daripada teori komunikasi massa.

Sampai akhirnya, dia mimpi buruk: Pak Bagus datang dalam bentuk Godzilla dan mengamuk di kampus karena Bab 2 belum dikumpulkan. Gani terbangun, basah kuyup oleh keringat, dan langsung mengetik Bab 2 tanpa napas. Hasilnya? Empat halaman penuh referensi... yang semua berasal dari artikel blog tidak jelas.

Dia sadar, kalau ini dibaca Pak Bagus, bisa langsung kena mental. Maka dia konsultasi dulu ke Rani, teman sekelas yang sudah sidang. Rani membaca cepat dan hanya berkata, “Gani, kamu mau skripsi... atau cerpen?”`,
			BookId: 8,
		},
		{
			Page: 6,
			Text: `Bab 2 dikumpulkan. Dan seperti yang sudah bisa ditebak, revisinya lebih panjang dari Bab 2 itu sendiri. Komentar dari Pak Bagus berwarna merah mencolok, memenuhi margin kiri sampai kanan. “Tolong jelaskan ini.” “Sumbernya mana?” “Ini opini, bukan akademik.” “Kenapa ada kata ‘kayaknya’ di kalimat ilmiah?”

Gani menghela napas sambil memutar lagu sedih di Spotify. Ia merasa seperti karakter utama drama tragedi. Tapi, ia tidak menyerah. Ia duduk di depan laptop, membuka kembali referensi (yang sekarang benar), dan mulai memperbaiki paragraf demi paragraf sambil menyeruput kopi sachet yang sudah dingin.

Malam makin larut, mata makin berat, tapi semangatnya justru naik. Ia merasa seperti pahlawan yang akhirnya menyadari tujuan hidupnya. Setiap koreksi dari Pak Bagus menjadi peluru semangat. “Gue harus buktiin gue bisa,” katanya sambil menghapus kata “mungkin” dari semua kalimatnya.

Keesokan paginya, Bab 2 versi revisi dikirimkan. Kali ini, balasan Pak Bagus hanya satu baris: “Sudah jauh lebih baik. Lanjutkan ke Bab 3.” Gani hampir menangis. Kalimat itu seperti surat cinta dari langit. Ia mencetaknya dan menempelkannya di dinding kosan.
			`,
			BookId: 8,
		},
		{
			Page:   7,
			Text:   `Bab 3 berarti pengumpulan data. Gani memutuskan untuk membuat kuesioner daring. Ia membuat 20 pertanyaan dengan desain warna-warni dan mengirimkannya ke semua grup WhatsApp, Line, Telegram, bahkan alumni. Tapi tak ada yang respon. Seakan dunia ikut skripsi juga—dan sedang hiatus.

Ia ganti strategi: menyamar jadi admin giveaway. “Isi survei ini dan menangkan pulsa 50 ribu!” Dalam sehari, 300 respon masuk. Sayangnya, 250 respon diisi oleh akun palsu bernama ‘Budi Skripsi’, ‘Skripsi Mulu’, dan ‘SkripsiSucks123’.

Panitia penelitian palsu ini membuat Gani frustasi. Ia putuskan untuk turun langsung ke kampus, membawa kertas kuesioner, dan memburu mahasiswa yang lewat. Ia dikejar satpam karena dikira jualan MLM.

Akhirnya ia dapat 50 data valid. Ia pulang dengan wajah berseri, seperti habis menang olimpiade. Tapi saat mau input data ke Excel... laptopnya hang. File rusak. Semua hilang. “Ya Tuhan,” bisik Gani sambil menatap langit-langit, “Aku nggak minta kaya, cuma minta fileku balik.”
`,
			BookId: 8,
		},
		{
			Page:   8,
			Text:   `Di tengah keputusasaan, Gani bertemu lagi dengan Arul. Ternyata Arul diam-diam juga sedang skripsi. Mereka berdua pun sepakat membuat grup WhatsApp bernama “Tumbal Skripsi”. Isinya hanya mereka berdua. Tujuannya: saling menyemangati dan... curhat.

Setiap malam mereka video call, saling baca Bab masing-masing, saling kritik, dan saling ketawa. Tapi makin lama, Gani merasa Arul makin cepat progress-nya. “Bro, Bab 4 gue udah kelar. Tinggal daftar sidang.” Gani menelan ludah.

Rasa iri mulai muncul. Gani mulai banding-bandingin diri. Kenapa dia stuck, sedangkan Arul bisa begitu cepat? Tapi dalam satu percakapan, Arul bilang, “Gue cepet bukan karena pinter, Gan. Tapi karena gue takut dikeluarin sama orang tua.”

Gani tertawa. “Sama, bro. Gue juga takut... tapi kayaknya udah pasrah.” Mereka berdua pun tertawa bersama. Ternyata bukan kompetisi. Mereka saling dorong.`,
			BookId: 8,
		},
		{
			Page:   9,
			Text:   `Akhirnya, hari yang ditunggu tiba. Gani berhasil menyelesaikan semua bab. ACC didapat. Ia mendaftar sidang. Dosen penguji sudah ditentukan. Jadwal keluar: Selasa, jam 10 pagi. Gani tidur cepat, baju rapi sudah disiapkan, power point sudah dicetak.

Tapi pagi itu... dia kesiangan. Alarm tidak bunyi. Listrik kosan padam. Gani bangun jam 09.45. Dengan kecepatan sonic, ia sikat gigi, ganti baju (masih pakai sandal jepit), dan naik ojek ke kampus. Dalam perjalanan, hujan turun.

Baju basah, rambut acak-acakan, dan sepatu pinjaman hilang satu di selokan. Gani masuk ke ruang sidang seperti peserta bentrok demo. Tapi ajaibnya, sidang belum dimulai karena dosen penguji telat.

Saat presentasi, Gani grogi setengah mati. Tapi setelah lima menit, ia mulai lancar. Ia menjelaskan latar belakang, data, sampai kesimpulan. Suaranya mantap. Matanya menyala. Dosen penguji mengangguk.`,
			BookId: 8,
		},
		{
			Page:   10,
			Text:   `Setelah 30 menit presentasi dan 20 menit tanya-jawab, Gani diminta keluar ruangan. Lima menit menunggu terasa seperti lima tahun. Ia menggigil, bukan karena AC, tapi karena takut ditanya, “Kenapa referensinya dari Wikipedia?”

Pintu dibuka. Pak Bagus memanggil. “Selamat, Mas Gani. Anda dinyatakan... LULUS.” Dunia hening sesaat. Gani tidak percaya. Ia menatap langit-langit dan hampir menangis. “Beneran, Pak?”

“Iya. Tapi revisi ya,” kata Pak Bagus. Gani tertawa lemas. Tapi itu tidak masalah. Yang penting: LULUS.

Gani berjalan keluar ruang sidang dengan senyum lebar, disambut oleh Arul yang sudah bawa kopi. Mereka saling tos. “Bro, kita akhirnya bebas.”

Dan di sore itu, Gani duduk di taman kampus, membuka laptop, dan mengubah nama file: “Skripsi_Gani_FINAL_beneran_FINAL_SIDANG.docx.” Untuk pertama kalinya... tanpa tanda bintang, tanpa versi.`,
			BookId: 8,
		},
	}

	for _, row := range data {
		err := db.Table("pages").Create(&row).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("succes seed page")

	return nil
}
