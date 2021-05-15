import Nav from '../components/Nav'
const Header = () =>{
    return (
    <>
           <header className="w-full container mx-auto">
        <div className="flex flex-col items-center py-5">
            <div className="py-2">
            <a className=" aref font-bold text-gray-800  hover:text-gray-700 text-6xl" href="/">
               موجز
            </a>
            </div>
            <p className="aref text-xl text-gray-600  ">
              اخبار واعمال ورياضة وترفيه بشكل موجز
            </p>
        </div>
    </header>
        <Nav />
    </>
    )
}


export default Header