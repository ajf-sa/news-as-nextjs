 import Header from 'components/Header'
 import Footer from 'components/Footer'
import SideBar from 'components/SideBar'
import Nav from './Nav'
 const Layout = ({ children }) => {

    return (
        <>
           
            <div className="flex flex-wrap  overflow-hidden justify-center  py-4 ">
            <div className="my-3 w-full overflow-hidden  lg:w-1/2">
            {children}
             </div>
             <div className="my-3 w-full overflow-hidden sm:my-2  md:my-3 lg:mx-2 lg:px-2 lg:w-1/4">
                <SideBar />
            </div>
            </div>
            <Footer />
        </>
    )
}

export default Layout