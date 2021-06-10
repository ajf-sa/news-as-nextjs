import axios from 'axios';
import Post from 'components/Post';
import { useRouter } from 'next/router';
import { PostType } from 'lib/interface';
import { useState } from 'react';

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

const Category: React.FC<{ post: PostType }> = (props) => {
  const { post } = props;
  const [enableLoadComments, setEnableLoadComments] = useState<boolean>(true);
  const router = useRouter();

  if (router.isFallback) {
    return <h1>Loading...</h1>;
  }

  return <>{<Post key={post.id.toString()} post={post} />}</>;
};

export default Category;
