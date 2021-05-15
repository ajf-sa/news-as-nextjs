import Loyout from 'components/layouts/Layouts'
import axios from 'axios'
import { PostType } from 'lib/interface'
import Post from 'components/Post'
export default function getPostById({post}){

    return (
        <>
         <Loyout>
        <Post key={post.id}
                        id={post.id}
                        title={post.title}
                        description={post.description}
                        tags={post.tags}
                        />
         </Loyout>
        </>
    )
}


export async function getServerSideProps(context) {
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/posts/${context.query.id}`)
    const data:PostType = res.data
    return {
      props: {post:data}, // will be passed to the page component as props
    }
  }
  