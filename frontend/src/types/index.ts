export interface Review {
  id: string;
  author: string;
  rating: number;
  comment: string;
  date: string;
}

export interface Product {
  id: string;
  name: string;
  category: string;
  price: number;
  oldPrice?: number;
  rating: number;
  imageUrl: string;
  isNew?: boolean;
}

export interface DetailedProduct extends Product {
  images: string[];
  stock: number;
  description: string;
  variants: {
    type: string;
    options: string[];
  };
  reviews: Review[];
}

export interface CartItem {
  id: string;
  name: string;
  price: number;
  quantity: number;
  imageUrl: string;
}

export interface ShippingOption {
  id: string;
  name: string;
  eta: string;
  price: number;
}

export interface PaymentOption {
  id: string;
  name: string;
}

export type OrderStatus =
  | "created"
  | "processing"
  | "shipping"
  | "completed"
  | "cancelled"
  | "returned";

export interface OrderItem {
  id: string;
  name: string;
  imageUrl: string;
  quantity: number;
  price: number;
}

export interface Order {
  id: string;
  date: string;
  status: OrderStatus;
  items: OrderItem[];
  total: number;
  subtotal: number;
  shippingCost: number;
  trackingNumber?: string;
  estimatedDelivery?: string;

  shippingAddress: UserAddress;
  paymentMethod: string;
  statusHistory: {
    status: OrderStatus;
    date: string;
  }[];
}

export interface UserAddress {
  id: string;
  label: string;
  isDefault: boolean;
  recipientName: string;
  phone: string;
  address: string;
  city: string;
  postalCode: string;
}

export interface User {
  id: string;
  fullName: string;
  email: string;
  avatarUrl?: string;
  isConfirmed: boolean;
  memberSince: string;
  addresses: UserAddress[];
}
