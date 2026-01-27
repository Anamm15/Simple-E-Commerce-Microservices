import { Order } from "../types";
import { dummyUser } from "./user";
const defaultAddress = dummyUser.addresses.find((a) => a.isDefault);

export const dummyDetailedOrder: Order = {
  id: "ORD-77449",
  date: "30 Oktober 2025, 10:30 WIB",
  status: "shipping",

  statusHistory: [
    { status: "created", date: "30 Oktober 2025, 10:30 WIB" },
    { status: "processing", date: "31 Oktober 2025, 09:00 WIB" },
    { status: "shipping", date: "31 Oktober 2025, 17:00 WIB" },
    { status: "completed", date: "" },
  ],

  items: [
    {
      id: "p4",
      name: "Premium Running Shoes",
      imageUrl: "https://placehold.co/100x100/e83e8c/white?text=Sepatu",
      quantity: 1,
      price: 120.0,
    },
  ],

  subtotal: 120.0,
  shippingCost: 12.0,
  total: 132.0,

  trackingNumber: "JN-1002348B",
  estimatedDelivery: "3 November 2025",

  shippingAddress: defaultAddress!,
  paymentMethod: "Kartu Kredit (Visa **** 4321)",
};
