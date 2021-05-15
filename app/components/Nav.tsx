import axios from 'axios'
import { Tag } from 'lib/interface'
import Link from 'next/link'

const Nav = ({tags}) => {

    return (
        <>
            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                        {
                        tags?.map(tag => (
                            <Link href={`/${tag.slug}`}>
                            <a  className="hover:bg-gray-400 rounded py-1 px-3">{tag.name}</a>
                            </Link>
                        ))}
                    </div>
                </div>
            </nav>
        </>
    )
}

export async function getServerSideProps(context) {
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/tags`)
    const data = await res.data
    console.log(JSON.stringify(data));
    
    return {
      props: {tags:data}, // will be passed to the page component as props
    }
  }
  

export default Nav




