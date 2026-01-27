import React from "react";
import { useNavigate } from "react-router-dom";
import { Order } from "../types";
import {
  FiLoader,
  FiTruck,
  FiCheckCircle,
  FiChevronRight,
} from "react-icons/fi";

const statusConfig = {
  processing: {
    icon: <FiLoader className="mr-1.5 animate-spin" />,
    text: "Sedang Diproses",
    color: "bg-yellow-100 text-yellow-800",
  },
  shipping: {
    icon: <FiTruck className="mr-1.5" />,
    text: "Dalam Perjalanan",
    color: "bg-blue-100 text-blue-800",
  },
  completed: {
    icon: <FiCheckCircle className="mr-1.5" />,
    text: "Selesai",
    color: "bg-green-100 text-green-800",
  },
  cancelled: {
    icon: <FiCheckCircle className="mr-1.5" />,
    text: "Dibatalkan",
    color: "bg-red-100 text-red-800",
  },
  returned: {
    icon: <FiCheckCircle className="mr-1.5" />,
    text: "Dikembalikan",
    color: "bg-purple-100 text-purple-800",
  },
};

interface OrderCardProps {
  order: Order;
}

const OrderCard: React.FC<OrderCardProps> = ({ order }) => {
  const config = statusConfig[order.status];
  const navigate = useNavigate();

  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden transition-all duration-300 hover:shadow-lg">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center p-4 bg-gray-50 border-b border-gray-200">
        <div>
          <h2 className="text-base font-semibold text-gray-900">
            Order #{order.id}
          </h2>
          <p className="text-sm text-gray-500">Ordered on: {order.date}</p>
        </div>
        <div className="mt-2 sm:mt-0 text-left sm:text-right">
          <span className="text-sm text-gray-500">Total:</span>
          <p className="text-lg font-bold text-gray-900">
            ${order.total.toFixed(2)}
          </p>
        </div>
      </div>

      <div className="p-4">
        {order.items.map((item, index) => (
          <div
            key={item.id + index}
            className="flex items-center mb-4 last:mb-0"
          >
            <img
              src={item.imageUrl}
              alt={item.name}
              className="w-16 h-16 rounded-md object-cover flex-shrink-0"
            />
            <div className="ml-4 flex-1 min-w-0">
              <h3 className="text-sm font-medium text-gray-900 truncate">
                {item.name}
              </h3>
              <p className="text-sm text-gray-500">Quantity: {item.quantity}</p>
            </div>
          </div>
        ))}
      </div>

      <div className="flex justify-between items-center p-4 bg-gray-50 border-t border-gray-200">
        <span
          className={`
            inline-flex items-center px-3 py-1 rounded-full 
            text-xs font-semibold ${config.color}
          `}
        >
          {config.icon}
          {config.text}
        </span>

        <button
          onClick={() => navigate(`/order/${order.id}`)}
          className="flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
        >
          View Details
          <FiChevronRight className="ml-1" size={16} />
        </button>
      </div>

      {order.status === "shipping" && (
        <div className="p-4 border-t border-gray-200 bg-blue-50">
          <p className="text-sm text-blue-700">
            <span className="font-semibold">Tracking Number:</span>{" "}
            {order.trackingNumber}
          </p>
          <p className="text-sm text-blue-700 mt-1">
            <span className="font-semibold">Estimated Delivery:</span>{" "}
            {order.estimatedDelivery}
          </p>
        </div>
      )}
    </div>
  );
};

export default OrderCard;
