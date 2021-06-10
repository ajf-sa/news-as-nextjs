import Post from 'components/Post';
import axios from 'axios';
import { PostType } from 'lib/interface';

type Url = {
  posts: Pos[];
};
type Pos = {
  id: string;
  title: string;
  slug: string;
  description: string;
  created_at: any;
  feature_image: any;
};

async function getPosts() {
  const { APP_URL, GHOST_CONTENT_KEY } = process.env;
  const res = await axios(
    `${APP_URL}/ghost/api/v4/content/posts/?key=${GHOST_CONTENT_KEY}`
  );
  const data: PostType = await res.data;
  return data;
}

export const getStaticProps = async ({ params }) => {
  const posts = await getPosts();
  return {
    revalidate: 10,
    props: { posts },
  };
};

const Home: React.FC<{ posts: Url }> = (props) => {
  const { posts } = props;
  return (
    <>
      <section className="w-full flex flex-col items-center px-3">
        {posts.posts.map((post) => (
          <Post
            key={post.id.toString()}
            id={post.id}
            title={post.title}
            slug={post.slug}
            description={post.description}
            created_at={post.created_at}
            image={post.feature_image}
          />
        ))}

        <div className="flex items-center py-12">
          <a
            href="#"
            className="h-10 w-10 bg-cyan-800 hover:bg-cyan-600 font-semibold text-white text-sm flex items-center justify-center ml-3"
          >
            1
          </a>
          <a
            href="#"
            className="h-10 w-10 font-semibold text-gray-800 hover:bg-cyan-600 hover:text-white text-sm flex items-center justify-center ml-3"
          >
            2
          </a>
          <a
            href="#"
            className="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3"
          >
            التالي <i className="fas fa-arrow-right ml-2"></i>
          </a>
        </div>
      </section>
    </>
  );
};

export default Home;
