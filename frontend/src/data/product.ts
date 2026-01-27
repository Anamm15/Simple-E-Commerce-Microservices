import { Product } from "../types";

export const dummyProducts: Product[] = [
  {
    id: "p1",
    name: "Classic Leather Jacket",
    category: "Pria",
    price: 149.99,
    oldPrice: 199.99,
    rating: 4.7,
    imageUrl: "https://placehold.co/400x400/333/white?text=Jaket",
  },
  {
    id: "p2",
    name: "Wireless Bluetooth Headphones",
    category: "Elektronik",
    price: 89.5,
    rating: 4.5,
    imageUrl: "https://placehold.co/400x400/007bff/white?text=Headphone",
    isNew: true,
  },
  {
    id: "p3",
    name: "Minimalist Wrist Watch",
    category: "Aksesoris",
    price: 220.0,
    rating: 4.8,
    imageUrl: "https://placehold.co/400x400/f0f0f0/555?text=Jam",
  },
  {
    id: "p4",
    name: "Premium Running Shoes",
    category: "Wanita",
    price: 120.0,
    oldPrice: 150.0,
    rating: 4.2,
    imageUrl: "https://placehold.co/400x400/e83e8c/white?text=Sepatu",
  },
];
