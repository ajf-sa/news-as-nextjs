import axios from 'axios';
import { PostType } from 'lib/interface';
import Post from 'components/Post';
import { useRouter } from 'next/router';

const { APP_URL, GHOST_CONTENT_KEY } = process.env;

async function getPost(slug: string) {
  const res = await axios(
    `${APP_URL}/ghost/api/v4/content/posts/slug/${slug}?key=${GHOST_CONTENT_KEY}` //&fields=title,slug,html
  );
  const data = await res.data;

  return data.posts[0];
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

const getPostById: React.FC<{ post: PostType }> = (props) => {
  const { post } = props;

  const router = useRouter();
  if (router.isFallback) {
    return <h1>تحميل...</h1>;
  }
  return (
    <>
      <Post key={post.slug} post={post} />
    </>
  );
};

export default getPostById;
