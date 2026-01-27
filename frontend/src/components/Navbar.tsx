import React, { useState } from "react";
import {
  FiShoppingCart,
  FiUser,
  FiSearch,
  FiMenu,
  FiX,
  FiShoppingBag,
} from "react-icons/fi";

const Navbar: React.FC = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const cartItemCount = 3;

  return (
    <nav
      className="
      sticky top-0 z-50 w-full
      bg-white/70 backdrop-blur-md 
      shadow-sm
    "
    >
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex-shrink-0">
            <a
              href="/"
              className="text-2xl font-bold text-gray-900 hover:text-blue-600 transition-colors"
            >
              E-Store
            </a>
          </div>

          <div className="hidden md:flex md:space-x-8">
            <a
              href="/category/men"
              className="text-gray-600 hover:text-blue-600 transition-colors font-medium"
            >
              Pria
            </a>
            <a
              href="/category/women"
              className="text-gray-600 hover:text-blue-600 transition-colors font-medium"
            >
              Wanita
            </a>
            <a
              href="/category/electronics"
              className="text-gray-600 hover:text-blue-600 transition-colors font-medium"
            >
              Elektronik
            </a>
            <a
              href="/deals"
              className="text-red-500 hover:text-red-700 transition-colors font-medium"
            >
              Promo
            </a>
          </div>

          <div className="flex items-center space-x-4">
            <button className="hidden md:block text-gray-500 hover:text-blue-600 transition-colors">
              <FiSearch size={22} />
            </button>

            <a
              href="/cart"
              className="relative text-gray-500 hover:text-blue-600 transition-colors"
            >
              <FiShoppingCart size={22} />
              {cartItemCount > 0 && (
                <span
                  className="
                  absolute -top-2 -right-2 w-5 h-5 
                  rounded-full bg-red-500 text-white 
                  text-xs flex items-center justify-center
                "
                >
                  {cartItemCount}
                </span>
              )}
            </a>

            <a
              href="/order"
              className="relative text-gray-500 hover:text-blue-600 transition-colors"
            >
              <FiShoppingBag size={22} />
            </a>

            <a
              href="/profile"
              className="hidden md:block text-gray-500 hover:text-blue-600 transition-colors"
            >
              <FiUser size={22} />
            </a>

            <div className="md:hidden">
              <button
                onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                className="text-gray-600 focus:outline-none"
              >
                {isMobileMenuOpen ? <FiX size={24} /> : <FiMenu size={24} />}
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* --- MOBILE MENU --- */}
      <div
        className={`
          md:hidden absolute top-16 left-0 w-full 
          bg-white/95 backdrop-blur-md shadow-lg
          transition-all duration-300 ease-in-out
          ${
            isMobileMenuOpen
              ? "translate-y-0 opacity-100"
              : "-translate-y-4 opacity-0 pointer-events-none"
          }
        `}
      >
        <div className="px-4 pt-2 pb-4 space-y-2">
          <div className="relative">
            <input
              type="text"
              placeholder="Cari produk..."
              className="
                w-full pl-10 pr-4 py-2 
                rounded-full border border-gray-300 
                focus:outline-none focus:ring-2 focus:ring-blue-500
              "
            />
            <span className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
              <FiSearch size={20} />
            </span>
          </div>

          <a
            href="/category/men"
            className="block py-2 px-3 text-lg text-gray-700 hover:bg-gray-100 rounded-md"
          >
            Pria
          </a>
          <a
            href="/category/women"
            className="block py-2 px-3 text-lg text-gray-700 hover:bg-gray-100 rounded-md"
          >
            Wanita
          </a>
          <a
            href="/category/electronics"
            className="block py-2 px-3 text-lg text-gray-700 hover:bg-gray-100 rounded-md"
          >
            Elektronik
          </a>
          <a
            href="/deals"
            className="block py-2 px-3 text-lg text-red-500 hover:bg-gray-100 rounded-md"
          >
            Promo
          </a>

          <hr className="my-2 border-gray-200" />

          <a
            href="/profile"
            className="block py-2 px-3 text-lg text-gray-700 hover:bg-gray-100 rounded-md"
          >
            Akun Saya
          </a>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
