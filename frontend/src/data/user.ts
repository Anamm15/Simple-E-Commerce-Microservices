import { User } from "../types";

export const dummyUser: User = {
  id: "u1",
  fullName: "Budi Santoso",
  email: "budi.santoso@example.com",
  avatarUrl: "https://example.com/avatar.jpg",
  isConfirmed: true,
  memberSince: "15 Agustus 2025",
  addresses: [
    {
      id: "a1",
      label: "Rumah",
      isDefault: true,
      recipientName: "Budi Santoso",
      phone: "081234567890",
      address: "Jl. Jenderal Sudirman No. 123, Apartemen Cendana Lt. 5",
      city: "Jakarta Selatan",
      postalCode: "12190",
    },
    {
      id: "a2",
      label: "Kantor",
      isDefault: false,
      recipientName: "Budi (Work)",
      phone: "081987654321",
      address: "Gedung Cyber 2, Lantai 10, Jl. H.R. Rasuna Said",
      city: "Jakarta Selatan",
      postalCode: "12950",
    },
  ],
};
