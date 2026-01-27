import AuthLayout from "../../layouts/auth";

const RecoveryPasswordPage = () => {
    return (
        <>
            <AuthLayout>
                <form className="w-3/5 gap-4">
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password" name="password" className="border border-black py-2 rounded-md px-2.5" placeholder="Masukkan password Anda" />
                    </div>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="confirm_password">Konfirmasi Password</label>
                        <input type="password" id="confirm_password" name="confirm_password" className="border border-black py-2 rounded-md px-2.5" placeholder="Konfirmasi password Anda" />
                    </div>
                    <button type="submit" className="py-2 w-full bg-primary-1 mt-5 text-white font-semibold">Submit</button>
                </form>
            </AuthLayout>
        </>
    )
}

export default RecoveryPasswordPage;