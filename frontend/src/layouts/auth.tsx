import React, { ReactNode } from "react";
import logo from "../assets/img/logo.jpg";

interface AuthLayoutProps {
  children: ReactNode;
}

const AuthLayout: React.FC<AuthLayoutProps> = ({ children }) => {
  return (
    <main className="w-screen h-screen overflow-hidden flex">
        <div className="flex-1 bg-primary-1 flex justify-center items-center">
            <img src={logo} alt="Company Logo" />
        </div>
        <div className="flex-1 flex flex-col justify-center items-center">
            {children}
        </div>
        </main>
  );
};

export default AuthLayout;
