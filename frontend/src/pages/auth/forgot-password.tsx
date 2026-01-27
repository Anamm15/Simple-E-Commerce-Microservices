import AuthLayout from "../../layouts/auth";

const ForgotPasswordPage = () => {
    return (
        <>
            <AuthLayout>
                <form className="w-3/5 gap-4">
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="email">Email</label>
                        <input type="text" id="email" name="email" className="border border-black py-2 rounded-md px-2.5" placeholder="Masukkan email Anda" />
                    </div>
                    <div className="input__group flex flex-col gap-2 mt-4">
                        <label htmlFor="code">Kode</label>
                        <input type="text" id="code" name="code" className="border border-black py-2 rounded-md px-2.5" placeholder="Masukkan kode yang dikirimkan" />
                    </div>
                    <button type="submit" className="py-2 w-full bg-primary-1 mt-5 text-white font-semibold">Submit</button>
                </form>
            </AuthLayout>
        </>
    )
}

export default ForgotPasswordPage;