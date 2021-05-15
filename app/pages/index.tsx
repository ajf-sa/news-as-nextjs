import Loyout from 'components/layouts/Layouts'
import { DateNow } from 'components/dateNow'
import Post from 'components/Post'

import axios from 'axios'
import { PostType } from 'lib/interface'
const Home = ({posts}) => {

    return (
        <>
            <Loyout>
                <section className="w-full flex flex-col items-center px-3">
                    {posts.map(post=>(
                        <Post key={post.id}
                        id={post.id}
                        title={post.title}
                        description={post.description}
                        tags={post.tags}
                        />
                    ))}

                    <div className="flex items-center py-8">
                        <a href="#" className="h-10 w-10 bg-cyan-800 hover:bg-cyan-600 font-semibold text-white text-sm flex items-center justify-center ml-3">1</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:bg-cyan-600 hover:text-white text-sm flex items-center justify-center ml-3">2</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3">التالي <i className="fas fa-arrow-right ml-2"></i></a>
                    </div>

                </section>


            </Loyout>
        </>
    )
}

//https://admin.ultraify.com/posts
export async function getServerSideProps(context) {
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/posts`)
    const data:PostType = await res.data
    return {
      props: {posts:data}, // will be passed to the page component as props
    }
  }
  
export default Home