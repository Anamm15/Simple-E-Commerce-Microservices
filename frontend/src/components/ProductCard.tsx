import React from "react";
import { Product } from "../types";
import { FaStar, FaStarHalfAlt } from "react-icons/fa";
import { FiShoppingCart } from "react-icons/fi";
import { useNavigate } from "react-router-dom";
interface ProductCardProps {
  product: Product;
}

const Rating: React.FC<{ rating: number }> = ({ rating }) => {
  const fullStars = Math.floor(rating);
  const hasHalfStar = rating % 1 >= 0.5;
  const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);

  return (
    <div className="flex items-center text-yellow-400">
      {[...Array(fullStars)].map((_, i) => (
        <FaStar key={`full-${i}`} />
      ))}
      {hasHalfStar && <FaStarHalfAlt />}
      {[...Array(emptyStars)].map((_, i) => (
        <FaStar key={`empty-${i}`} className="text-gray-300" />
      ))}
      <span className="ml-2 text-sm text-gray-500">({rating})</span>
    </div>
  );
};

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const { name, category, price, oldPrice, rating, imageUrl, isNew } = product;
  const navigate = useNavigate();
  const discountPercent = oldPrice
    ? Math.round(((oldPrice - price) / oldPrice) * 100)
    : 0;

  return (
    <div
      className="
      group relative w-full max-w-sm 
      bg-white rounded-lg shadow-md 
      overflow-hidden cursor-pointer 
      transition-all duration-300 hover:shadow-xl
    "
      onClick={() => navigate(`/product/${product.id}`)}
    >
      {/* IMAGE N BADGE */}
      <div className="relative overflow-hidden h-64">
        <img
          src={imageUrl}
          alt={name}
          className="w-full h-full object-cover 
                     transition-transform duration-500 ease-in-out 
                     group-hover:scale-110"
        />

        {/* Badge */}
        {discountPercent > 0 ? (
          <span className="absolute top-3 left-3 bg-red-500 text-white text-xs font-semibold px-2 py-1 rounded">
            -{discountPercent}%
          </span>
        ) : isNew ? (
          <span className="absolute top-3 left-3 bg-blue-500 text-white text-xs font-semibold px-2 py-1 rounded">
            BARU
          </span>
        ) : null}
      </div>

      {/* PRODUCT INFORMATION */}
      <div className="p-5">
        <h3 className="text-xs text-gray-500 uppercase tracking-wide">
          {category}
        </h3>
        <h2
          className="mt-1 text-lg font-semibold text-gray-900 truncate"
          title={name}
        >
          {name}
        </h2>

        <div className="mt-2">
          <Rating rating={rating} />
        </div>

        <div className="mt-3 flex items-baseline space-x-2">
          <span className="text-2xl font-bold text-gray-900">
            ${price.toFixed(2)}
          </span>
          {oldPrice && (
            <span className="text-sm text-gray-400 line-through">
              ${oldPrice.toFixed(2)}
            </span>
          )}
        </div>
      </div>

      {/* CTA */}
      <div
        className="
        absolute bottom-0 left-0 right-0 p-4 
        bg-white/70 backdrop-blur-sm 
        transform translate-y-full 
        transition-transform duration-300 ease-in-out 
        group-hover:translate-y-0
      "
      >
        <button
          className="
          w-full flex items-center justify-center 
          bg-gray-900 text-white text-sm font-medium 
          py-2.5 px-4 rounded-lg 
          hover:bg-blue-600 transition-colors
        "
        >
          <FiShoppingCart className="mr-2" size={18} />
          Tambah ke Keranjang
        </button>
      </div>
    </div>
  );
};

export default ProductCard;
