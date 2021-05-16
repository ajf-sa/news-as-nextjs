
import Link from 'next/link'
import axios from 'axios'
import useSWR from 'swr'
import { Tag } from 'lib/interface'


const Nav = () => {
    const tags : Array<Tag> = [
        {id:1,name:"محليات",slug:"local"},
        {id:2,name:"رياضة",slug:"sports"}
    ]    
    return (
        <>
            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                        {tags?.map(tag => (
                            <Link href={`/c/${tag.slug}`} key={tag.id}>
                                <a className="hover:bg-gray-400 rounded py-1 px-3">{tag.name}</a>
                            </Link>
                        ))}

                    </div>
                </div>
            </nav>
        </>
    )
}

export default Nav




