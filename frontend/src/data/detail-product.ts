import { DetailedProduct } from "../types";

export const detailedProductData: DetailedProduct = {
  id: "p1",
  name: "Classic Leather Jacket",
  category: "Pria",
  price: 149.99,
  oldPrice: 199.99,
  rating: 4.7,
  imageUrl: "https://placehold.co/600x600/333/white?text=Jaket+Utama",
  images: [
    "https://placehold.co/600x600/333/white?text=Jaket+Utama",
    "https://placehold.co/600x600/555/white?text=Jaket+Detail+1",
    "https://placehold.co/600x600/777/white?text=Jaket+Detail+2",
    "https://placehold.co/600x600/999/white?text=Jaket+Belakang",
  ],
  stock: 24,
  description: `
    <p>Tingkatkan gaya Anda dengan Jaket Kulit Klasik kami. Dibuat dari 100% kulit domba asli, jaket ini menawarkan kenyamanan dan daya tahan yang tak tertandingi.</p>
    <p class="mt-4">Fitur utama:</p>
    <ul class="list-disc list-inside mt-2">
      <li>Bahan kulit domba premium</li>
      <li>Desain ritsleting asimetris</li>
      <li>Beberapa saku fungsional</li>
      <li>Lapisan dalam yang lembut</li>
    </ul>
  `,
  variants: {
    type: "Ukuran",
    options: ["S", "M", "L", "XL"],
  },
  reviews: [
    {
      id: "r1",
      author: "Ahmad Subagja",
      rating: 5,
      comment:
        "Kualitasnya luar biasa! Kulitnya sangat lembut dan pas di badan. Pengiriman juga cepat. Sangat direkomendasikan.",
      date: "10 Oktober 2025",
    },
    {
      id: "r2",
      author: "Siti Aminah",
      rating: 4,
      comment:
        "Jaketnya bagus, tapi ukurannya sedikit lebih kecil dari yang saya kira. Sarankan untuk naik satu ukuran.",
      date: "5 Oktober 2025",
    },
  ],
};
