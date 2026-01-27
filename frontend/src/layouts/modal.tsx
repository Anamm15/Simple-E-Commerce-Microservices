import React, { ReactNode } from "react";
import closeIcon from "../assets/icons/close-icon.svg";

interface ModalProps {
    children: ReactNode;
    onClick: () => void;
}

const Modal: React.FC<ModalProps> = ({ children, onClick }) => {
    return (
        <div className="fixed z-50 top-0 left-0 flex justify-center items-center w-screen h-screen">
            <div className="w-1/4 max-h-[80vh] bg-white shadow-lg relative px-4 py-6 rounded-md overflow-auto">
                <button 
                    onClick={onClick} 
                    className="absolute top-2 right-2">
                    <img src={closeIcon} alt="" />
                </button>
                {children}
            </div>
        </div>
    );
};

export default Modal;
