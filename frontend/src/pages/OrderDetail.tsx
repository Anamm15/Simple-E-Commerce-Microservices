import React from "react";
import { dummyDetailedOrder } from "../data/detail-order";
import { Order, OrderStatus } from "../types";
import {
  FiClipboard,
  FiLoader,
  FiTruck,
  FiCheckCircle,
  FiArrowRight,
  FiMapPin,
  FiCreditCard,
} from "react-icons/fi";
import Navbar from "../components/Navbar";

const statusConfig = {
  created: {
    icon: FiClipboard,
    title: "Order Created",
    color: "text-blue-500",
  },
  processing: {
    icon: FiLoader,
    title: "Order Processing",
    color: "text-yellow-500",
  },
  shipping: {
    icon: FiTruck,
    title: "Order Shipping",
    color: "text-blue-500",
  },
  completed: {
    icon: FiCheckCircle,
    title: "Order Completed",
    color: "text-green-500",
  },
};

const getStatusIndex = (status: OrderStatus) => {
  const order: OrderStatus[] = [
    "created",
    "processing",
    "shipping",
    "completed",
  ];
  return order.indexOf(status);
};

const OrderDetailPage: React.FC = () => {
  const order = dummyDetailedOrder;
  const currentStatusIndex = getStatusIndex(order.status);

  const OrderTracker: React.FC = () => {
    const steps: OrderStatus[] = [
      "created",
      "processing",
      "shipping",
      "completed",
    ];

    return (
      <nav aria-label="Order tracker">
        <ol className="relative flex flex-col sm:flex-row justify-between w-full">
          {steps.map((status, index) => {
            const config = statusConfig[status];
            const historyEntry = order.statusHistory.find(
              (h) => h.status === status
            );

            const isCompleted = index < currentStatusIndex;
            const isActive = index === currentStatusIndex;
            const isUpcoming = index > currentStatusIndex;

            return (
              <li key={status} className="flex-1 relative sm:text-center p-2">
                {index < steps.length - 1 && (
                  <div
                    className={`
                      hidden sm:block absolute top-1/2 left-1/2 w-full h-0.5 
                      -translate-y-1/2
                      ${isCompleted || isActive ? "bg-blue-600" : "bg-gray-300"}
                    `}
                  />
                )}

                <div className="relative z-10 flex flex-row sm:flex-col items-center sm:justify-center">
                  <span
                    className={`
                      flex items-center justify-center w-12 h-12 rounded-full 
                      ${isCompleted ? "bg-blue-600 text-white" : ""}
                      ${isActive ? "bg-blue-600 text-white animate-pulse" : ""}
                      ${isUpcoming ? "bg-gray-200 text-gray-500" : ""}
                    `}
                  >
                    <config.icon size={24} />
                  </span>
                  <div className="ml-4 sm:ml-0 sm:mt-3 text-left sm:text-center">
                    <h3 className="mt-2 text-sm sm:text-base font-semibold text-gray-900">
                      {config.title}
                    </h3>
                    <p className="text-xs sm:text-sm text-gray-500">
                      {historyEntry?.date || "Not yet"}
                    </p>
                  </div>
                </div>
              </li>
            );
          })}
        </ol>
      </nav>
    );
  };

  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen py-12">
        <main className="container mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="mb-8">
            <h1 className="text-3xl font-bold text-gray-900">Order Details</h1>
            <p className="mt-1 text-lg text-gray-600">Order ID: #{order.id}</p>
          </div>

          <div className="lg:grid lg:grid-cols-3 lg:gap-12 items-start">
            {/* ========== 1. LEFT COLUMN: TRACKER & ITEM ========= */}
            <div className="lg:col-span-2 space-y-8">
              {/* --- TRACKER CARD --- */}
              <section className="bg-white p-6 rounded-lg shadow-md">
                <OrderTracker />
                {order.status === "shipping" && (
                  <div className="mt-6 p-4 bg-blue-50 rounded-lg">
                    <p className="text-sm text-blue-700">
                      <span className="font-semibold">Estimated Delivery:</span>{" "}
                      {order.estimatedDelivery}
                    </p>
                    <p className="text-sm text-blue-700 mt-1">
                      <span className="font-semibold">Tracking Number:</span>{" "}
                      {order.trackingNumber}
                    </p>
                    <button className="mt-2 py-1.5 px-10 text-white rounded-md font-medium bg-blue-600 hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 duration-200 active:scale-95">
                      Track
                    </button>
                  </div>
                )}
              </section>

              {/* --- ORDER ITEMS --- */}
              <section className="bg-white rounded-lg shadow-md">
                <h2 className="text-xl font-semibold text-gray-900 p-6 border-b border-gray-200">
                  Ordered Products
                </h2>
                <ul role="list" className="divide-y divide-gray-200">
                  {order.items.map((item) => (
                    <li key={item.id} className="flex py-6 px-6">
                      <img
                        src={item.imageUrl}
                        alt={item.name}
                        className="w-24 h-24 rounded-md object-cover"
                      />
                      <div className="ml-4 flex-1 flex flex-col justify-between">
                        <div>
                          <h3 className="text-base font-medium text-gray-900">
                            {item.name}
                          </h3>
                          <p className="mt-1 text-sm text-gray-500">
                            Qty: {item.quantity}
                          </p>
                        </div>
                        <span className="mt-2 text-base font-medium text-gray-900">
                          ${item.price.toFixed(2)}
                        </span>
                      </div>
                    </li>
                  ))}
                </ul>
              </section>
            </div>

            {/* ========== 2. RIGHT COLUMN: SUMMARY & ADDRESS ========== */}
            <div className="lg:col-span-1 mt-10 lg:mt-0">
              <div className="lg:sticky lg:top-24 space-y-6">
                {/* --- Summary Card --- */}
                <div className="bg-white rounded-lg shadow-md p-6">
                  <h2 className="text-lg font-semibold text-gray-900 mb-4">
                    Cost Summary
                  </h2>
                  <div className="space-y-2">
                    <div className="flex justify-between text-sm text-gray-600">
                      <span>Subtotal</span>
                      <span>${order.subtotal.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-sm text-gray-600">
                      <span>Shipping</span>
                      <span>${order.shippingCost.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-lg font-bold text-gray-900 mt-2 pt-2 border-t">
                      <span>Total Paid</span>
                      <span>${order.total.toFixed(2)}</span>
                    </div>
                  </div>
                </div>

                {/* --- Shipping Address Card --- */}
                <div className="bg-white rounded-lg shadow-md p-6">
                  <h2 className="flex items-center text-lg font-semibold text-gray-900 mb-4">
                    <FiMapPin className="mr-2" /> Shipping Address
                  </h2>
                  <div className="text-sm text-gray-700 space-y-1">
                    <p className="font-medium">
                      {order.shippingAddress.recipientName}
                    </p>
                    <p>{order.shippingAddress.phone}</p>
                    <p>{order.shippingAddress.address}</p>
                    <p>
                      {order.shippingAddress.city},{" "}
                      {order.shippingAddress.postalCode}
                    </p>
                  </div>
                </div>

                {/* --- Payment Method Card --- */}
                <div className="bg-white rounded-lg shadow-md p-6">
                  <h2 className="flex items-center text-lg font-semibold text-gray-900 mb-4">
                    <FiCreditCard className="mr-2" /> Payment Method
                  </h2>
                  <p className="text-sm text-gray-700">{order.paymentMethod}</p>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </>
  );
};

export default OrderDetailPage;
