const Nav = () => {
    return (
        <>
            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">محليات</a>
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">عالمي</a>
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">رياضة</a>
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">اعمال</a>
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">تقنية</a>
                        <a href="#" className="hover:bg-gray-400 rounded py-1 px-3">ترفيه</a>
                    </div>
                </div>
            </nav>
        </>
    )
}


export default Nav