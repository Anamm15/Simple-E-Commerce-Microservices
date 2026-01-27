import { CartItem, ShippingOption, PaymentOption } from "../types";

export const cartData: CartItem[] = [
  {
    id: "p1",
    name: "Classic Leather Jacket",
    price: 149.99,
    quantity: 1,
    imageUrl: "https://placehold.co/100x100/333/white?text=Jaket",
  },
  {
    id: "p2",
    name: "Wireless Bluetooth Headphones",
    price: 89.5,
    quantity: 1,
    imageUrl: "https://placehold.co/100x100/007bff/white?text=Headphone",
  },
];

export const shippingOptionsData: ShippingOption[] = [
  { id: "ship-1", name: "Reguler", eta: "3-5 hari kerja", price: 5.0 },
  { id: "ship-2", name: "Express", eta: "1-2 hari kerja", price: 12.0 },
  { id: "ship-3", name: "Same Day", eta: "Hari ini", price: 18.0 },
];

export const paymentOptionsData: PaymentOption[] = [
  { id: "pay-1", name: "Kartu Kredit/Debit" },
  { id: "pay-2", name: "Virtual Account" },
  { id: "pay-3", name: "E-Wallet (GoPay, OVO)" },
];
