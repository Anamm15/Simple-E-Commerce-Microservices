import React, { useState } from "react";
import { UserProfile, UserAddress, Order } from "../types";
import { dummyUser } from "../data/user";
import { dummyOrders } from "../data/order";
import OrderCard from "../components/OrderCard";
import {
  FiUser,
  FiMapPin,
  FiArchive,
  FiLogOut,
  FiEdit3,
  FiPlus,
  FiTrash2,
  FiHome,
} from "react-icons/fi";
import Navbar from "../components/Navbar";

type ProfileTab = "profile" | "addresses" | "orders";

const ProfileInfoSection: React.FC<{ user: UserProfile }> = ({ user }) => {
  return (
    <div className="space-y-6">
      {/* --- PERSONAL INFORMATION CARD --- */}
      <div className="bg-white p-6 rounded-lg shadow-md">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold text-gray-900">
            Personal Information
          </h2>
          <button className="flex items-center text-sm font-medium text-blue-600 hover:text-blue-500">
            <FiEdit3 className="mr-1" size={16} /> Edit
          </button>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-500">
              Full Name
            </label>
            <p className="mt-1 text-base text-gray-900">{user.fullName}</p>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-500">
              Email
            </label>
            <p className="mt-1 text-base text-gray-900">{user.email}</p>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-500">
              Member Since
            </label>
            <p className="mt-1 text-base text-gray-900">{user.memberSince}</p>
          </div>
        </div>
      </div>

      {/* --- CHANGE PASSWORD CARD --- */}
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold text-gray-900 mb-4">
          Change Password
        </h2>
        <form className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Current Password
            </label>
            <input
              type="password"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">
              New Password
            </label>
            <input
              type="password"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Confirm New Password
            </label>
            <input
              type="password"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
          </div>
          <button
            type="submit"
            className="inline-flex justify-center rounded-lg border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700"
          >
            Update Password
          </button>
        </form>
      </div>
    </div>
  );
};

const ManageAddressesSection: React.FC<{ addresses: UserAddress[] }> = ({
  addresses,
}) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-semibold text-gray-900">Addresses Book</h2>
        <button className="flex items-center text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg px-4 py-2">
          <FiPlus className="mr-1" size={16} /> Add Address
        </button>
      </div>

      <div className="space-y-5">
        {addresses.map((addr) => (
          <div key={addr.id} className="border border-gray-200 rounded-lg p-4">
            <div className="flex justify-between items-start">
              <div>
                <div className="flex items-center space-x-2">
                  <FiHome className="text-gray-600" />
                  <span className="text-lg font-semibold text-gray-900">
                    {addr.label}
                  </span>
                  {addr.isDefault && (
                    <span className="text-xs font-medium text-green-800 bg-green-100 px-2 py-0.5 rounded-full">
                      Default
                    </span>
                  )}
                </div>
                <p className="mt-2 text-sm text-gray-700">
                  {addr.recipientName}
                </p>
                <p className="text-sm text-gray-500">{addr.phone}</p>
                <p className="mt-1 text-sm text-gray-500">{addr.address}</p>
                <p className="text-sm text-gray-500">
                  {addr.city}, {addr.postalCode}
                </p>
              </div>
              <div className="flex space-x-3 flex-shrink-0 ml-4">
                <button
                  className="text-blue-600 hover:text-blue-500"
                  title="Edit"
                >
                  <FiEdit3 size={18} />
                </button>
                <button
                  className="text-red-600 hover:text-red-500"
                  title="Hapus"
                >
                  <FiTrash2 size={18} />
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

const OrderHistorySection: React.FC<{ orders: Order[] }> = ({ orders }) => {
  return (
    <div className="space-y-6">
      <h2 className="text-xl font-semibold text-gray-900">Order History</h2>
      {orders.length > 0 ? (
        orders.map((order) => <OrderCard key={order.id} order={order} />)
      ) : (
        <p className="text-gray-500">You have no order history.</p>
      )}
    </div>
  );
};

const ProfilePage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<ProfileTab>("profile");
  const user = dummyUser;
  const orders = dummyOrders;

  const navItems = [
    { id: "profile", label: "Account", icon: FiUser },
    { id: "addresses", label: "Addresses Book", icon: FiMapPin },
    { id: "orders", label: "Order History", icon: FiArchive },
  ];

  const renderContent = () => {
    switch (activeTab) {
      case "profile":
        return <ProfileInfoSection user={user} />;
      case "addresses":
        return <ManageAddressesSection addresses={user.addresses} />;
      case "orders":
        return <OrderHistorySection orders={orders} />;
      default:
        return null;
    }
  };

  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen py-12">
        <main className="container mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="mb-8">
            <h1 className="text-3xl font-bold text-gray-900">My Account</h1>
            <p className="mt-1 text-lg text-gray-600">
              Welcome back, {user.fullName}!
            </p>
          </div>

          <div className="lg:grid lg:grid-cols-4 lg:gap-8">
            {/* --- MOBILE --- */}
            <aside className="lg:col-span-1">
              {/* DEKSTOP */}
              <nav className="hidden lg:flex flex-col space-y-2">
                {navItems.map((item) => (
                  <button
                    key={item.id}
                    onClick={() => setActiveTab(item.id as ProfileTab)}
                    className={`
                    flex items-center px-4 py-3 rounded-lg text-base font-medium
                    transition-colors
                    ${
                      activeTab === item.id
                        ? "bg-blue-50 text-blue-600"
                        : "text-gray-600 hover:bg-gray-100 hover:text-gray-900"
                    }
                  `}
                  >
                    <item.icon className="mr-3" size={20} />
                    {item.label}
                  </button>
                ))}
                <button
                  onClick={() => alert("Logout!")}
                  className="flex items-center px-4 py-3 rounded-lg text-base font-medium text-gray-600 hover:bg-gray-100 hover:text-gray-900"
                >
                  <FiLogOut className="mr-3" size={20} />
                  Logout
                </button>
              </nav>

              {/* MOBILE */}
              <nav className="lg:hidden mb-6">
                <div className="flex space-x-2 overflow-x-auto pb-2">
                  {navItems.map((item) => (
                    <button
                      key={item.id}
                      onClick={() => setActiveTab(item.id as ProfileTab)}
                      className={`
                      flex-shrink-0 flex items-center px-4 py-2 rounded-lg text-sm font-medium
                      transition-colors
                      ${
                        activeTab === item.id
                          ? "bg-blue-600 text-white"
                          : "bg-white text-gray-600 shadow-sm"
                      }
                    `}
                    >
                      <item.icon className="mr-2" size={16} />
                      {item.label}
                    </button>
                  ))}
                </div>
              </nav>
            </aside>

            <div className="lg:col-span-3">{renderContent()}</div>
          </div>
        </main>
      </div>
    </>
  );
};

export default ProfilePage;
