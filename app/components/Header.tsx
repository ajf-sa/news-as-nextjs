import Nav from '../components/Nav'
const Header = () =>{
    return (
    <>
           <header className="w-full container mx-auto">
        <div className="flex flex-col items-center py-12">
            <div className="py-6">
            <a className="font-bold text-gray-800 uppercase hover:text-gray-700 text-5xl" href="/">
               موجز
            </a>
            </div>
            <p className="text-lg text-gray-600 ">
              اخبار واعمال ورياضة وترفيه بشكل موجز
            </p>
        </div>
    </header>
        <Nav />
    </>
    )
}


export default Header