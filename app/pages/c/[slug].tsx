import axios from 'axios'
import Post from 'components/Post'

const Category = ({tag}) =>{
    
  if(tag.length > 0){
    return (
      <>
      {tag?.map(t =>{
        return(
        t.posts?.map(p =>{
          return (
            <Post 
                key={p.id.toString()}
                id={p.id}
                title={p.title}
                description={p.description}
                tags={tag}
                created_at={p.created_at}
                />
          )
        })
        )
      })}
      </>
    )
  }

}



export async function getServerSideProps(context) {
    const {slug} = context.query
    const {APP_URL} = process.env
    const res = await axios(`${APP_URL}/tags?slug=${slug}`)
    const data = await res.data
    console.log(slug);
    
    return {
      props: {tag:data},
        }
  }
  
export default Category