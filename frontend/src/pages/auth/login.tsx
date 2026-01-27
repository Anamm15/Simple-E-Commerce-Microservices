import { useState } from "react";
import AuthLayout from "../../layouts/auth";
import { login } from "../../services/api/user";
import { NavLink, useNavigate } from "react-router-dom";

const LoginPage = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!username || !password) {
            setErrorMessage("Username dan password harus diisi");
            return;
        }
    
        try {
            const res = await login({ username, password });
            if (res.status === 200) {
                localStorage.setItem('token', res.data.accessToken);
                navigate('/');
            }
        } catch {
            setErrorMessage("Username atau password salah");
        }
    };    

    return (
        <>
            <AuthLayout>
                <form className="w-3/5 gap-4" onSubmit={handleLogin}>
                    <h4 className="text-4xl font-semibold text-center">Selamat Datang di Toko Kami</h4>
                    <p className="mt-2 text-lg text-center">Dapatkan pengalaman belanja yang menyenangkan di palform ini</p>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="username">Username</label>
                        <input type="text" id="username" name="username" 
                            className="border border-black py-2 rounded-md px-2.5" 
                            placeholder="Masukkan username Anda"
                            onChange={(e) => setUsername(e.target.value)} />
                    </div>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="password">Password</label>
                        <input type="password" value={password} id="password" name="password" 
                            className="border border-black py-2 rounded-md px-2.5" 
                            placeholder="Masukkan password Anda"
                            onChange={(e) => setPassword(e.target.value)} />
                        <NavLink to="/forgot-password" className="text-blue-600 text-right">Lupa Password?</NavLink>
                    </div>
                    <button type="submit" 
                        className="py-2 w-full bg-primary-1 mt-5 text-white font-semibold">
                        Masuk
                    </button>
                    <p className="mt-2 flex gap-1">Belum punya akun?
                        <NavLink to="/register" className="text-blue-600">Daftar sekarang!</NavLink>
                    </p>
                    {errorMessage && <p className="text-red-600 mt-2">{errorMessage}</p>}
                </form>
            </AuthLayout>
        </>
    )
}

export default LoginPage;