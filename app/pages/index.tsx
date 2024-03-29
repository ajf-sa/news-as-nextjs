import Post from 'components/Post'
import axios from 'axios'
import { PostType } from 'lib/interface'

const Home = ({ allPosts, preview }) => {

    const heroPost = allPosts[0]
    const morePosts = allPosts.slice(1)

    return (
        <>
                <section className="w-full flex flex-col items-center px-3">
                    
                    {morePosts.length > 0 && allPosts.map(post=>(
                        <Post 
                        key={post.id.toString()}
                        id={post.id}
                        title={post.title}
                        description={post.description}
                        tags={post.tags}
                        created_at={post.created_at}
                        image={post.image}
                        />
                    ))}

                    <div className="flex items-center py-12">
                        <a href="#" className="h-10 w-10 bg-cyan-800 hover:bg-cyan-600 font-semibold text-white text-sm flex items-center justify-center ml-3">1</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:bg-cyan-600 hover:text-white text-sm flex items-center justify-center ml-3">2</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3">التالي <i className="fas fa-arrow-right ml-2"></i></a>
                    </div>

                </section>
        </>
    )
}

// //https://admin.ultraify.com/posts
export async function getServerSideProps({ preview = null }) {
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/posts`)
    const allPosts:PostType = ( await res.data ) || []
    return {
      props: { allPosts, preview },
    }
  }


// export async function getStaticProps({ preview = null }) {

//     const {APP_URL} = process.env
//     const res = await axios(`${APP_URL}/posts`)
//     const allPosts:PostType = ( await res.data ) || []
//     return {
//       props: { allPosts, preview },
//     }
//   }

  
export default Home