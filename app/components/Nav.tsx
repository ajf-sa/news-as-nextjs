import axios from 'axios'
import { Tag } from 'lib/interface'
import Link from 'next/link'


const Nav = ({tags}) => {

    return (
        <>
            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                        {tags? ("tttt"):("sss") }
                    </div>
                </div>
            </nav>
        </>
    )
}


// Nav.getInitialProps = async (ctx) => {
//     const { APP_URL } = process.env
//     const res = await axios(`${APP_URL}/tags`)
//     const data = await res.data
//     console.log(JSON.stringify('ddddddddddddddddddddddddddddddddddddddddddd'));

//     return {
//         props: { tags: data }, // will be passed to the page component as props
//     }
//   }

// // export async function getServerSideProps(context) {
   
// // }


export default Nav




