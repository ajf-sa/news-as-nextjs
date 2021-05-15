import  Loyout from 'components/layouts/Layouts'
const Home = ({slug})=> {

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