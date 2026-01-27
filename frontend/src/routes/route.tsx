import { createBrowserRouter } from "react-router-dom";
import LandingPage from "../pages/Landing";
import LoginPage from "../pages/auth/login";
import RegisterPage from "../pages/auth/register";
import ForgotPasswordPage from "../pages/auth/forgot-password";
import RecoveryPasswordPage from "../pages/auth/recovery-password";
import DetailProdukPage from "../pages/ProductDetail";
import CartPage from "../pages/Cart";
import OrderTrackingPage from "../pages/OrderTracking";
import ProfilePage from "../pages/Profile";
import CheckoutPage from "../pages/Checkout";
import OrderDetailPage from "../pages/OrderDetail";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
  {
    path: "/forgot-password",
    element: <ForgotPasswordPage />,
  },
  {
    path: "/recovery-password",
    element: <RecoveryPasswordPage />,
  },
  {
    path: "/profile",
    element: <ProfilePage />,
  },
  {
    path: "/product/:id",
    element: <DetailProdukPage />,
  },
  {
    path: "checkout",
    element: <CheckoutPage />,
  },
  {
    path: "cart",
    element: <CartPage />,
  },
  {
    path: "/order",
    element: <OrderTrackingPage />,
  },
  {
    path: "/order/:id",
    element: <OrderDetailPage />,
  },
]);

export default router;
