
const Home = ({slug})=> {

  return (
    <>
    {slug}
    </>
  )
}

export async function getServerSideProps(context) { 

    return {
        props:{slug:context.params.slug}
    }

}


export default Home