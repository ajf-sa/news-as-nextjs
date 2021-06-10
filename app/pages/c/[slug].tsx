import axios from 'axios';
import Post from 'components/Post';
import { useRouter } from 'next/router';

async function getPost(slug: string) {
  const { APP_URL, GHOST_CONTENT_KEY } = process.env;
  const res = await axios(
    `${APP_URL}/ghost/api/v4/content/posts/slug/${slug}?key=${GHOST_CONTENT_KEY}`
  );
  const data = await res.data;
  return data[0];
}

export const getStaticProps = async ({ params }) => {
  const post = await getPost(params.slug);
  return {
    props: { post },
    revalidate: 10,
  };
};

export const getStaticPaths = () => {
  // paths -> slugs which are allowed
  // fallback ->
  return {
    paths: [],
    fallback: true,
  };
};

type Pos = {
  id: string;
  title: string;
  slug: string;
  description: string;
  created_at: any;
  feature_image: any;
};

const Category: React.FC<{ post: Pos }> = (props) => {
  const { post } = props;
  const router = useRouter();

  if (router.isFallback) {
    return <h1>Loading...</h1>;
  }
  return (
    <>
      {
        <Post
          key={post.id.toString()}
          id={post.id}
          title={post.title}
          slug={post.slug}
          description={post.description}
          created_at={post.created_at}
          image={post.feature_image}
        />
      }
    </>
  );
};

export default Category;
