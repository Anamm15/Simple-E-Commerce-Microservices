import { NavLink } from "react-router-dom";
import AuthLayout from "../../layouts/auth";
import { register } from "../../services/api/user";
import { UserDataType } from "../../utils/data-types";
import { useState } from "react";

const RegisterPage = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [gender, setGender] = useState('');
    const [birthPlace, setBirthPlace] = useState('');
    const [birthDate, setBirthDate] = useState('');
    const [password, setPassword] = useState('');
    const [authMessage, setAuthMessage] = useState('');

    const handleRegister = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const data : UserDataType = {
            name,
            email,
            username,
            gender,
            birth_date: birthDate,
            birth_place: birthPlace,
            password
        }
        try {
            const res = await register(data);
            if (res.status === 201) {
                setAuthMessage("Registrasi berhasil");
            }
        } catch (e: any) {
            if (e.response.status === 409) {
                setAuthMessage("Email sudah digunakan. Silakan gunakan email lain.");
            } else {
                setAuthMessage("Registrasi gagal" + e.message);
            }
        }
    };

    return (
        <>
            <AuthLayout>
                <form className="w-3/5 gap-4" onSubmit={handleRegister}>
                    <h4 className="text-4xl font-semibold text-center">Selamat Datang di Toko Kami</h4>
                    <p className="mt-2 text-lg text-center">Dapatkan pengalaman belanja yang menyenangkan di palform ini</p>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="name">Nama</label>
                        <input type="text" id="name" name="name" 
                            value={name} onChange={(e) => setName(e.target.value)} 
                            className="border border-black py-2 rounded-md px-2.5" 
                            placeholder="Masukkan nama Anda" />
                    </div>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="email">Email</label>
                        <input type="text" id="email" name="email" 
                            value={email} onChange={(e) => setEmail(e.target.value)} 
                            className="border border-black py-2 rounded-md px-2.5" 
                            placeholder="Masukkan email Anda" />
                    </div>
                    <div className="flex gap-2">
                        <div className="input__group flex flex-col gap-2 mt-4 flex-1">
                            <label htmlFor="username">Username</label>
                            <input type="text" id="username" name="username" 
                                value={username} onChange={(e) => setUsername(e.target.value)} 
                                className="border border-black py-2 rounded-md px-2.5" 
                                placeholder="Masukkan username Anda" />
                        </div>
                        <div className="input__group flex flex-col gap-2 mt-4 flex-1">
                            <label htmlFor="gender">Gender</label>
                            <select id="gender" name="gender" 
                                value={gender} onChange={(e) => setGender(e.target.value)} 
                                className="border border-black py-2 rounded-md px-2.5">
                                    <option value="M">Male</option>
                                    <option value="F">Female</option>
                            </select>
                        </div>
                    </div>
                    <div className="flex gap-2">
                        <div className="input__group flex flex-col gap-2 mt-4 flex-1">
                            <label htmlFor="birth_place">Tempat Lahir</label>
                            <input type="text" id="birth_place" name="birth_place" 
                                value={birthPlace} onChange={(e) => setBirthPlace(e.target.value)} 
                                className="border border-black py-2 rounded-md px-2.5" 
                                placeholder="Masukkan tempat lahir Anda" />
                        </div>
                        <div className="input__group flex flex-col gap-2 mt-4 flex-1">
                            <label htmlFor="birth_date">Tanggal Lahir</label>
                            <input type="date" id="birth_date" name="birth_date" 
                                value={birthDate} onChange={(e) => setBirthDate(e.target.value)} 
                                className="border border-black py-2 rounded-md px-2.5" />
                        </div>
                    </div>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password" name="password" 
                            value={password} onChange={(e) => setPassword(e.target.value)}
                            className="border border-black py-2 rounded-md px-2.5" 
                            placeholder="Masukkan password Anda" />
                        <a href="/forgot-password" className="text-blue-600 text-right">Lupa Password?</a>
                    </div>
                    <button type="submit" className="py-2 w-full bg-primary-1 mt-5 text-white font-semibold">Masuk</button>
                    <p className="text-red-600 mt-2">{authMessage}</p>
                    <p className="mt-2 flex gap-1">Sudah punya akun?
                        <NavLink to="/login" className="text-blue-600">Login disini!</NavLink>
                    </p>
                </form>
            </AuthLayout>
        </>
    )
}

export default RegisterPage;

