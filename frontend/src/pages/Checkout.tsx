import React, { useState, useMemo } from "react";
import {
  cartData,
  shippingOptionsData,
  paymentOptionsData,
} from "../data/payment";
import {
  FiCreditCard,
  FiPackage,
  FiSmartphone,
  FiCheckCircle,
  FiLock,
} from "react-icons/fi";
import Navbar from "../components/Navbar";

const paymentIcons: { [key: string]: React.ReactElement } = {
  "pay-1": <FiCreditCard className="text-blue-500" size={24} />,
  "pay-2": <FiPackage className="text-green-500" size={24} />,
  "pay-3": <FiSmartphone className="text-purple-500" size={24} />,
};

const CheckoutPage: React.FC = () => {
  const [email, setEmail] = useState("");
  const [fullName, setFullName] = useState("");
  const [address, setAddress] = useState("");
  const [city, setCity] = useState("");
  const [postalCode, setPostalCode] = useState("");

  const [selectedShipping, setSelectedShipping] = useState(
    shippingOptionsData[0].id
  );
  const [selectedPayment, setSelectedPayment] = useState(
    paymentOptionsData[0].id
  );

  const { subtotal, shippingCost, total } = useMemo(() => {
    const subtotal = cartData.reduce(
      (acc, item) => acc + item.price * item.quantity,
      0
    );
    const shippingCost =
      shippingOptionsData.find((opt) => opt.id === selectedShipping)?.price ||
      0;
    const total = subtotal + shippingCost;
    return { subtotal, shippingCost, total };
  }, [selectedShipping]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    alert("Pesanan berhasil dibuat! (Simulasi)");
    console.log({
      email,
      fullName,
      address,
      city,
      postalCode,
      selectedShipping,
      selectedPayment,
      total,
    });
  };

  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen">
        <main className="container mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-12">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">Checkout</h1>

          <form
            onSubmit={handleSubmit}
            className="lg:grid lg:grid-cols-3 lg:gap-12"
          >
            {/* ========== 1. LEFT COLUMN: FORM ========== */}
            <div className="lg:col-span-2 space-y-8">
              <section className="bg-white p-6 rounded-lg shadow-md">
                <h2 className="text-xl font-semibold text-gray-900 mb-4">
                  Shipment Information
                </h2>
                <div className="grid grid-cols-1 gap-y-4 sm:grid-cols-2 sm:gap-x-6">
                  <div>
                    <label
                      htmlFor="email"
                      className="block text-sm font-medium text-gray-700"
                    >
                      Email
                    </label>
                    <input
                      type="email"
                      id="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    />
                  </div>
                  <div>
                    <label
                      htmlFor="fullName"
                      className="block text-sm font-medium text-gray-700"
                    >
                      Full Name
                    </label>
                    <input
                      type="text"
                      id="fullName"
                      value={fullName}
                      onChange={(e) => setFullName(e.target.value)}
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    />
                  </div>
                  <div className="sm:col-span-2">
                    <label
                      htmlFor="address"
                      className="block text-sm font-medium text-gray-700"
                    >
                      Address
                    </label>
                    <textarea
                      id="address"
                      rows={3}
                      value={address}
                      onChange={(e) => setAddress(e.target.value)}
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    />
                  </div>
                  <div>
                    <label
                      htmlFor="city"
                      className="block text-sm font-medium text-gray-700"
                    >
                      City
                    </label>
                    <input
                      type="text"
                      id="city"
                      value={city}
                      onChange={(e) => setCity(e.target.value)}
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    />
                  </div>
                  <div>
                    <label
                      htmlFor="postalCode"
                      className="block text-sm font-medium text-gray-700"
                    >
                      Postal Code
                    </label>
                    <input
                      type="text"
                      id="postalCode"
                      value={postalCode}
                      onChange={(e) => setPostalCode(e.target.value)}
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                      required
                    />
                  </div>
                </div>
              </section>

              <section className="bg-white p-6 rounded-lg shadow-md">
                <h2 className="text-xl font-semibold text-gray-900 mb-4">
                  Shipping Method
                </h2>
                <div className="space-y-4">
                  {shippingOptionsData.map((option) => (
                    <label
                      key={option.id}
                      className={`
                      relative flex justify-between items-center 
                      border rounded-lg p-4 cursor-pointer 
                      transition-all
                      ${
                        selectedShipping === option.id
                          ? "border-blue-600 ring-2 ring-blue-500 bg-blue-50"
                          : "border-gray-300"
                      }
                    `}
                    >
                      <input
                        type="radio"
                        name="shipping"
                        value={option.id}
                        checked={selectedShipping === option.id}
                        onChange={() => setSelectedShipping(option.id)}
                        className="peer absolute -z-10 opacity-0"
                      />
                      <div>
                        <span className="block text-sm font-medium text-gray-900">
                          {option.name}
                        </span>
                        <span className="block text-sm text-gray-500">
                          {option.eta}
                        </span>
                      </div>
                      <span className="text-sm font-semibold text-gray-900">
                        ${option.price.toFixed(2)}
                      </span>
                      <FiCheckCircle
                        className={`
                        absolute top-4 right-4 text-blue-600 
                        transition-opacity
                        ${
                          selectedShipping === option.id
                            ? "opacity-100"
                            : "opacity-0"
                        }
                      `}
                        size={20}
                      />
                    </label>
                  ))}
                </div>
              </section>

              <section className="bg-white p-6 rounded-lg shadow-md">
                <h2 className="text-xl font-semibold text-gray-900 mb-4">
                  Payment Method
                </h2>
                <div className="space-y-4">
                  {paymentOptionsData.map((option) => (
                    <label
                      key={option.id}
                      className={`
                      relative flex items-center 
                      border rounded-lg p-4 cursor-pointer 
                      transition-all
                      ${
                        selectedPayment === option.id
                          ? "border-blue-600 ring-2 ring-blue-500 bg-blue-50"
                          : "border-gray-300"
                      }
                    `}
                    >
                      <input
                        type="radio"
                        name="payment"
                        value={option.id}
                        checked={selectedPayment === option.id}
                        onChange={() => setSelectedPayment(option.id)}
                        className="peer absolute -z-10 opacity-0"
                      />
                      {paymentIcons[option.id]}
                      <span className="ml-4 block text-sm font-medium text-gray-900">
                        {option.name}
                      </span>
                      <FiCheckCircle
                        className={`
                        absolute top-1/2 -translate-y-1/2 right-4 text-blue-600 
                        transition-opacity
                        ${
                          selectedPayment === option.id
                            ? "opacity-100"
                            : "opacity-0"
                        }
                      `}
                        size={20}
                      />
                    </label>
                  ))}
                </div>
              </section>
            </div>

            {/* ========== 2. RIGHT COLUMN: ORDER SUMMARY ========== */}
            <div className="lg:col-span-1 mt-10 lg:mt-0">
              <div className="lg:sticky lg:top-24">
                <div className="bg-white rounded-lg shadow-md p-6">
                  <h2 className="text-xl font-semibold text-gray-900 mb-4">
                    Order Summary
                  </h2>

                  <div className="space-y-4">
                    {cartData.map((item) => (
                      <div
                        key={item.id}
                        className="flex items-start justify-between"
                      >
                        <div className="flex items-start">
                          <img
                            src={item.imageUrl}
                            alt={item.name}
                            className="w-16 h-16 rounded-md object-cover mr-4"
                          />
                          <div>
                            <h3 className="text-sm font-medium text-gray-900">
                              {item.name}
                            </h3>
                            <p className="text-sm text-gray-500">
                              Qty: {item.quantity}
                            </p>
                          </div>
                        </div>
                        <span className="text-sm font-medium text-gray-900">
                          ${(item.price * item.quantity).toFixed(2)}
                        </span>
                      </div>
                    ))}
                  </div>

                  <div className="mt-6 pt-6 border-t border-gray-200 space-y-2">
                    <div className="flex justify-between text-sm text-gray-600">
                      <span>Subtotal</span>
                      <span>${subtotal.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-sm text-gray-600">
                      <span>Shipment</span>
                      <span>${shippingCost.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-lg font-semibold text-gray-900 mt-2">
                      <span>Total</span>
                      <span>${total.toFixed(2)}</span>
                    </div>
                  </div>

                  {/* CTA */}
                  <button
                    type="submit"
                    className="
                    w-full mt-6 py-3 px-4 
                    rounded-lg bg-blue-600 text-white 
                    font-semibold text-lg shadow-md 
                    hover:bg-blue-700 transition-colors
                    flex items-center justify-center
                  "
                  >
                    <FiLock className="mr-2" size={18} />
                    Pay Now
                  </button>
                </div>
              </div>
            </div>
          </form>
        </main>
      </div>
    </>
  );
};

export default CheckoutPage;
