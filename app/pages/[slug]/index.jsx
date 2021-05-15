import {useRouter} from 'next/router'
import  Loyout from '../../components/layouts/Layouts'
const Home = ({slug})=> {
    // const router = useRouter()
    // const {slug} = router.query
  return (
    <>
    <Loyout>
    {slug}
    </Loyout>
      </>
  )
}

export async function getServerSideProps(context) { 

    return {
        props:{slug:context.params.slug}
    }

}


export default Home