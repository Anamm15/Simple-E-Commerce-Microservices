import React, { useState, useMemo } from "react";
import { OrderStatus } from "../types";
import { dummyOrders } from "../data/order";
import OrderCard from "../components/OrderCard";
import { FiArchive } from "react-icons/fi";
import Navbar from "../components/Navbar";

const OrderTrackingPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<OrderStatus>("processing");

  const filteredOrders = useMemo(() => {
    return dummyOrders.filter((order) => order.status === activeTab);
  }, [activeTab]);

  const tabs: { id: OrderStatus; label: string }[] = [
    { id: "processing", label: "Processing" },
    { id: "shipping", label: "Shipping" },
    { id: "completed", label: "Completed" },
    { id: "cancelled", label: "Cancelled" },
    { id: "returned", label: "Returned" },
  ];

  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen py-12">
        <main className="container mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">
            Your Orders Status
          </h1>

          {/* --- NAVIGATION TABS --- */}
          <div className="border-b border-gray-200 mb-8">
            <nav className="-mb-px flex space-x-8" aria-label="Tabs">
              {tabs.map((tab) => (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`
                  whitespace-nowrap py-4 px-1 border-b-2 
                  font-medium text-lg
                  transition-colors
                  ${
                    activeTab === tab.id
                      ? "border-blue-500 text-blue-600"
                      : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
                  }
                `}
                >
                  {tab.label}
                </button>
              ))}
            </nav>
          </div>

          {/* --- ORDER LIST --- */}
          <div className="space-y-6">
            {filteredOrders.length > 0 ? (
              filteredOrders.map((order) => (
                <OrderCard key={order.id} order={order} />
              ))
            ) : (
              // --- EMPTY STATE ---
              <div className="bg-white rounded-lg shadow-sm text-center p-12">
                <FiArchive className="mx-auto text-gray-400" size={56} />
                <h2 className="mt-4 text-xl font-semibold text-gray-900">
                  No orders found yet
                </h2>
                <p className="mt-2 text-gray-500">
                  You don't have any orders with the status "{activeTab}".
                </p>
              </div>
            )}
          </div>
        </main>
      </div>
    </>
  );
};

export default OrderTrackingPage;
