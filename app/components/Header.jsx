import Link from "next/link"
import { Tag } from 'lib/interface'
import { useContext } from 'react'
import { useRouter } from 'next/router'
import HeaderContext from '../contexts/HeaderContext'
const Header = () => {
    const router = useRouter()
   
    const {menuItems} = useContext(HeaderContext)
 
    return (
        <>
            <header className="w-full container mx-auto">
                <div className="flex flex-col items-center py-5">
                    <div className="py-2">
                        <Link href="/">
                            <a className=" aref font-bold text-gray-800  hover:text-gray-700 text-6xl">
                                موجز
                            </a>
                        </Link>
                    </div>
                    <p className="aref text-xl text-gray-600  ">
                        اخبار واعمال ورياضة وترفيه بشكل موجز
                    </p>
                </div>
            </header>

            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                         {menuItems.map(item => (
                        <Link href={`/c/${item.slug}`} key={item.id}>
                            <a className={router.query.slug === item.slug ? 'bg-gray-400 rounded py-1 px-4 mx-2' : 'hover:bg-gray-400 rounded py-1 px-4 mx-2' }>{item.name}</a>
                        </Link>
                   
                ))}

                        

                    </div>
                </div>
            </nav>

        </>
    )
}


export default Header