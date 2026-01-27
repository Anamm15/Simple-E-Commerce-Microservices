import React, { useState } from "react";
import { FiPlus, FiMinus, FiCheckCircle, FiAlertCircle } from "react-icons/fi";
import { FaStar } from "react-icons/fa";

import { detailedProductData } from "../data/detail-product";
import Navbar from "../components/Navbar";

const StarRatingDisplay: React.FC<{ rating: number }> = ({ rating }) => {
  return (
    <div className="flex items-center">
      {[...Array(5)].map((_, i) => (
        <FaStar
          key={i}
          className={
            i < Math.round(rating) ? "text-yellow-400" : "text-gray-300"
          }
        />
      ))}
      <span className="ml-2 text-sm text-gray-600">
        {rating.toFixed(1)} ({detailedProductData.reviews.length} review)
      </span>
    </div>
  );
};

const ReviewForm: React.FC = () => {
  const [rating, setRating] = useState(0);
  const [hoverRating, setHoverRating] = useState(0);
  const [comment, setComment] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Logika untuk submit review (misal: kirim ke API)
    console.log({ rating, comment });
    alert("Review Anda telah dikirim!");
    setRating(0);
    setComment("");
  };

  return (
    <form onSubmit={handleSubmit} className="mt-8">
      <h3 className="text-xl font-semibold text-gray-900">Tulis Review Anda</h3>
      <div className="mt-4">
        <label className="block text-sm font-medium text-gray-700">
          Rating Anda
        </label>
        <div className="flex space-x-1 mt-1">
          {[...Array(5)].map((_, index) => {
            const starValue = index + 1;
            return (
              <FaStar
                key={starValue}
                size={24}
                className={`cursor-pointer ${
                  starValue <= (hoverRating || rating)
                    ? "text-yellow-400"
                    : "text-gray-300"
                }`}
                onClick={() => setRating(starValue)}
                onMouseEnter={() => setHoverRating(starValue)}
                onMouseLeave={() => setHoverRating(0)}
              />
            );
          })}
        </div>
      </div>
      <div className="mt-4">
        <label
          htmlFor="comment"
          className="block text-sm font-medium text-gray-700"
        >
          Komentar Anda
        </label>
        <textarea
          id="comment"
          name="comment"
          rows={4}
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
          placeholder="Bagikan pemikiran Anda tentang produk ini..."
          required
        ></textarea>
      </div>
      <button
        type="submit"
        className="mt-4 inline-flex items-center justify-center rounded-lg border border-transparent bg-blue-600 px-6 py-2 text-base font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
      >
        Kirim Review
      </button>
    </form>
  );
};

const ProductDetailPage: React.FC = () => {
  const product = detailedProductData;

  const [selectedImageIndex, setSelectedImageIndex] = useState(0);
  const [selectedVariant, setSelectedVariant] = useState(
    product.variants.options[0],
  );
  const [quantity, setQuantity] = useState(1);
  const [activeTab, setActiveTab] = useState<"description" | "reviews">(
    "description",
  );

  const handleQuantityChange = (amount: number) => {
    setQuantity((prev) => {
      const newQty = prev + amount;
      if (newQty < 1) return 1;
      if (newQty > product.stock) return product.stock;
      return newQty;
    });
  };

  return (
    <>
      <Navbar />
      <div className="bg-white">
        <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
            <div className="flex flex-col">
              <div className="w-full h-96 md:h-[650px] bg-gray-100 rounded-lg overflow-hidden flex items-center justify-center">
                <img
                  src={product.images[selectedImageIndex]}
                  alt={`${product.name} view ${selectedImageIndex + 1}`}
                  className="w-full h-full object-cover transition-opacity duration-300"
                />
              </div>

              <div className="flex space-x-2 mt-4 overflow-x-auto">
                {product.images.map((img, index) => (
                  <button
                    key={index}
                    onClick={() => setSelectedImageIndex(index)}
                    className={`
                    flex-shrink-0 w-20 h-20 rounded-md overflow-hidden
                    border-2 
                    ${
                      selectedImageIndex === index
                        ? "border-blue-500"
                        : "border-transparent"
                    }
                    focus:outline-none focus:ring-2 focus:ring-blue-500
                  `}
                  >
                    <img
                      src={img}
                      alt={`Thumbnail ${index + 1}`}
                      className="w-full h-full object-cover"
                    />
                  </button>
                ))}
              </div>
            </div>

            <div>
              <span className="text-sm font-medium text-blue-600 uppercase">
                {product.category}
              </span>

              <h1 className="mt-2 text-3xl lg:text-4xl font-bold text-gray-900">
                {product.name}
              </h1>

              <div className="mt-4">
                <StarRatingDisplay rating={product.rating} />
              </div>

              <div className="mt-4 flex items-baseline space-x-3">
                <span className="text-3xl font-bold text-gray-900">
                  ${product.price.toFixed(2)}
                </span>
                {product.oldPrice && (
                  <span className="text-xl text-gray-400 line-through">
                    ${product.oldPrice.toFixed(2)}
                  </span>
                )}
              </div>

              <div className="mt-6">
                {product.stock > 0 ? (
                  <span className="flex items-center text-green-600">
                    <FiCheckCircle className="mr-2" />
                    Available ({product.stock} remaining)
                  </span>
                ) : (
                  <span className="flex items-center text-red-600">
                    <FiAlertCircle className="mr-2" />
                    Out of Stock
                  </span>
                )}
              </div>

              <div className="mt-6">
                <h3 className="text-sm font-medium text-gray-900">
                  {product.variants.type}
                </h3>
                <div className="flex flex-wrap gap-3 mt-2">
                  {product.variants.options.map((option) => (
                    <label key={option} className="relative">
                      <input
                        type="radio"
                        name="variant"
                        value={option}
                        checked={selectedVariant === option}
                        onChange={() => setSelectedVariant(option)}
                        className="peer absolute -z-10 opacity-0"
                      />
                      <span
                        className={`
                        flex items-center justify-center w-12 h-10 
                        border rounded-md cursor-pointer
                        text-sm font-medium
                        transition-colors
                        peer-checked:bg-blue-600 peer-checked:text-white peer-checked:border-blue-600
                        hover:bg-gray-100
                      `}
                      >
                        {option}
                      </span>
                    </label>
                  ))}
                </div>
              </div>

              <div className="mt-8">
                <label className="text-sm font-medium text-gray-900">
                  Quantity
                </label>
                <div className="flex items-center border border-gray-300 rounded-md w-32 mt-2">
                  <button
                    onClick={() => handleQuantityChange(-1)}
                    className="px-3 py-2 text-gray-600 hover:bg-gray-100 rounded-l-md"
                    disabled={quantity <= 1}
                  >
                    <FiMinus />
                  </button>
                  <span className="flex-1 text-center font-medium">
                    {quantity}
                  </span>
                  <button
                    onClick={() => handleQuantityChange(1)}
                    className="px-3 py-2 text-gray-600 hover:bg-gray-100 rounded-r-md"
                    disabled={quantity >= product.stock}
                  >
                    <FiPlus />
                  </button>
                </div>
              </div>

              {/* (CTA) */}
              <div className="mt-8 flex flex-col sm:flex-row gap-4">
                <button
                  className="
                  flex-1 flex items-center justify-center 
                  bg-blue-600 text-white text-base font-medium 
                  py-3 px-8 rounded-lg shadow-md 
                  hover:bg-blue-700 transition-colors
                  disabled:opacity-50
                "
                  disabled={product.stock === 0}
                >
                  Add to Cart
                </button>
                <button
                  className="
                  flex-1 flex items-center justify-center 
                  bg-transparent text-blue-600 border border-blue-600 
                  text-base font-medium py-3 px-8 rounded-lg
                  hover:bg-blue-50 transition-colors
                  disabled:opacity-50
                "
                  disabled={product.stock === 0}
                >
                  Buy Now
                </button>
              </div>
            </div>
          </div>

          {/* DESCRIPTION AND REVIEWS */}
          <div className="mt-16">
            <div className="border-b border-gray-200">
              <nav className="-mb-px flex space-x-8" aria-label="Tabs">
                <button
                  onClick={() => setActiveTab("description")}
                  className={`
                  whitespace-nowrap py-4 px-1 border-b-2 
                  font-medium text-lg
                  ${
                    activeTab === "description"
                      ? "border-blue-500 text-blue-600"
                      : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
                  }
                `}
                >
                  Description
                </button>
                <button
                  onClick={() => setActiveTab("reviews")}
                  className={`
                  whitespace-nowrap py-4 px-1 border-b-2 
                  font-medium text-lg
                  ${
                    activeTab === "reviews"
                      ? "border-blue-500 text-blue-600"
                      : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
                  }
                `}
                >
                  Review ({product.reviews.length})
                </button>
              </nav>
            </div>

            <div className="mt-8">
              {activeTab === "description" && (
                <div
                  className="prose prose-lg text-gray-700"
                  dangerouslySetInnerHTML={{ __html: product.description }}
                />
              )}

              {activeTab === "reviews" && (
                <div>
                  <h2 className="text-2xl font-bold text-gray-900">
                    Customer Reviews
                  </h2>

                  <div className="mt-6 space-y-6">
                    {product.reviews.length > 0 ? (
                      product.reviews.map((review) => (
                        <div
                          key={review.id}
                          className="pb-6 border-b border-gray-200"
                        >
                          <div className="flex items-center mb-2">
                            <h4 className="text-base font-semibold text-gray-900">
                              {review.author}
                            </h4>
                            <span className="ml-4 text-sm text-gray-500">
                              {review.date}
                            </span>
                          </div>
                          <StarRatingDisplay rating={review.rating} />
                          <p className="mt-3 text-gray-700">{review.comment}</p>
                        </div>
                      ))
                    ) : (
                      <p className="text-gray-500">
                        No reviews for this product yet.
                      </p>
                    )}
                  </div>

                  <div className="mt-10">
                    <ReviewForm />
                  </div>
                </div>
              )}
            </div>
          </div>
        </main>
      </div>
    </>
  );
};

export default ProductDetailPage;
