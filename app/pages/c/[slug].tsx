import axios from 'axios'

const Category = ({tag}) =>{
    console.log(tag.posts);

    return (
        <div>
         {tag.posts}
        </div>
    )
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