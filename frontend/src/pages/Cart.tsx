import React, { useState, useMemo } from "react";
import { cartData } from "../data/payment";
import { CartItem } from "../types";
import {
  FiPlus,
  FiMinus,
  FiTrash2,
  FiShoppingCart,
  FiArrowRight,
} from "react-icons/fi";
import Navbar from "../components/Navbar";

const CartPage: React.FC = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>(cartData);

  const handleQuantityChange = (id: string, amount: number) => {
    setCartItems((currentItems) =>
      currentItems.map((item) =>
        item.id === id
          ? { ...item, quantity: Math.max(1, item.quantity + amount) }
          : item
      )
    );
  };

  const handleRemoveItem = (id: string) => {
    setCartItems((currentItems) =>
      currentItems.filter((item) => item.id !== id)
    );
  };

  const subtotal = useMemo(() => {
    return cartItems.reduce((acc, item) => acc + item.price * item.quantity, 0);
  }, [cartItems]);

  if (cartItems.length === 0) {
    return (
      <>
        <Navbar />
        <div className="bg-gray-50 min-h-screen py-12">
          <main className="container mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
            <div className="bg-white rounded-lg shadow-md text-center p-12">
              <FiShoppingCart className="mx-auto text-blue-500" size={64} />
              <h1 className="mt-6 text-2xl font-bold text-gray-900">
                Keranjang Anda Kosong
              </h1>
              <p className="mt-2 text-gray-600">
                Sepertinya Anda belum menambahkan produk apapun ke keranjang.
              </p>
              <a
                href="/"
                className="
                mt-8 inline-flex items-center justify-center 
                rounded-lg bg-blue-600 px-6 py-3 
                text-base font-medium text-white shadow-sm 
                hover:bg-blue-700 transition-colors
              "
              >
                Mulai Belanja
              </a>
            </div>
          </main>
        </div>
      </>
    );
  }

  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen py-12">
        <main className="container mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">
            Shopping Cart
          </h1>

          <div className="lg:grid lg:grid-cols-3 lg:gap-12 items-start">
            {/* ITEM LIST */}
            <div className="lg:col-span-2 bg-white rounded-lg shadow-md">
              <ul role="list" className="divide-y divide-gray-200">
                {cartItems.map((item) => (
                  <li
                    key={item.id}
                    className="flex flex-col sm:flex-row py-6 px-6"
                  >
                    {/* Gambar */}
                    <div className="flex-shrink-0">
                      <img
                        src={item.imageUrl}
                        alt={item.name}
                        className="w-24 h-24 sm:w-32 sm:h-32 rounded-md object-cover"
                      />
                    </div>

                    <div className="ml-0 sm:ml-6 mt-4 sm:mt-0 flex-1 flex flex-col justify-between">
                      <div>
                        <div className="flex justify-between items-start">
                          <h3 className="text-lg font-medium text-gray-900">
                            <a href={`/product/${item.id}`}>{item.name}</a>
                          </h3>
                          <span className="ml-4 text-lg font-medium text-gray-900">
                            ${(item.price * item.quantity).toFixed(2)}
                          </span>
                        </div>
                        <p className="mt-1 text-sm text-gray-500">
                          Unit Price: ${item.price.toFixed(2)}
                        </p>
                      </div>

                      {/* CTA */}
                      <div className="flex items-center justify-between mt-4">
                        {/* Quantity Control */}
                        <div className="flex items-center border border-gray-300 rounded-md">
                          <button
                            onClick={() => handleQuantityChange(item.id, -1)}
                            className="px-3 py-2 text-gray-600 hover:bg-gray-100 rounded-l-md disabled:opacity-50"
                            disabled={item.quantity <= 1}
                          >
                            <FiMinus size={16} />
                          </button>
                          <span className="px-4 py-1 text-center font-medium text-gray-900">
                            {item.quantity}
                          </span>
                          <button
                            onClick={() => handleQuantityChange(item.id, 1)}
                            className="px-3 py-2 text-gray-600 hover:bg-gray-100 rounded-r-md"
                          >
                            <FiPlus size={16} />
                          </button>
                        </div>

                        <button
                          onClick={() => handleRemoveItem(item.id)}
                          type="button"
                          className="ml-4 inline-flex items-center text-sm font-medium text-red-500 hover:text-red-700"
                        >
                          <FiTrash2 className="mr-1.5" size={16} />
                          Delete
                        </button>
                      </div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>

            {/* ORDER SUMMARY */}
            <div className="lg:col-span-1 mt-10 lg:mt-0">
              <div className="lg:sticky lg:top-24">
                <div className="bg-white rounded-lg shadow-md p-6">
                  <h2 className="text-xl font-semibold text-gray-900 mb-4">
                    Order Summary
                  </h2>

                  <div className="space-y-2">
                    <div className="flex justify-between text-base text-gray-600">
                      <span>Subtotal</span>
                      <span>${subtotal.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-sm text-gray-500">
                      <span>Shipping</span>
                      <span>Calculated at checkout</span>
                    </div>
                    <div className="flex justify-between text-sm text-gray-500">
                      <span>Tax</span>
                      <span>Calculated at checkout</span>
                    </div>
                  </div>

                  <div className="mt-6 pt-4 border-t border-gray-200">
                    <div className="flex justify-between text-lg font-bold text-gray-900">
                      <span>Estimation Total</span>
                      <span>${subtotal.toFixed(2)}</span>
                    </div>
                  </div>

                  {/* CTA */}
                  <a
                    href="/checkout"
                    className="
                    w-full mt-6 py-3 px-4 
                    rounded-lg bg-blue-600 text-white 
                    font-semibold text-lg shadow-md 
                    hover:bg-blue-700 transition-colors
                    flex items-center justify-center
                  "
                  >
                    Proceed to Checkout
                    <FiArrowRight className="ml-2" size={20} />
                  </a>

                  <div className="mt-4 text-center">
                    <a
                      href="/"
                      className="text-sm font-medium text-blue-600 hover:text-blue-500"
                    >
                      or Continue Shopping
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </>
  );
};

export default CartPage;
