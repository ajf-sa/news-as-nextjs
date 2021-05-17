import axios from 'axios'
import { PostType } from 'lib/interface'
import Post from 'components/Post'
import { useRouter } from 'next/router'
export default function getPostById({post}){
  const router = useRouter()

  if (router.isFallback) {
    return <div>Loading...</div>
  }

    return (
        <>
        <Post key={post.id}
                        id={post.id}
                        title={post.title}
                        description={post.description}
                        tags={post.tags}
                        created_at ={post.created_at}
                        image = {post.image}
                        />
        </>
    )
}


export async function getStaticProps(context) {
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/posts/${context.query.id}`)
    const data:PostType = res.data
    return {
      props: {post:data}, // will be passed to the page component as props
      revalidate: 60,
    }
  }

export async function getStaticPaths() {
   return { 
      paths: [],
      fallback: true }
  }
  