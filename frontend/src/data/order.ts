import { Order } from "../types";

export const dummyOrders: Order[] = [
  {
    id: "ORD-77451",
    date: "1 November 2025",
    status: "processing",
    total: 35.0,
    items: [
      {
        id: "p5",
        name: "Minimalist Graphic T-Shirt",
        imageUrl: "https://placehold.co/100x100/eee/333?text=T-Shirt",
        quantity: 1,
      },
    ],
  },
  {
    id: "ORD-77450",
    date: "31 Oktober 2025",
    status: "processing",
    total: 89.5,
    items: [
      {
        id: "p2",
        name: "Wireless Bluetooth Headphones",
        imageUrl: "https://placehold.co/100x100/007bff/white?text=Headphone",
        quantity: 1,
      },
    ],
  },

  {
    id: "ORD-77449",
    date: "30 Oktober 2025",
    status: "shipping",
    total: 120.0,
    items: [
      {
        id: "p4",
        name: "Premium Running Shoes",
        imageUrl: "https://placehold.co/100x100/e83e8c/white?text=Sepatu",
        quantity: 1,
      },
    ],
    trackingNumber: "JN-1002348B",
    estimatedDelivery: "3 November 2025",
  },

  {
    id: "ORD-77448",
    date: "25 Oktober 2025",
    status: "completed",
    total: 149.99,
    items: [
      {
        id: "p1",
        name: "Classic Leather Jacket",
        imageUrl: "https://placehold.co/100x100/333/white?text=Jaket",
        quantity: 1,
      },
    ],
  },
  {
    id: "ORD-77447",
    date: "22 Oktober 2025",
    status: "completed",
    total: 220.0,
    items: [
      {
        id: "p3",
        name: "Minimalist Wrist Watch",
        imageUrl: "https://placehold.co/100x100/f0f0f0/555?text=Jam",
        quantity: 1,
      },
    ],
  },
];
