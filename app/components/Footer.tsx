const Footer = () =>{
    return (
       
    <footer className="w-full border-t bg-white pb-12">
    <div className="w-full container mx-auto flex flex-col items-center">
        <div className="flex flex-row sm:flex-row text-center  justify-between py-6">
            <a href="#" className=" px-2 text-sm lg:text-lg">من نحن</a>
            <a href="#" className=" px-2 text-sm lg:text-lg">سياسة الخصوصية</a>
            <a href="#" className=" px-2 text-sm lg:text-lg">الشروط والاحكام</a>
            <a href="#" className=" px-2 text-sm lg:text-lg">اتصل بنا</a>
        </div>
        <div className="pb-6">&copy; ajf.sa</div>
    </div>
</footer>
    )
}


export default Footer